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

// DeployAccountV1Transaction is version 1 of the deploy account transaction.
type DeployAccountV1Transaction struct {
	TransactionHash     types.Hash           `json:"transaction_hash"`
	Type                TransactionType      `json:"type"`
	MaxFee              types.Number         `json:"max_fee"`
	Version             TransactionVersion   `json:"version"`
	Signature           []types.FieldElement `json:"signature"`
	Nonce               types.Number         `json:"nonce"`
	ContractAddressSalt types.FieldElement   `json:"contract_address_salt"`
	ConstructorCalldata []types.FieldElement `json:"constructor_calldata"`
	ClassHash           types.Hash           `json:"class_hash"`
}

// String returns a string version of the structure.
func (t *DeployAccountV1Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
