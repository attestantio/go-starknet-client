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
	"errors"
	"fmt"
	"strings"
)

// TransactionVersion defines the spec version of a transaction.
//
//nolint:recvcheck
type TransactionVersion uint8

const (
	// TransactionVersionUnknown is an unknown version.
	TransactionVersionUnknown TransactionVersion = iota
	// TransactionVersion0 is a version 0 transaction.
	TransactionVersion0
	// TransactionVersion0Query is a query-only version 0 transaction.
	TransactionVersion0Query
	// TransactionVersion1 is a version 1 transaction.
	TransactionVersion1
	// TransactionVersion1Query is a query-only version 1 transaction.
	TransactionVersion1Query
	// TransactionVersion2 is a version 2 transaction.
	TransactionVersion2
	// TransactionVersion2Query is a query-only version 2 transaction.
	TransactionVersion2Query
	// TransactionVersion3 is a version 3 transaction.
	TransactionVersion3
	// TransactionVersion3Query is a query-only version 3 transaction.
	TransactionVersion3Query
)

var transactionVersionStrings = [...]string{
	"UNKNOWN",
	"0x0",
	"0x100000000000000000000000000000000",
	"0x1",
	"0x100000000000000000000000000000001",
	"0x2",
	"0x100000000000000000000000000000002",
	"0x3",
	"0x100000000000000000000000000000003",
}

// MarshalJSON implements json.Marshaler.
func (v TransactionVersion) MarshalJSON() ([]byte, error) {
	if int(v) >= len(transactionVersionStrings) {
		return nil, errors.New("invalid transaction version")
	}

	return []byte(fmt.Sprintf("%q", transactionVersionStrings[v])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *TransactionVersion) UnmarshalJSON(input []byte) error {
	var err error
	switch strings.ToLower(strings.Trim(string(input), `"`)) {
	case "unknown":
		*v = TransactionVersionUnknown
	case "0x0":
		*v = TransactionVersion0
	case "0x100000000000000000000000000000000":
		*v = TransactionVersion0Query
	case "0x1":
		*v = TransactionVersion1
	case "0x100000000000000000000000000000001":
		*v = TransactionVersion1Query
	case "0x2":
		*v = TransactionVersion2
	case "0x100000000000000000000000000000002":
		*v = TransactionVersion2Query
	case "0x3":
		*v = TransactionVersion3
	case "0x100000000000000000000000000000003":
		*v = TransactionVersion3Query
	default:
		err = fmt.Errorf("unrecognised transaction version %s", string(input))
	}

	return err
}

// String returns a string representation of the item.
func (v TransactionVersion) String() string {
	if int(v) >= len(transactionVersionStrings) {
		return transactionVersionStrings[0]
	}

	return transactionVersionStrings[v]
}
