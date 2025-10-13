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

// FinalityStatus defines the finality status of a transaction.
//
//nolint:recvcheck
type FinalityStatus uint32

const (
	// FinalityStatusUnknown is an unknown finality status.
	FinalityStatusUnknown FinalityStatus = iota
	// FinalityStatusPending means the transaction is pending acceptance.
	FinalityStatusPending
	// FinalityStatusAcceptedOnL2 means the transaction has been accepted on layer 2.
	FinalityStatusAcceptedOnL2
	// FinalityStatusAcceptedOnL1 means the transaction has been accepted on layer 1.
	FinalityStatusAcceptedOnL1
	// FinalityStatusRejected means the transaction has been rejected.
	FinalityStatusRejected
)

var finalityStatusStrings = [...]string{
	"UNKNOWN",
	"PENDING",
	"ACCEPTED_ON_L2",
	"ACCEPTED_ON_L1",
	"REJECTED",
}

// MarshalJSON implements json.Marshaler.
func (f FinalityStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", finalityStatusStrings[f])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *FinalityStatus) UnmarshalJSON(input []byte) error {
	var err error

	switch strings.ToUpper(string(input)) {
	case `"PENDING"`:
		*f = FinalityStatusPending
	case `"ACCEPTED_ON_L2"`:
		*f = FinalityStatusAcceptedOnL2
	case `"ACCEPTED_ON_L1"`:
		*f = FinalityStatusAcceptedOnL1
	case `"REJECTED"`:
		*f = FinalityStatusRejected
	default:
		err = fmt.Errorf("unrecognised finality status %s", string(input))
	}

	return err
}

// String returns a string representation of the struct.
func (f FinalityStatus) String() string {
	if uint32(f) >= uint32(len(finalityStatusStrings)) {
		return finalityStatusStrings[0]
	}

	return finalityStatusStrings[f]
}
