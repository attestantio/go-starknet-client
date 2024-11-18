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

func TestInvokeV1Transaction(t *testing.T) {
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
			err:   "json: cannot unmarshal array into Go value of type spec.InvokeV1Transaction",
		},
		{
			name:  "Good",
			input: []byte(`{"type":"INVOKE","sender_address":"0x391d69afc1b49f01ad8d2e0e8a03756b694dd056fb6645781eb00f33dbd8caf","calldata":["0x1","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x83afd3f4caedc6eebf44246fe54e38c95e3179a5ec9ea81740eca5b482d12e","0x3","0x714792a41f3651e171c46c93fc53adeb922a414a891cc36d73029d23e99a6ec","0x2386f26fc10000","0x0"],"max_fee":"0x75e126529","version":"0x1","signature":["0x41cb23266d64d2ba4bbbb0a70355e249022153524fb2aeea3c41a2d0d9b785b","0x5f7bab89dcb1eef4e6a09a6da87494759c1816928012d05941c720fa7687920"],"nonce":"0x1"}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.InvokeV1Transaction
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
