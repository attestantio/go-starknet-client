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

// TransactionEvent contains an event emitted by a transaction.
type TransactionEvent struct {
	BlockHash       *types.Hash          `json:"block_hash,omitempty"`
	BlockNumber     *uint32              `json:"block_number,omitempty"`
	FromAddress     types.Address        `json:"from_address"`
	Keys            []types.FieldElement `json:"keys"`
	Data            []types.FieldElement `json:"data"`
	TransactionHash *types.Hash          `json:"transaction_hash,omitempty"`
}

// String returns a string version of the structure.
func (e *TransactionEvent) String() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
