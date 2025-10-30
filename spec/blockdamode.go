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
	"fmt"
	"strings"
)

// BlockDAMode defines the data availability mode of transaction data.
//
//nolint:recvcheck
type BlockDAMode uint32

const (
	// BlockDAModeUnknown is an unknown data availability mode.
	BlockDAModeUnknown BlockDAMode = iota
	// BlockDAModeL1 references data available on layer 1.
	BlockDAModeL1
	// BlockDAModeL2 references data available on layer 2.
	BlockDAModeL2
)

var blockDAModeStrings = [...]string{
	"UNKNOWN",
	"BLOB",
	"CALLDATA",
}

// MarshalJSON implements json.Marshaler.
func (d BlockDAMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", blockDAModeStrings[d])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *BlockDAMode) UnmarshalJSON(input []byte) error {
	var err error

	switch strings.ToUpper(string(input)) {
	case `"BLOB"`:
		*d = BlockDAModeL1
	case `"CALLDATA"`:
		*d = BlockDAModeL2
	default:
		err = fmt.Errorf("unrecognised block data availability mode %s", string(input))
	}

	return err
}

// String returns a string representation of the struct.
func (d BlockDAMode) String() string {
	if uint32(d) >= uint32(len(blockDAModeStrings)) {
		return blockDAModeStrings[0]
	}

	return blockDAModeStrings[d]
}
