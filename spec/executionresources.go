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
)

// ExecutionResources contains information about the resources used in executing a transaction.
type ExecutionResources struct {
	Steps                         uint64            `json:"steps"`
	MemoryHoles                   uint64            `json:"memory_holes,omitempty"`
	PedersenBuiltinApplications   uint64            `json:"pedersen_builtin_applications,omitempty"`
	RangeCheckBuiltinApplications uint64            `json:"range_check_builtin_applications,omitempty"`
	BitwiseBuiltinApplications    uint64            `json:"bitwise_builtin_applications,omitempty"`
	EcOpBuiltinApplications       uint64            `json:"ec_op_builtin_applications,omitempty"`
	PoseidonBuiltinApplications   uint64            `json:"poseidon_builtin_applications,omitempty"`
	EcdsaBuiltinApplications      uint64            `json:"ecdsa_builtin_applications,omitempty"`
	KeccakBuiltinApplications     uint64            `json:"keccak_builtin_applications,omitempty"`
	SegmentArenaBuiltin           uint64            `json:"segment_arena_builtin,omitempty"`
	DataAvailability              *DataAvailability `json:"data_availability,omitempty"`
}

// String returns a string version of the structure.
func (e *ExecutionResources) String() string {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
