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

package mock

import (
	"context"

	"github.com/attestantio/go-starknet-client/api"
	"github.com/attestantio/go-starknet-client/spec"
	"github.com/attestantio/go-starknet-client/types"
)

// Service is a mock starknet client service.
type Service struct{}

// New creates a new mock.
func New() (*Service, error) {
	return &Service{}, nil
}

// Name returns the name of the client implementation.
func (*Service) Name() string { return "mock" }

// Address returns the address of the client.
func (*Service) Address() string { return "mock" }

// BlockHashAndNumber returns the hash and number of the latest block as understood by the node.
func (*Service) BlockHashAndNumber(_ context.Context,
	_ *api.BlockHashAndNumberOpts,
) (
	*api.Response[*api.BlockHashAndNumber],
	error,
) {
	return &api.Response[*api.BlockHashAndNumber]{}, nil
}

// BlockNumber returns the number of the latest block as understood by the node.
func (*Service) BlockNumber(_ context.Context,
	_ *api.BlockNumberOpts,
) (
	*api.Response[uint32],
	error,
) {
	return &api.Response[uint32]{}, nil
}

// Block returns the block as per the given parameters.
func (*Service) Block(_ context.Context,
	_ *api.BlockOpts,
) (
	*api.Response[*spec.Block],
	error,
) {
	return &api.Response[*spec.Block]{}, nil
}

// Call makes a call to the execution client.
func (*Service) Call(_ context.Context,
	_ *api.CallOpts,
) (
	*api.Response[[]types.FieldElement],
	error,
) {
	return &api.Response[[]types.FieldElement]{}, nil
}

// ChainID returns the chain ID.
func (*Service) ChainID(_ context.Context,
	_ *api.ChainIDOpts,
) (
	*api.Response[types.Data],
	error,
) {
	return &api.Response[types.Data]{}, nil
}

// Events returns the events matching the filter.
func (*Service) Events(_ context.Context,
	_ *api.EventsOpts,
) (
	*api.Response[[]*spec.TransactionEvent],
	error,
) {
	return &api.Response[[]*spec.TransactionEvent]{}, nil
}

// Nonce returns the nonce of the given contract at the given block.
func (*Service) Nonce(_ context.Context,
	_ *api.NonceOpts,
) (
	*api.Response[uint32],
	error,
) {
	return &api.Response[uint32]{}, nil
}

// ProtocolVersion returns the protocol version of the node.
func (*Service) ProtocolVersion(_ context.Context,
	_ *api.ProtocolVersionOpts,
) (
	*api.Response[uint32],
	error,
) {
	return &api.Response[uint32]{}, nil
}

// SpecVersion returns the version of the specification followed by the node.
func (*Service) SpecVersion(_ context.Context,
	_ *api.SpecVersionOpts,
) (
	*api.Response[string],
	error,
) {
	return &api.Response[string]{}, nil
}

// Syncing obtains information about the sync state of the node.
func (*Service) Syncing(_ context.Context,
	_ *api.SyncingOpts,
) (
	*api.Response[*api.SyncState],
	error,
) {
	return &api.Response[*api.SyncState]{}, nil
}
