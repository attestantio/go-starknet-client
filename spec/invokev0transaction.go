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

// InvokeV0Transaction is version 0 of the invoke transaction.
type InvokeV0Transaction struct {
	TransactionHash    *types.Hash          `json:"transaction_hash,omitempty"`
	Type               TransactionType      `json:"type"`
	MaxFee             types.Number         `json:"max_fee"`
	Version            TransactionVersion   `json:"version"`
	Signature          types.Signature      `json:"signature"`
	ContractAddress    types.Address        `json:"contract_address"`
	EntryPointSelector types.FieldElement   `json:"entry_point_selector"`
	Calldata           []types.FieldElement `json:"calldata"`
}

// String returns a string version of the structure.
func (t *InvokeV0Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
