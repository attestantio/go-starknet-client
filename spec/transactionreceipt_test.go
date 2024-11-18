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

func TestTransactionReceipt(t *testing.T) {
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
			err:   "json: cannot unmarshal array into Go value of type spec.TransactionReceipt",
		},
		{
			name:  "Full",
			input: []byte(`{"type":"INVOKE","transaction_hash":"0x4eea7ea635466c1253e9a3e1aeee6a85e29e8020155399711d90e45c6deaf6a","actual_fee":{"amount":"0x118dcc7fe511833","unit":"FRI"},"execution_status":"SUCCEEDED","finality_status":"ACCEPTED_ON_L2","messages_sent":[],"events":[{"from_address":"0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","keys":["0x96982abd597114bdaa4a60612f87fabfcc7206aa12d61c50e7ba1e6c291100"],"data":["0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x0","0xd0","0x0","0xb8284","0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x84c140","0x1","0x84c070","0x1","0x0","0x1","0x0","0x1"]},{"from_address":"0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","keys":["0x5dacf59794364ad1555bb3c9b2346afa81e57e5c19bb6bae0d22721c96c4e5"],"data":["0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x0","0xd0","0x0","0xb8284","0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x84c140","0x1","0x84c070","0x1","0x0","0x1","0x0","0x1"]},{"from_address":"0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","keys":["0x3a7adca3546c213ce791fabf3b04090c163e419c808c9830fb343a4a395946e"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x0","0xd0","0x0","0xb8284","0x84c140","0x1","0x84c070","0x1","0x12eb365618b4bcc4b9e1d1","0x1","0x27071607066b3642d1f1","0x1","0x0","0x0"]},{"from_address":"0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","keys":["0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9"],"data":["0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","0x207c098face5cbb03f3a3a17f765ac54f8706730dd573a8e73c7496722a84ce","0x27071607066b3642d1f1","0x0"]},{"from_address":"0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","keys":["0x157717768aca88da4ac4279765f09f4d0151823d573537fbbeb950cdbd9a870"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x0","0xd0","0x0","0x0","0x0","0x0","0x1000003f7f1380b75","0x0","0x0","0x0","0x0","0x0","0x1","0x325de9a14c148956f5e383fe7b1a00c","0x0","0x863ac7","0x1","0x0"]},{"from_address":"0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","keys":["0x3a7adca3546c213ce791fabf3b04090c163e419c808c9830fb343a4a395946e"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x0","0xd0","0x0","0x0","0x549cd80","0x1","0x549cd80","0x0","0x0","0x0","0x0","0x0","0x0","0x0"]},{"from_address":"0x7b696af58c967c1b14c9dde0ace001720635a660a8e90c565ea459345318b30","keys":["0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9"],"data":["0x207c098face5cbb03f3a3a17f765ac54f8706730dd573a8e73c7496722a84ce","0x64691fcf3e0421406c8ab1a8028bc05affbb16aecba4f0352f8d3d7a3386212","0xb8284","0x0"]},{"from_address":"0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","keys":["0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9"],"data":["0x207c098face5cbb03f3a3a17f765ac54f8706730dd573a8e73c7496722a84ce","0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x27071607066b3642d1f1","0x0"]},{"from_address":"0x7b696af58c967c1b14c9dde0ace001720635a660a8e90c565ea459345318b30","keys":["0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9"],"data":["0x0","0x207c098face5cbb03f3a3a17f765ac54f8706730dd573a8e73c7496722a84ce","0xb828c","0x0"]},{"from_address":"0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","keys":["0x3a7adca3546c213ce791fabf3b04090c163e419c808c9830fb343a4a395946e"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","0x0","0xd0","0x0","0xb828c","0x841f90","0x1","0x841ec0","0x1","0x13507de31a01b130234c2b","0x0","0x27071607066b3642d1f1","0x0","0x0","0x0"]},{"from_address":"0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","keys":["0x134692b230b9e1ffa39098904722134159652b09c5bc41d88d6698779d228ff"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","0x27071607066b3642d1f1","0x0"]},{"from_address":"0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","keys":["0x134692b230b9e1ffa39098904722134159652b09c5bc41d88d6698779d228ff"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","0x0","0x0"]},{"from_address":"0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","keys":["0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9"],"data":["0x2e0af29598b407c8716b17f6d2795eca1b471413fa03fb145a5e33722184067","0x5dd3d2f4429af886cd1a3b08289dbcea99a294197e9eb43b0e0325b4b","0x27071607066b3642d1f1","0x0"]},{"from_address":"0x7461e8a41459e52d5ca62e6faee68c8149d3d8e974ca120ed8cb752192b3f5a","keys":["0x1dcde06aabdbca2f80aa51392b345d7549d7757aa855f7e37f5d335ac8243b1","0x4eea7ea635466c1253e9a3e1aeee6a85e29e8020155399711d90e45c6deaf6a"],"data":["0x2","0x0","0x0"]},{"from_address":"0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d","keys":["0x99cd8bde557814842a3121e8ddfd433a539b8c9f14bf31ebf108d12e6196e9"],"data":["0x7461e8a41459e52d5ca62e6faee68c8149d3d8e974ca120ed8cb752192b3f5a","0x1176a1bd84444c89232ec27754698e5d2e7e1a7f1539f12027f28b23ec9f3d8","0x118dcc7fe511833","0x0"]}],"execution_resources":{"steps":145888,"pedersen_builtin_applications":520,"range_check_builtin_applications":7261,"bitwise_builtin_applications":270,"ec_op_builtin_applications":3,"poseidon_builtin_applications":21,"data_availability":{"l1_gas":0,"l1_data_gas":1280}}}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.TransactionReceipt
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
