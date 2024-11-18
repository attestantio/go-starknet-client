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

package spec_test

import (
	"encoding/json"
	"testing"

	"github.com/attestantio/go-starknet-client/spec"
	"github.com/stretchr/testify/require"
)

func TestInvokeV3Transaction(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
		err      string
	}{
		{
			name: "Empty",
			err:  "unexpected end of JSON input",
		},
		{
			name:  "JSONBad",
			input: []byte("[]"),
			err:   "json: cannot unmarshal array into Go value of type spec.InvokeV3Transaction",
		},
		{
			name:  "Good",
			input: []byte(`{"type":"INVOKE","sender_address":"0x391d69afc1b49f01ad8d2e0e8a03756b694dd056fb6645781eb00f33dbd8caf","calldata":["0x1","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x83afd3f4caedc6eebf44246fe54e38c95e3179a5ec9ea81740eca5b482d12e","0x3","0x714792a41f3651e171c46c93fc53adeb922a414a891cc36d73029d23e99a6ec","0x2386f26fc10000","0x0"],"version":"0x3","signature":["0x30793b441461b3627061238a8370037cfe4ed310d1df5c74b6302ab2f1ee1ce","0x2b08553d33515ba29e8235e163718731a2c4ae55a10b036a848485c9601d236"],"nonce":"0x2","resource_bounds":{"l1_gas":{"max_amount":"0x22","max_price_per_unit":"0x9819642cfe1"},"l2_gas":{"max_amount":"0x0","max_price_per_unit":"0x0"}},"tip":"0x0","paymaster_data":[],"account_deployment_data":[],"nonce_data_availability_mode":"L1","fee_data_availability_mode":"L1"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.InvokeV3Transaction
			err := json.Unmarshal(test.input, &res)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				rt, err := json.Marshal(&res)
				require.NoError(t, err)
				if len(test.expected) == 0 {
					require.JSONEq(t, string(test.input), string(rt))
				} else {
					require.JSONEq(t, string(test.expected), string(rt))
				}
			}
		})
	}
}
