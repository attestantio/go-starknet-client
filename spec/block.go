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

package spec

import (
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-starknet-client/types"
)

// Block contains a block.
type Block struct {
	Status           *FinalityStatus          `json:"status,omitempty"`
	BlockHash        *types.Hash              `json:"block_hash,omitempty"`
	ParentHash       types.Hash               `json:"parent_hash"`
	BlockNumber      *uint64                  `json:"block_number,omitempty"`
	NewRoot          *types.Root              `json:"new_root,omitempty"`
	Timestamp        uint64                   `json:"timestamp"`
	SequencerAddress types.Address            `json:"sequencer_address"`
	L1GasPrice       Price                    `json:"l1_gas_price"`
	L1DataGasPrice   Price                    `json:"l1_data_gas_price"`
	L1DAMode         BlockDAMode              `json:"l1_da_mode"`
	StarknetVersion  string                   `json:"starknet_version"`
	Transactions     []*TransactionAndReceipt `json:"transactions"`
}

// String returns a string version of the structure.
func (t *Block) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
