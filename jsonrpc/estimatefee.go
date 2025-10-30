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

// EstimateFee estimates the fee for a transaction.
func (s *Service) EstimateFee(ctx context.Context,
	opts *api.EstimateFeeOpts,
) (
	*api.Response[[]api.FeeEstimate],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	if opts.Transaction == nil {
		return nil, errors.Join(errors.New("no transaction specified"), client.ErrInvalidOptions)
	}

	tx := preFlightTransaction(ctx, opts.Transaction)
	tx.SetQueryBit()

	var data []api.FeeEstimate

	err := s.client.CallFor(&data, "starknet_estimateFee", []any{tx}, []any{"SKIP_VALIDATE"}, opts.Block)
	if err != nil {
		return nil, errors.Join(errors.New("starknet_estimateFee failed"), err)
	}

	return &api.Response[[]api.FeeEstimate]{
		Data:     data,
		Metadata: map[string]any{},
	}, nil
}
