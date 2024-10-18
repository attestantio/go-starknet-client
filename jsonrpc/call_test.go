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
	"github.com/attestantio/go-starknet-client/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name     string
		opts     *api.CallOpts
		expected []types.FieldElement
		err      string
		errCode  int
	}{
		{
			name: "SingleReturnElement",
			opts: &api.CallOpts{
				Block:              "latest",
				Contract:           strToAddress("0x0028c3ac8a0d8e8505486cd2857c309f1557cab0f93d9bb3686704d3cd26af96"),
				EntryPointSelector: strToFieldElement("0x1a35984e05126dbecb7c3bb9929e7dd9106d460c59b1633739a5c733a5fb13b"),
			},
			expected: []types.FieldElement{
				strToFieldElement("0x1ff755e135eac251b4a10bc7aadd11e37a4cab7a552b52d99827c85605ba333"),
			},
		},
		{
			name: "MultipleReturnElements",
			opts: &api.CallOpts{
				Block:              "latest",
				Contract:           strToAddress("0x0590e76a2e65435b7288bf3526cfa5c3ec7748d2f3433a934c931cce62460fc5"),
				EntryPointSelector: strToFieldElement("0x36735aa694184cd8116c479c296d9431cc04a470e0467c07067e4586f647ece"),
			},
			expected: []types.FieldElement{
				strToFieldElement("0xde0b6b3a7640000"),
				strToFieldElement("0x4718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d"),
				strToFieldElement("0x1f777aef22"),
				strToFieldElement("0x3660fc6334c9485065394f6432933c2f04b4716c67511ba174384c65faebc19"),
				strToFieldElement("0x52bf6ec001452dbace0ca4c7db1b232b8031b5a0ccb117bb47a05569df435de"),
				strToFieldElement("0x12c"),
			},
		},
	}

	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := s.Call(ctx, test.opts)
			switch {
			case test.err != "":
				require.ErrorContains(t, err, test.err)
			// case test.errCode != 0:
			//	var apiErr *api.Error
			//	if errors.As(err, &apiErr) {
			//		require.Equal(t, test.errCode, apiErr.StatusCode)
			//	}
			default:
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotNil(t, test.expected, response.Data)
			}
		})
	}
}
