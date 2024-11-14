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

	client "github.com/attestantio/go-starknet-client"
	"github.com/attestantio/go-starknet-client/api"
	"github.com/attestantio/go-starknet-client/spec"
)

type eventsResJSON struct {
	Events []*spec.TransactionEvent `json:"events"`
}

// Events returns the events matching the filter.
func (s *Service) Events(ctx context.Context,
	opts *api.EventsOpts,
) (
	*api.Response[[]*spec.TransactionEvent],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}
	if opts.Limit <= 0 {
		return nil, errors.Join(errors.New("limit must be specified"), client.ErrInvalidOptions)
	}

	var res eventsResJSON
	if err := s.client.CallFor(&res, "starknet_getEvents", opts); err != nil {
		return nil, err
	}

	return &api.Response[[]*spec.TransactionEvent]{
		Data:     res.Events,
		Metadata: map[string]any{},
	}, nil
}
