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

package api

import (
	"github.com/attestantio/go-starknet-client/spec"
	"github.com/attestantio/go-starknet-client/types"
)

// EstimateFeeOpts are the options for estimating fees for a transaction.
type EstimateFeeOpts struct {
	Common CommonOpts

	// Block is the block at which the fee is estimated.
	// It can be a block number, block hash, or one of the special values "latest" or "pending".
	Block types.BlockID

	// Transaction is the transaction for which to estimate fees.
	Transaction *spec.Transaction

	// Simulation flags?
}
