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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// Transaction is a struct that covers all transaction types.
type Transaction struct {
	DeployV0Transaction        *DeployV0Transaction
	InvokeV0Transaction        *InvokeV0Transaction
	InvokeV1Transaction        *InvokeV1Transaction
	InvokeV3Transaction        *InvokeV3Transaction
	DeclareV0Transaction       *DeclareV0Transaction
	DeclareV1Transaction       *DeclareV1Transaction
	DeclareV2Transaction       *DeclareV2Transaction
	DeclareV3Transaction       *DeclareV3Transaction
	DeployAccountV1Transaction *DeployAccountV1Transaction
	DeployAccountV3Transaction *DeployAccountV3Transaction
	L1HandlerV0Transaction     *L1HandlerV0Transaction
}

// transactionTypeAndVersionJSON is a simple struct to fetch the transaction type and version.
type transactionTypeAndVersionJSON struct {
	Type    TransactionType    `json:"type"`
	Version TransactionVersion `json:"version"`
}

// MarshalJSON marshals a typed transaction.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	switch {
	case t.DeployV0Transaction != nil:
		return json.Marshal(t.DeployV0Transaction)
	case t.InvokeV0Transaction != nil:
		return json.Marshal(t.InvokeV0Transaction)
	case t.InvokeV1Transaction != nil:
		return json.Marshal(t.InvokeV1Transaction)
	case t.InvokeV3Transaction != nil:
		return json.Marshal(t.InvokeV3Transaction)
	case t.DeclareV0Transaction != nil:
		return json.Marshal(t.DeclareV0Transaction)
	case t.DeclareV1Transaction != nil:
		return json.Marshal(t.DeclareV1Transaction)
	case t.DeclareV2Transaction != nil:
		return json.Marshal(t.DeclareV2Transaction)
	case t.DeclareV3Transaction != nil:
		return json.Marshal(t.DeclareV3Transaction)
	case t.DeployAccountV1Transaction != nil:
		return json.Marshal(t.DeployAccountV1Transaction)
	case t.DeployAccountV3Transaction != nil:
		return json.Marshal(t.DeployAccountV3Transaction)
	case t.L1HandlerV0Transaction != nil:
		return json.Marshal(t.L1HandlerV0Transaction)
	default:
		return nil, errors.New("unhandled transaction")
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	var data transactionTypeAndVersionJSON

	err := json.Unmarshal(input, &data)
	if err != nil {
		return errors.Join(errors.New("invalid JSON"), err)
	}

	switch data.Type {
	case TransactionTypeDeploy:
		switch data.Version {
		case TransactionVersion0:
			t.DeployV0Transaction = &DeployV0Transaction{}
			err = json.Unmarshal(input, t.DeployV0Transaction)
		default:
			return fmt.Errorf("unsupported deploy transaction version: %s", data.Version)
		}
	case TransactionTypeInvoke:
		switch data.Version {
		case TransactionVersion0:
			t.InvokeV0Transaction = &InvokeV0Transaction{}
			err = json.Unmarshal(input, t.InvokeV0Transaction)
		case TransactionVersion1:
			t.InvokeV1Transaction = &InvokeV1Transaction{}
			err = json.Unmarshal(input, t.InvokeV1Transaction)
		case TransactionVersion3:
			t.InvokeV3Transaction = &InvokeV3Transaction{}
			err = json.Unmarshal(input, t.InvokeV3Transaction)
		default:
			return fmt.Errorf("unsupported invoke transaction version: %s", data.Version)
		}
	case TransactionTypeDeclare:
		switch data.Version {
		case TransactionVersion0:
			t.DeclareV0Transaction = &DeclareV0Transaction{}
			err = json.Unmarshal(input, t.DeclareV0Transaction)
		case TransactionVersion1:
			t.DeclareV1Transaction = &DeclareV1Transaction{}
			err = json.Unmarshal(input, t.DeclareV1Transaction)
		case TransactionVersion2:
			t.DeclareV2Transaction = &DeclareV2Transaction{}
			err = json.Unmarshal(input, t.DeclareV2Transaction)
		case TransactionVersion3:
			t.DeclareV3Transaction = &DeclareV3Transaction{}
			err = json.Unmarshal(input, t.DeclareV3Transaction)
		default:
			return fmt.Errorf("unsupported declare transaction version: %s", data.Version)
		}
	case TransactionTypeDeployAccount:
		switch data.Version {
		case TransactionVersion1:
			t.DeployAccountV1Transaction = &DeployAccountV1Transaction{}
			err = json.Unmarshal(input, t.DeployAccountV1Transaction)
		case TransactionVersion3:
			t.DeployAccountV3Transaction = &DeployAccountV3Transaction{}
			err = json.Unmarshal(input, t.DeployAccountV3Transaction)
		default:
			return fmt.Errorf("unsupported deploy account transaction version: %s", data.Version)
		}
	case TransactionTypeL1Handler:
		switch data.Version {
		case TransactionVersion0:
			t.L1HandlerV0Transaction = &L1HandlerV0Transaction{}
			err = json.Unmarshal(input, t.L1HandlerV0Transaction)
		default:
			return fmt.Errorf("unsupported L1 handler transaction version: %s", data.Version)
		}
	default:
		err = fmt.Errorf("unhandled transaction %v %v", data.Type, data.Version)
	}

	return err
}

// SetQueryBit sets the query bit for the transaction.
func (t *Transaction) SetQueryBit() {
	switch {
	case t.DeployV0Transaction != nil:
		t.DeployV0Transaction.Version = TransactionVersion0Query
	case t.InvokeV0Transaction != nil:
		t.InvokeV0Transaction.Version = TransactionVersion0Query
	case t.InvokeV1Transaction != nil:
		t.InvokeV1Transaction.Version = TransactionVersion1Query
	case t.InvokeV3Transaction != nil:
		t.InvokeV3Transaction.Version = TransactionVersion3Query
	case t.DeclareV0Transaction != nil:
		t.DeclareV0Transaction.Version = TransactionVersion0Query
	case t.DeclareV1Transaction != nil:
		t.DeclareV1Transaction.Version = TransactionVersion1Query
	case t.DeclareV2Transaction != nil:
		t.DeclareV2Transaction.Version = TransactionVersion2Query
	case t.DeclareV3Transaction != nil:
		t.DeclareV3Transaction.Version = TransactionVersion3Query
	case t.DeployAccountV1Transaction != nil:
		t.DeployAccountV1Transaction.Version = TransactionVersion1Query
	case t.DeployAccountV3Transaction != nil:
		t.DeployAccountV3Transaction.Version = TransactionVersion3Query
	case t.L1HandlerV0Transaction != nil:
		t.L1HandlerV0Transaction.Version = TransactionVersion1Query
	default:
		// No transaction type matched - this shouldn't happen with valid data
	}
}

// String returns a string version of the structure.
func (t *Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(bytes.TrimSuffix(data, []byte("\n")))
}
