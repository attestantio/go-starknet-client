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

// ExecutionStatus defines the execution status of a transaction.
type ExecutionStatus uint32

const (
	// ExecutionStatusUnknown is an unknown execution status.
	ExecutionStatusUnknown ExecutionStatus = iota
	// ExecutionStatusSucceeded is a succeeded status.
	ExecutionStatusSucceeded
	// ExecutionStatusReverted is a reverted status.
	ExecutionStatusReverted
)

var executionStatusStrings = [...]string{
	"UNKNOWN",
	"SUCCEEDED",
	"REVERTED",
}

// MarshalJSON implements json.Marshaler.
func (e *ExecutionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", executionStatusStrings[*e])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ExecutionStatus) UnmarshalJSON(input []byte) error {
	var err error
	switch strings.ToUpper(string(input)) {
	case `"SUCCEEDED"`:
		*e = ExecutionStatusSucceeded
	case `"REVERTED"`:
		*e = ExecutionStatusReverted
	default:
		err = fmt.Errorf("unrecognised execution status %s", string(input))
	}

	return err
}

// String returns a string representation of the struct.
func (e ExecutionStatus) String() string {
	if uint32(e) >= uint32(len(executionStatusStrings)) {
		return executionStatusStrings[0]
	}

	return executionStatusStrings[e]
}
