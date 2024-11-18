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

package jsonrpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/attestantio/go-starknet-client/api"
	"github.com/attestantio/go-starknet-client/jsonrpc"
	"github.com/attestantio/go-starknet-client/spec"
	"github.com/attestantio/go-starknet-client/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestSubmitTransaction(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	tests := []struct {
		name        string
		transaction *spec.Transaction
	}{
		{
			name: "InvokeV1",
			transaction: &spec.Transaction{
				InvokeV1Transaction: &spec.InvokeV1Transaction{
					Type:          spec.TransactionTypeInvoke,
					SenderAddress: strToAddress("0x391d69afc1b49f01ad8d2e0e8a03756b694dd056fb6645781eb00f33dbd8caf"),
					Calldata: []types.FieldElement{
						strToFieldElement("0x1"),
						strToFieldElement("0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"),
						strToFieldElement("0x83afd3f4caedc6eebf44246fe54e38c95e3179a5ec9ea81740eca5b482d12e"),
						strToFieldElement("0x3"),
						strToFieldElement("0x714792a41f3651e171c46c93fc53adeb922a414a891cc36d73029d23e99a6ec"),
						strToFieldElement("0x2386f26fc10000"),
						strToFieldElement("0x0"),
					},
					Version: spec.TransactionVersion1,
					Signature: types.Signature{
						strToFieldElement("0x5cc5059ca03b3270080de3d87e9c05a99a140c47875cbb1c43d283a0098be13"),
						strToFieldElement("0x7bc7316ebada49485c49c7a8b5cf0bc756486888f1fc2cb60acdadf0f54afc0"),
					},
					Nonce:  3,
					MaxFee: 1000000000000000000,
				},
			},
		},
		{
			name: "InvokeV3",
			transaction: &spec.Transaction{
				InvokeV3Transaction: &spec.InvokeV3Transaction{
					Type:          spec.TransactionTypeInvoke,
					SenderAddress: strToAddress("0x391d69afc1b49f01ad8d2e0e8a03756b694dd056fb6645781eb00f33dbd8caf"),
					Calldata: []types.FieldElement{
						strToFieldElement("0x1"),
						strToFieldElement("0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"),
						strToFieldElement("0x83afd3f4caedc6eebf44246fe54e38c95e3179a5ec9ea81740eca5b482d12e"),
						strToFieldElement("0x3"),
						strToFieldElement("0x714792a41f3651e171c46c93fc53adeb922a414a891cc36d73029d23e99a6ec"),
						strToFieldElement("0x2386f26fc10000"),
						strToFieldElement("0x0"),
					},
					Version: spec.TransactionVersion3,
					Signature: types.Signature{
						strToFieldElement("0x5cc5059ca03b3270080de3d87e9c05a99a140c47875cbb1c43d283a0098be13"),
						strToFieldElement("0x7bc7316ebada49485c49c7a8b5cf0bc756486888f1fc2cb60acdadf0f54afc0"),
					},
					Nonce: 3,
					ResourceBounds: spec.ResourceBounds{
						L1Gas: spec.ResourceBound{
							MaxAmount:       1000,
							MaxPricePerUnit: 1000000000000000,
						},
					},
					Tip:                       0,
					NonceDataAvailabilityMode: spec.TxDAModeL1,
					FeeDataAvailabilityMode:   spec.TxDAModeL1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := s.SubmitTransaction(ctx, &api.SubmitTransactionOpts{
				Transaction: test.transaction,
			})
			require.NoError(t, err)
			require.NotNil(t, response.Data)
			require.False(t, response.Data.TransactionHash.IsZero())
		})
	}
}
