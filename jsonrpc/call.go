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

// Call makes a call to the client.
func (s *Service) Call(ctx context.Context,
	opts *api.CallOpts,
) (
	*api.Response[[]types.FieldElement],
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

	callOpts := make(map[string]any)
	request := make(map[string]any)
	request["contract_address"] = opts.Contract.String()
	request["entry_point_selector"] = opts.EntryPointSelector.String()

	calldata := make([]string, 0, len(opts.Calldata))
	for i := range opts.Calldata {
		calldata = append(calldata, opts.Calldata[i].String())
	}
	request["calldata"] = calldata
	callOpts["request"] = request
	callOpts["block_id"] = opts.Block

	var callResults []types.FieldElement
	err := s.client.CallFor(&callResults, "starknet_call", callOpts)
	if err != nil {
		return nil, errors.Join(errors.New("starknet_call failed"), err)
	}

	return &api.Response[[]types.FieldElement]{
		Data:     callResults,
		Metadata: map[string]any{},
	}, nil
}
