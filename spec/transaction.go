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
	Type                       TransactionType
	Version                    TransactionVersion
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
type transactionTypeJSON struct {
	Type    TransactionType    `json:"type"`
	Version TransactionVersion `json:"version"`
}

// MarshalJSON marshals a typed transaction.
//
//nolint:gocritic
func (t *Transaction) MarshalJSON() ([]byte, error) {
	switch t.Type {
	case TransactionTypeDeploy:
		switch t.Version {
		case TransactionVersion0:
			return json.Marshal(t.DeployV0Transaction)
		}
	case TransactionTypeInvoke:
		switch t.Version {
		case TransactionVersion0:
			return json.Marshal(t.InvokeV0Transaction)
		case TransactionVersion1:
			return json.Marshal(t.InvokeV1Transaction)
		case TransactionVersion3:
			return json.Marshal(t.InvokeV3Transaction)
		}
	case TransactionTypeDeclare:
		switch t.Version {
		case TransactionVersion0:
			return json.Marshal(t.DeclareV0Transaction)
		case TransactionVersion1:
			return json.Marshal(t.DeclareV1Transaction)
		case TransactionVersion2:
			return json.Marshal(t.DeclareV2Transaction)
		case TransactionVersion3:
			return json.Marshal(t.DeclareV3Transaction)
		}
	case TransactionTypeDeployAccount:
		switch t.Version {
		case TransactionVersion1:
			return json.Marshal(t.DeployAccountV1Transaction)
		case TransactionVersion3:
			return json.Marshal(t.DeployAccountV3Transaction)
		}
	case TransactionTypeL1Handler:
		switch t.Version {
		case TransactionVersion0:
			return json.Marshal(t.L1HandlerV0Transaction)
		}
	}

	return nil, fmt.Errorf("unhandled transaction %v %v", t.Type, t.Version)
}

// UnmarshalJSON implements json.Unmarshaler.
//
//nolint:gocritic
func (t *Transaction) UnmarshalJSON(input []byte) error {
	var data transactionTypeJSON
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
		}
	case TransactionTypeDeployAccount:
		switch data.Version {
		case TransactionVersion1:
			t.DeployAccountV1Transaction = &DeployAccountV1Transaction{}
			err = json.Unmarshal(input, t.DeployAccountV1Transaction)
		case TransactionVersion3:
			t.DeployAccountV3Transaction = &DeployAccountV3Transaction{}
			err = json.Unmarshal(input, t.DeployAccountV3Transaction)
		}
	case TransactionTypeL1Handler:
		switch data.Version {
		case TransactionVersion0:
			t.L1HandlerV0Transaction = &L1HandlerV0Transaction{}
			err = json.Unmarshal(input, t.L1HandlerV0Transaction)
		}
	default:
		err = fmt.Errorf("unhandled transaction %v %v", data.Type, data.Version)
	}

	t.Type = data.Type
	t.Version = data.Version

	return err
}

