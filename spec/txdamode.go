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

// TxDAMode defines the data availability mode of transaction data.
//
//nolint:recvcheck
type TxDAMode uint32

const (
	// TxDAModeUnknown is an unknown data availability mode.
	TxDAModeUnknown TxDAMode = iota
	// TxDAModeL1 references data available on layer 1.
	TxDAModeL1
	// TxDAModeL2 references data available on layer 2.
	TxDAModeL2
)

var txDAModeStrings = [...]string{
	"UNKNOWN",
	"L1",
	"L2",
}

// MarshalJSON implements json.Marshaler.
func (d TxDAMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", txDAModeStrings[d])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *TxDAMode) UnmarshalJSON(input []byte) error {
	var err error

	switch strings.ToUpper(string(input)) {
	case `"L1"`:
		*d = TxDAModeL1
	case `"L2"`:
		*d = TxDAModeL2
	default:
		err = fmt.Errorf("unrecognised tx data availability mode %s", string(input))
	}

	return err
}

// String returns a string representation of the struct.
func (d TxDAMode) String() string {
	if uint32(d) >= uint32(len(txDAModeStrings)) {
		return txDAModeStrings[0]
	}

	return txDAModeStrings[d]
}
