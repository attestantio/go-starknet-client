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
)

// BlockNumber returns the number of the latest block as understood by the node.
func (s *Service) BlockNumber(ctx context.Context,
	opts *api.BlockNumberOpts,
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

	data := uint32(0)
	if err := s.client.CallFor(&data, "starknet_blockNumber"); err != nil {
		return nil, errors.Join(err, client.ErrRPCCallFailed)
	}

	return &api.Response[uint32]{
		Data:     data,
		Metadata: map[string]any{},
	}, nil
}