// // AccessList returns the access list of the transaction.
// // This value can be nil, if the transaction does not support access lists.
// func (t *Transaction) AccessList() []*AccessListEntry {
// 	switch t.Type {
// 	case TransactionType0:
// 		return nil
// 	case TransactionType1:
// 		return t.Type1Transaction.AccessList
// 	case TransactionType2:
// 		return t.Type2Transaction.AccessList
// 	case TransactionType3:
// 		return t.Type3Transaction.AccessList
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // BlobGasUsed returns the blob gas used by the transaction.
// // This value can be nil, if the transaction does not support this (e.g. type 0 transactions).
// func (t *Transaction) BlobGasUsed() *uint32 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return nil
// 	case TransactionType1:
// 		return nil
// 	case TransactionType2:
// 		return nil
// 	case TransactionType3:
// 		return t.Type3Transaction.BlobGasUsed
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // BlobVersionedHashes returns the blob versioned hashes of the transaction.
// // This value can be nil, if the transaction is not a blob transaction.
// func (t *Transaction) BlobVersionedHashes() []types.VersionedHash {
// 	switch t.Type {
// 	case TransactionType0:
// 		return nil
// 	case TransactionType1:
// 		return nil
// 	case TransactionType2:
// 		return nil
// 	case TransactionType3:
// 		return t.Type3Transaction.BlobVersionedHashes
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // BlockHash returns the block hash of the transaction.
// // This value can be nil, if the transaction is not included in a block.
// func (t *Transaction) BlockHash() *types.Hash {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.BlockHash
// 	case TransactionType1:
// 		return t.Type1Transaction.BlockHash
// 	case TransactionType2:
// 		return t.Type2Transaction.BlockHash
// 	case TransactionType3:
// 		return t.Type3Transaction.BlockHash
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // BlockNumber returns the block number of the transaction.
// // This value can be nil, if the transaction is not included in a block.
// func (t *Transaction) BlockNumber() *uint32 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.BlockNumber
// 	case TransactionType1:
// 		return t.Type1Transaction.BlockNumber
// 	case TransactionType2:
// 		return t.Type2Transaction.BlockNumber
// 	case TransactionType3:
// 		return t.Type3Transaction.BlockNumber
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // From returns the sender of the transaction.
// func (t *Transaction) From() types.Address {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.From
// 	case TransactionType1:
// 		return t.Type1Transaction.From
// 	case TransactionType2:
// 		return t.Type2Transaction.From
// 	case TransactionType3:
// 		return t.Type3Transaction.From
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // Gas returns the gas limit of the transaction.
// func (t *Transaction) Gas() uint32 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.Gas
// 	case TransactionType1:
// 		return t.Type1Transaction.Gas
// 	case TransactionType2:
// 		return t.Type2Transaction.Gas
// 	case TransactionType3:
// 		return t.Type3Transaction.Gas
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // GasPrice returns the gas price of the transaction.
// // This will be 0 for transactions that do not have an individual
// // gas price, for example type 2 transactions.
// func (t *Transaction) GasPrice() uint64 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.GasPrice
// 	case TransactionType1:
// 		return t.Type1Transaction.GasPrice
// 	case TransactionType2:
// 		return 0
// 	case TransactionType3:
// 		return 0
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // Hash returns the hash of the transaction.
// func (t *Transaction) Hash() types.Hash {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.Hash
// 	case TransactionType1:
// 		return t.Type1Transaction.Hash
// 	case TransactionType2:
// 		return t.Type2Transaction.Hash
// 	case TransactionType3:
// 		return t.Type3Transaction.Hash
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // Input returns the input data of the transaction.
// func (t *Transaction) Input() []byte {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.Input
// 	case TransactionType1:
// 		return t.Type1Transaction.Input
// 	case TransactionType2:
// 		return t.Type2Transaction.Input
// 	case TransactionType3:
// 		return t.Type3Transaction.Input
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // MaxFeePerGas returns the maximum fee per gas paid by the transaction.
// // This value can be 0, if the transaction does not support this (e.g. type 0 transactions).
// func (t *Transaction) MaxFeePerGas() uint64 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return 0
// 	case TransactionType1:
// 		return 0
// 	case TransactionType2:
// 		return t.Type2Transaction.MaxFeePerGas
// 	case TransactionType3:
// 		return t.Type3Transaction.MaxFeePerGas
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // MaxFeePerBlobGas returns the maximum fee per blob gas paid by the transaction.
// // This value can be 0, if the transaction does not support this (e.g. type 0 transactions).
// func (t *Transaction) MaxFeePerBlobGas() uint64 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return 0
// 	case TransactionType1:
// 		return 0
// 	case TransactionType2:
// 		return 0
// 	case TransactionType3:
// 		return t.Type3Transaction.MaxFeePerBlobGas
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // MaxPriorityFeePerGas returns the maximum priority fee per gas paid by the transaction.
// // This value can be 0, if the transaction does not support this (e.g. type 0 transactions).
// func (t *Transaction) MaxPriorityFeePerGas() uint64 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return 0
// 	case TransactionType1:
// 		return 0
// 	case TransactionType2:
// 		return t.Type2Transaction.MaxPriorityFeePerGas
// 	case TransactionType3:
// 		return t.Type3Transaction.MaxPriorityFeePerGas
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // Nonce returns the nonce of the transaction.
// func (t *Transaction) Nonce() uint64 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.Nonce
// 	case TransactionType1:
// 		return t.Type1Transaction.Nonce
// 	case TransactionType2:
// 		return t.Type2Transaction.Nonce
// 	case TransactionType3:
// 		return t.Type3Transaction.Nonce
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // R returns the R portion of the signature of the transaction.
// func (t *Transaction) R() *big.Int {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.R
// 	case TransactionType1:
// 		return t.Type1Transaction.R
// 	case TransactionType2:
// 		return t.Type2Transaction.R
// 	case TransactionType3:
// 		return t.Type3Transaction.R
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // S returns the S portion of the signature of the transaction.
// func (t *Transaction) S() *big.Int {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.S
// 	case TransactionType1:
// 		return t.Type1Transaction.S
// 	case TransactionType2:
// 		return t.Type2Transaction.S
// 	case TransactionType3:
// 		return t.Type3Transaction.S
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // To returns the recipient of the transaction.
// // This can be nil, for example on contract creation.
// func (t *Transaction) To() *types.Address {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.To
// 	case TransactionType1:
// 		return t.Type1Transaction.To
// 	case TransactionType2:
// 		return t.Type2Transaction.To
// 	case TransactionType3:
// 		return t.Type3Transaction.To
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // TransactionIndex returns the index of the transaction in its block.
// // This value can be nil, if the transaction is not included in a block.
// func (t *Transaction) TransactionIndex() *uint32 {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.TransactionIndex
// 	case TransactionType1:
// 		return t.Type1Transaction.TransactionIndex
// 	case TransactionType2:
// 		return t.Type2Transaction.TransactionIndex
// 	case TransactionType3:
// 		return t.Type3Transaction.TransactionIndex
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // V returns the V portion of the signature of the transaction.
// func (t *Transaction) V() *big.Int {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.V
// 	case TransactionType1:
// 		return t.Type1Transaction.V
// 	case TransactionType2:
// 		return t.Type2Transaction.V
// 	case TransactionType3:
// 		return t.Type3Transaction.V
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }
//
// // Value returns the value of the transaction.
// func (t *Transaction) Value() *big.Int {
// 	switch t.Type {
// 	case TransactionType0:
// 		return t.Type0Transaction.Value
// 	case TransactionType1:
// 		return t.Type1Transaction.Value
// 	case TransactionType2:
// 		return t.Type2Transaction.Value
// 	case TransactionType3:
// 		return t.Type3Transaction.Value
// 	default:
// 		panic(fmt.Errorf("unhandled transaction type %s", t.Type))
// 	}
// }

// String returns a string version of the structure.
func (t *Transaction) String() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(bytes.TrimSuffix(data, []byte("\n")))
}
