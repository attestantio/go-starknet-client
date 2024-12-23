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

// ChainID returns the chain ID of the node.
func (s *Service) ChainID(ctx context.Context,
	opts *api.ChainIDOpts,
) (
	*api.Response[types.Data],
	error,
) {
	if err := s.assertIsActive(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	var data types.Data
	if err := s.client.CallFor(&data, "starknet_chainId"); err != nil {
		return nil, errors.Join(err, client.ErrRPCCallFailed)
	}

	return &api.Response[types.Data]{
		Data:     data,
		Metadata: map[string]any{},
	}, nil
}
