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

package jsonrpc

import (
	"context"

	"github.com/attestantio/go-starknet-client/spec"
	"github.com/attestantio/go-starknet-client/types"
)

// preFlightTransaction tidies up a transaction before sending it to a client.
func preFlightTransaction(ctx context.Context,
	tx *spec.Transaction,
) *spec.Transaction {
	switch {
	case tx.InvokeV1Transaction != nil:
		return preFlightInvokeV1Transaction(ctx, tx)
	case tx.InvokeV3Transaction != nil:
		return preFlightInvokeV3Transaction(ctx, tx)
	}

	return tx
}

func preFlightInvokeV1Transaction(_ context.Context,
	tx *spec.Transaction,
) *spec.Transaction {
	cpTx := &spec.Transaction{
		InvokeV1Transaction: tx.InvokeV1Transaction.Copy(),
	}

	if cpTx.InvokeV1Transaction.Signature == nil {
		cpTx.InvokeV1Transaction.Signature = types.Signature{}
	}

	return cpTx
}

func preFlightInvokeV3Transaction(_ context.Context,
	tx *spec.Transaction,
) *spec.Transaction {
	cpTx := &spec.Transaction{
		InvokeV3Transaction: tx.InvokeV3Transaction.Copy(),
	}

	if cpTx.InvokeV3Transaction.Signature == nil {
		cpTx.InvokeV3Transaction.Signature = types.Signature{}
	}
	if cpTx.InvokeV3Transaction.PaymasterData == nil {
		cpTx.InvokeV3Transaction.PaymasterData = []types.FieldElement{}
	}
	if cpTx.InvokeV3Transaction.AccountDeploymentData == nil {
		cpTx.InvokeV3Transaction.AccountDeploymentData = []types.FieldElement{}
	}

	return cpTx
}
