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

// InvokeV3Transaction is version 3 of the invoke transaction.
type InvokeV3Transaction struct {
	TransactionHash           types.Hash           `json:"transaction_hash"`
	Type                      TransactionType      `json:"type"`
	SenderAddress             types.Address        `json:"sender_address"`
	Calldata                  []types.FieldElement `json:"calldata"`
	Version                   TransactionVersion   `json:"version"`
	Signature                 []types.FieldElement `json:"signature"`
	Nonce                     types.Number         `json:"nonce"`
	ResourceBounds            ResourceBounds       `json:"resource_bounds"`
	Tip                       types.Number         `json:"tip"`
	PaymasterData             []types.FieldElement `json:"paymaster_data"`
	AccountDeploymentData     []types.FieldElement `json:"account_deployment_data"`
	NonceDataAvailabilityMode TxDAMode             `json:"nonce_data_availability_mode"`
	FeeDataAvailabilityMode   TxDAMode             `json:"fee_data_availability_mode"`
}

// String returns a string version of the structure.
func (t *InvokeV3Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
