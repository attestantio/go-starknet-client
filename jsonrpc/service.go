// Copyright © 2024 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	client "github.com/attestantio/go-starknet-client"
	"github.com/attestantio/go-starknet-client/api"
	"github.com/rs/zerolog"
	zerologger "github.com/rs/zerolog/log"
	"github.com/ybbus/jsonrpc/v2"
	"golang.org/x/sync/semaphore"
)

// Service is an Starknet client service.
type Service struct {
	base             *url.URL
	address          string
	webSocketAddress string
	client           jsonrpc.RPCClient
	timeout          time.Duration
	// Endpoint support.
	pingSem          *semaphore.Weighted
	connectionMu     sync.RWMutex
	connectionActive bool
	connectionSynced bool
}

// log is a service-wide logger.
var log zerolog.Logger

// New creates a new execution client service, connecting with a standard HTTP.
func New(ctx context.Context, params ...Parameter) (*Service, error) {
	parameters, err := parseAndCheckParameters(params...)
	if err != nil {
		return nil, err
	}

	// Set logging.
	log = zerologger.With().Str("service", "client").Str("impl", "jsonrpc").Logger()
	if parameters.logLevel != log.GetLevel() {
		log = log.Level(parameters.logLevel)
	}

	if parameters.monitor != nil {
		if err := registerMetrics(ctx, parameters.monitor); err != nil {
			return nil, errors.Join(errors.New("failed to register metrics"), err)
		}
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:        64,
			MaxIdleConnsPerHost: 64,
			IdleConnTimeout:     384 * time.Second,
		},
	}

	base, address, err := parseAddress(parameters.address)
	if err != nil {
		return nil, err
	}

	webSocketAddress := parameters.webSocketAddress
	if strings.HasPrefix(webSocketAddress, "http://") {
		webSocketAddress = fmt.Sprintf("ws://%s", webSocketAddress[7:])
	}
	if strings.HasPrefix(webSocketAddress, "https://") {
		webSocketAddress = fmt.Sprintf("wss://%s", webSocketAddress[8:])
	}
	if !strings.HasPrefix(webSocketAddress, "ws") {
		webSocketAddress = fmt.Sprintf("ws://%s", webSocketAddress)
	}
	log.Trace().Stringer("address", address).Str("web_socket_address", webSocketAddress).Msg("Addresses configured")

	extraHeaders := map[string]string{
		"User-Agent": "go-starknet-client/0.1.0",
	}

	rpcClient := jsonrpc.NewClientWithOpts(address.String(), &jsonrpc.RPCClientOpts{
		HTTPClient:    httpClient,
		CustomHeaders: extraHeaders,
	})

	s := &Service{
		base:             base,
		client:           rpcClient,
		address:          address.String(),
		webSocketAddress: webSocketAddress,
		timeout:          parameters.timeout,
		pingSem:          semaphore.NewWeighted(1),
	}

	// Ping the client to see if it is ready to serve requests.
	s.CheckConnectionState(ctx)
	active := s.IsActive()

	if !active && !parameters.allowDelayedStart {
		return nil, client.ErrNotActive
	}

	// Fetch static values to confirm the connection is good.
	if err := s.fetchStaticValues(ctx); err != nil {
		return nil, errors.Join(errors.New("failed to confirm node connection"), err)
	}

	// Periodically ping the client for state updates.  We do this so that
	// even if we aren't actively using the connection its state will be
	// roughly up-to-date.
	s.periodicUpdateConnectionState(ctx)

	// Close the service on context done.
	go func(s *Service) {
		<-ctx.Done()
		log.Trace().Msg("Context done; closing connection")
		s.close()
	}(s)

	return s, nil
}

// periodicUpdateConnectionState periodically pings the client to update its active and synced status.
func (s *Service) periodicUpdateConnectionState(ctx context.Context) {
	go func(s *Service, ctx context.Context) {
		// Refresh every 30 seconds.
		refreshTicker := time.NewTicker(30 * time.Second)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				s.CheckConnectionState(ctx)
			case <-ctx.Done():
				return
			}
		}
	}(s, ctx)
}

