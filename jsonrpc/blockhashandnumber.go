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

type blockHashAndNumberRes struct {
	BlockHash   types.Hash `json:"block_hash"`
	BlockNumber uint32     `json:"block_number"`
}

// BlockHashAndNumber returns the hash and number of the latest block as understood by the node.
func (s *Service) BlockHashAndNumber(ctx context.Context,
	opts *api.BlockHashAndNumberOpts,
) (
	*api.Response[*api.BlockHashAndNumber],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	res := blockHashAndNumberRes{}
	if err := s.client.CallFor(&res, "starknet_blockHashAndNumber"); err != nil {
		return nil, errors.Join(err, client.ErrRPCCallFailed)
	}

	return &api.Response[*api.BlockHashAndNumber]{
		Data: &api.BlockHashAndNumber{
			Hash:   res.BlockHash,
			Number: res.BlockNumber,
		},
	}, nil
}
