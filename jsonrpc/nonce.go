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
	"github.com/attestantio/go-starknet-client/types"
)

// Nonce returns the nonce of the given contract at the given block.
func (s *Service) Nonce(ctx context.Context,
	opts *api.NonceOpts,
) (
	*api.Response[uint32],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}
	if opts.Block == "" {
		return nil, errors.Join(errors.New("no block specified"), client.ErrInvalidOptions)
	}

	var res types.Number
	err := s.client.CallFor(&res, "starknet_getNonce", opts.Block, opts.Contract.String())
	if err != nil {
		return nil, errors.Join(err, client.ErrRPCCallFailed)
	}

	return &api.Response[uint32]{
		Data:     uint32(res),
		Metadata: map[string]any{},
	}, nil
}