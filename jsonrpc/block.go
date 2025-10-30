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

// Block returns the block as per the given parameters.
func (s *Service) Block(ctx context.Context,
	opts *api.BlockOpts,
) (
	*api.Response[*spec.Block],
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

	rpcOpts := map[string]any{
		"block_id": opts.Block,
	}

	var data spec.Block

	err := s.client.CallFor(&data, "starknet_getBlockWithReceipts", rpcOpts)
	if err != nil {
		return nil, errors.Join(errors.New("starknet_getBlockWithReceipts failed"), err)
	}

	return &api.Response[*spec.Block]{
		Data:     &data,
		Metadata: map[string]any{},
	}, nil
}
