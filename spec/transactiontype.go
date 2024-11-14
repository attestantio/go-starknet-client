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

// TransactionType defines the type of a transaction.
//
//nolint:recvcheck
type TransactionType uint8

const (
	// TransactionTypeUnknown is an unknown transaction type.
	TransactionTypeUnknown TransactionType = iota
	// TransactionTypeDeploy is a deploy transaction.
	TransactionTypeDeploy
	// TransactionTypeInvoke is an invoke transaction.
	TransactionTypeInvoke
	// TransactionTypeDeclare is a declare transaction.
	TransactionTypeDeclare
	// TransactionTypeDeployAccount is a deploy account transaction.
	TransactionTypeDeployAccount
	// TransactionTypeL1Handler is a layer 1 handler transaction.
	TransactionTypeL1Handler
)

var transactionTypeStrings = [...]string{
	"UNKNOWN",
	"DEPLOY",
	"INVOKE",
	"DECLARE",
	"DEPLOY_ACCOUNT",
	"L1_HANDLER",
}

// MarshalJSON implements json.Marshaler.
func (t TransactionType) MarshalJSON() ([]byte, error) {
	if int(t) >= len(transactionVersionStrings) {
		return nil, errors.New("invalid transaction version")
	}

	return []byte(fmt.Sprintf("%q", transactionTypeStrings[t])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TransactionType) UnmarshalJSON(input []byte) error {
	var err error
	switch strings.ToUpper(strings.Trim(string(input), `"`)) {
	case "UNKNOWN":
		*t = TransactionTypeUnknown
	case "DEPLOY":
		*t = TransactionTypeDeploy
	case "INVOKE":
		*t = TransactionTypeInvoke
	case "DECLARE":
		*t = TransactionTypeDeclare
	case "DEPLOY_ACCOUNT":
		*t = TransactionTypeDeployAccount
	case "L1_HANDLER":
		*t = TransactionTypeL1Handler
	default:
		err = fmt.Errorf("unrecognised transaction type %s", string(input))
	}

	return err
}

// String returns a string representation of the item.
func (t TransactionType) String() string {
	if int(t) >= len(transactionTypeStrings) {
		return transactionTypeStrings[0]
	}

	return transactionTypeStrings[t]
}
