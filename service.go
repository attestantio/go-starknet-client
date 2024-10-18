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

package client

import (
	"context"

	"github.com/attestantio/go-starknet-client/api"
	"github.com/attestantio/go-starknet-client/spec"
	"github.com/attestantio/go-starknet-client/types"
)

// Service is the service providing a connection to a Starknet node.
type Service interface {
	// Name returns the name of the node implementation.
	Name() string

	// Address returns the address of the node.
	Address() string
}

// BlockHashAndNumberProvider is the interface for providing block hashes and numbers.
type BlockHashAndNumberProvider interface {
	// BlockHashAndNumber returns the hash and number of the latest block as understood by the node.
	BlockHashAndNumber(ctx context.Context,
		opts *api.BlockHashAndNumberOpts,
	) (
		*api.Response[*api.BlockHashAndNumber],
		error,
	)
}

// BlockNumberProvider is the interface for providing block numbers.
type BlockNumberProvider interface {
	// BlockNumber returns the number of the latest block as understood by the node.
	BlockNumber(ctx context.Context,
		opts *api.BlockNumberOpts,
	) (
		*api.Response[uint32],
		error,
	)
}

// BlockProvider is the interface for providing blocks.
type BlockProvider interface {
	// Block returns the block as per the given parameters.
	Block(ctx context.Context,
		opts *api.BlockOpts,
	) (
		*api.Response[*spec.Block],
		error,
	)
}

// CallProvider is the interface for making calls to the execution client.
type CallProvider interface {
	// Call makes a call to the execution client.
	Call(ctx context.Context,
		opts *api.CallOpts,
	) (
		*api.Response[[]types.FieldElement],
		error,
	)
}

// ChainIDProvider is the interface for providing the chain ID.
type ChainIDProvider interface {
	// ChainID returns the chain ID.
	ChainID(ctx context.Context, opts *api.ChainIDOpts) (*api.Response[types.Data], error)
}

// EventsProvider is the interface for providing events.
type EventsProvider interface {
	// Events returns the events matching the filter.
	Events(ctx context.Context, opts *api.EventsOpts) (*api.Response[[]*spec.TransactionEvent], error)
}

// NonceProvider is the interface for providing contract nonces.
type NonceProvider interface {
	// Nonce returns the nonce of the given contract at the given block.
	Nonce(ctx context.Context,
		opts *api.NonceOpts,
	) (
		*api.Response[uint32],
		error,
	)
}

// ProtocolVersionProvider is the interface for providing the protocol version of the node.
type ProtocolVersionProvider interface {
	// ProtocolVersion returns the protocol version of the node.
	ProtocolVersion(ctx context.Context,
		opts *api.ProtocolVersionOpts,
	) (
		*api.Response[uint32],
		error,
	)
}

// SpecVersionProvider is the interface for providing specifiation version information.
type SpecVersionProvider interface {
	// SpecVersion returns the version of the specification followed by the node.
	SpecVersion(ctx context.Context,
		opts *api.SpecVersionOpts,
	) (
		*api.Response[string],
		error,
	)
}

// SyncingProvider is the interface for providing syncing information.
type SyncingProvider interface {
	// Syncing obtains information about the sync state of the node.
	Syncing(ctx context.Context,
		opts *api.SyncingOpts,
	) (
		*api.Response[*api.SyncState],
		error,
	)
}
