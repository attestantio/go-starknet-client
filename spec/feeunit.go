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

// FeeUnit defines the type of a transaction.
type FeeUnit uint32

const (
	// FeeUnitUnknown is an unknown fee unit.
	FeeUnitUnknown FeeUnit = iota
	// FeeUnitWei is the Wei unit.
	FeeUnitWei
	// FeeUnitFri is the Friday unit.
	FeeUnitFri
)

var feeUnitStrings = [...]string{
	"unknown",
	"WEI",
	"FRI",
}

// MarshalJSON implements json.Marshaler.
func (f *FeeUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", feeUnitStrings[*f])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *FeeUnit) UnmarshalJSON(input []byte) error {
	var err error
	switch strings.ToUpper(string(input)) {
	case `"WEI"`:
		*f = FeeUnitWei
	case `"FRI"`:
		*f = FeeUnitFri
	default:
		err = fmt.Errorf("unrecognised fee unit %s", string(input))
	}

	return err
}

// String returns a string representation of the struct.
func (f FeeUnit) String() string {
	if uint32(f) >= uint32(len(feeUnitStrings)) {
		return "unknown"
	}

	return feeUnitStrings[f]
}
