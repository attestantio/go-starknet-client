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

// TransactionReceipt contains execution results of a transaction.
type TransactionReceipt struct {
	Type               TransactionType     `json:"type"`
	TransactionHash    types.Hash          `json:"transaction_hash"`
	ActualFee          Fee                 `json:"actual_fee"`
	ExecutionStatus    ExecutionStatus     `json:"execution_status"`
	FinalityStatus     FinalityStatus      `json:"finality_status"`
	BlockHash          *types.Hash         `json:"block_hash,omitempty"`
	BlockNumber        *uint64             `json:"block_number,omitempty"`
	MessagesSent       []*Message          `json:"messages_sent"`
	RevertReason       string              `json:"revert_reason,omitempty"`
	Events             []*TransactionEvent `json:"events"`
	ContractAddress    *types.Address      `json:"contract_address,omitempty"`
	ExecutionResources ExecutionResources  `json:"execution_resources"`
}

// String returns a string version of the structure.
func (r *TransactionReceipt) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
