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

// SubmitTransaction submits a transaction to the client.
func (s *Service) SubmitTransaction(ctx context.Context,
	opts *api.SubmitTransactionOpts,
) (
	*api.Response[*api.SubmitTransactionResponse],
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

	opts.Transaction = preFlightTransaction(ctx, opts.Transaction)

	switch {
	case opts.Transaction.InvokeV1Transaction != nil:
		return s.invokeV1Transaction(ctx, opts)
	case opts.Transaction.InvokeV3Transaction != nil:
		return s.invokeV3Transaction(ctx, opts)
	default:
		return nil, errors.New("unhandled transaction type")
	}
}

func (s *Service) invokeV1Transaction(_ context.Context,
	opts *api.SubmitTransactionOpts,
) (
	*api.Response[*api.SubmitTransactionResponse],
	error,
) {
	var data api.SubmitTransactionResponse
	err := s.client.CallFor(&data, "starknet_addInvokeTransaction", []*spec.Transaction{opts.Transaction})
	if err != nil {
		return nil, errors.Join(errors.New("starknet_call failed"), err)
	}

	return &api.Response[*api.SubmitTransactionResponse]{
		Data:     &data,
		Metadata: map[string]any{},
	}, nil
}

func (s *Service) invokeV3Transaction(_ context.Context,
	opts *api.SubmitTransactionOpts,
) (
	*api.Response[*api.SubmitTransactionResponse],
	error,
) {
	var data api.SubmitTransactionResponse
	err := s.client.CallFor(&data, "starknet_addInvokeTransaction", []*spec.Transaction{opts.Transaction})
	if err != nil {
		return nil, errors.Join(errors.New("starknet_call failed"), err)
	}

	return &api.Response[*api.SubmitTransactionResponse]{
		Data:     &data,
		Metadata: map[string]any{},
	}, nil
}
