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

// ResourceBound constains a specific layer resource bound.
type ResourceBound struct {
	MaxAmount       types.Number `json:"max_amount"`
	MaxPricePerUnit types.Number `json:"max_price_per_unit"`
}

// String returns a string version of the structure.
func (r *ResourceBound) String() string {
	data, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