// CheckConnectionState checks the connection state for the client, potentially updating
// its activation and sync states.
// This will call hooks supplied when creating the client if the state changes.
func (s *Service) CheckConnectionState(ctx context.Context) {
	log := zerolog.Ctx(ctx)

	s.connectionMu.Lock()
	wasActive := s.connectionActive
	wasSynced := s.connectionSynced
	s.connectionMu.Unlock()

	var active bool
	var synced bool

	acquired := s.pingSem.TryAcquire(1)
	if !acquired {
		// Means there is another ping running, just use current info.
		active = wasActive
		synced = wasSynced
	} else {
		response, err := s.Syncing(ctx, &api.SyncingOpts{})
		if err != nil {
			log.Debug().Err(err).Msg("Failed to obtain sync state from node")
			active = false
			synced = false
		} else {
			active = true
			synced = !response.Data.Syncing
		}
		s.pingSem.Release(1)
	}

	// if !wasActive && active {
	// 	// Switched from not active to active.
	// }

	// if wasActive && !active {
	// 	// Switched from active to not active.
	// }

	// if !wasSynced && synced {
	// 	// Switched from not synced to synced.
	// }

	// if wasSynced && !synced {
	// 	// Switched from synced to not synced.
	// }

	log.Trace().
		Bool("was_active", wasActive).
		Bool("active", active).
		Bool("was_synced", wasSynced).
		Bool("synced", synced).
		Msg("Updated connection state")

	s.connectionMu.Lock()
	s.connectionActive = active
	s.connectionSynced = synced
	s.connectionMu.Unlock()

	switch {
	case synced:
		s.monitorState("synced")
	case active:
		s.monitorState("active")
	default:
		s.monitorState("inactive")
	}
}

// fetchStaticValues fetches values that never change.
// This caches the values, avoiding future API calls.
func (*Service) fetchStaticValues(_ context.Context) error {
	return nil
}

// Name provides the name of the service.
func (*Service) Name() string {
	return "json-rpc"
}

// Address provides the address for the connection.
func (s *Service) Address() string {
	return s.address
}

// close closes the service, freeing up resources.
func (*Service) close() {
}

// IsActive returns true if the client is active.
func (s *Service) IsActive() bool {
	s.connectionMu.RLock()
	active := s.connectionActive
	s.connectionMu.RUnlock()

	return active
}

// IsSynced returns true if the client is synced.
func (s *Service) IsSynced() bool {
	s.connectionMu.RLock()
	synced := s.connectionSynced
	s.connectionMu.RUnlock()

	return synced
}

func (s *Service) assertIsActive(ctx context.Context) error {
	active := s.IsActive()
	if active {
		return nil
	}

	s.CheckConnectionState(ctx)
	active = s.IsActive()
	if !active {
		return client.ErrNotActive
	}

	return nil
}

func (s *Service) assertIsSynced(ctx context.Context) error {
	synced := s.IsSynced()
	if synced {
		return nil
	}

	s.CheckConnectionState(ctx)
	active := s.IsActive()
	if !active {
		return client.ErrNotActive
	}

	synced = s.IsSynced()
	if !synced {
		return client.ErrNotSynced
	}

	return nil
}

func parseAddress(address string) (*url.URL, *url.URL, error) {
	if !strings.HasPrefix(address, "http") {
		address = fmt.Sprintf("http://%s", address)
	}
	base, err := url.Parse(address)
	if err != nil {
		return nil, nil, errors.Join(errors.New("invalid URL"), err)
	}
	// Remove any trailing slash from the path.
	base.Path = strings.TrimSuffix(base.Path, "/")

	// Attempt to mask any sensitive information in the URL, for logging purposes.
	baseAddress := *base
	if _, pwExists := baseAddress.User.Password(); pwExists {
		// Mask the password.
		user := baseAddress.User.Username()
		baseAddress.User = url.UserPassword(user, "xxxxx")
	}
	if baseAddress.Path != "" {
		// Mask the path.
		baseAddress.Path = "xxxxx"
	}
	if baseAddress.RawQuery != "" {
		// Mask all query values.
		sensitiveRegex := regexp.MustCompile("=([^&]*)(&)?")
		baseAddress.RawQuery = sensitiveRegex.ReplaceAllString(baseAddress.RawQuery, "=xxxxx$2")
	}

	return base, &baseAddress, nil
}
