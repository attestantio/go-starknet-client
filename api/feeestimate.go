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
	"encoding/json"
	"fmt"

	"github.com/attestantio/go-starknet-client/types"
)

// FeeEstimate contains a fee estimate.
type FeeEstimate struct {
	GasConsumed     types.Number `json:"gas_consumed"`
	GasPrice        types.Number `json:"gas_price"`
	DataGasConsumed types.Number `json:"data_gas_consumed"`
	DataGasPrice    types.Number `json:"data_gas_price"`
	OverallFee      types.Number `json:"overall_fee"`
	Unit            string       `json:"unit"`
}

// String returns a string version of the structure.
func (f *FeeEstimate) String() string {
	data, err := json.Marshal(f)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
