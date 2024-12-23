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
	"errors"
	"time"

	"github.com/attestantio/go-starknet-client/metrics"
	"github.com/rs/zerolog"
)

type parameters struct {
	logLevel          zerolog.Level
	monitor           metrics.Service
	address           string
	webSocketAddress  string
	timeout           time.Duration
	allowDelayedStart bool
}

// Parameter is the interface for service parameters.
type Parameter interface {
	apply(p *parameters)
}

type parameterFunc func(*parameters)

func (f parameterFunc) apply(p *parameters) {
	f(p)
}

// WithLogLevel sets the log level for the module.
func WithLogLevel(logLevel zerolog.Level) Parameter {
	return parameterFunc(func(p *parameters) {
		p.logLevel = logLevel
	})
}

// WithMonitor sets the monitor for the service.
func WithMonitor(monitor metrics.Service) Parameter {
	return parameterFunc(func(p *parameters) {
		p.monitor = monitor
	})
}

// WithAddress provides the address for the endpoint.
func WithAddress(address string) Parameter {
	return parameterFunc(func(p *parameters) {
		p.address = address
	})
}

// WithWebSocketAddress provides the address for the websocket endpoint.
// If not supplied it will use the value supplied as the address.
func WithWebSocketAddress(address string) Parameter {
	return parameterFunc(func(p *parameters) {
		p.webSocketAddress = address
	})
}

// WithTimeout sets the maximum duration for all requests to the endpoint.
func WithTimeout(timeout time.Duration) Parameter {
	return parameterFunc(func(p *parameters) {
		p.timeout = timeout
	})
}

// WithAllowDelayedStart allows the service to start even if the client is unavailable.
func WithAllowDelayedStart(allowDelayedStart bool) Parameter {
	return parameterFunc(func(p *parameters) {
		p.allowDelayedStart = allowDelayedStart
	})
}

// parseAndCheckParameters parses and checks parameters to ensure that mandatory parameters are present and correct.
func parseAndCheckParameters(params ...Parameter) (*parameters, error) {
	parameters := parameters{
		logLevel: zerolog.GlobalLevel(),
		timeout:  2 * time.Second,
	}
	for _, p := range params {
		if params != nil {
			p.apply(&parameters)
		}
	}

	if parameters.address == "" {
		return nil, errors.New("no address specified")
	}
	if parameters.webSocketAddress == "" {
		parameters.webSocketAddress = parameters.address
	}
	if parameters.timeout == 0 {
		return nil, errors.New("no timeout specified")
	}

	return &parameters, nil
}
