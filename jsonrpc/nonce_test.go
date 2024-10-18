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
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestNonce(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tests := []struct {
		name    string
		opts    *api.NonceOpts
		err     string
		errCode int
	}{
		{
			name: "Empty",
			opts: &api.NonceOpts{
				Block:    "latest",
				Contract: strToAddress("0x0028c3ac8a0d8e8505486cd2857c309f1557cab0f93d9bb3686704d3cd26af96"),
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
			response, err := s.Nonce(ctx, test.opts)
			switch {
			case test.err != "":
				require.EqualError(t, err, test.err)
			// case test.errCode != 0:
			//	var apiErr *api.Error
			//	if errors.As(err, &apiErr) {
			//		require.Equal(t, test.errCode, apiErr.StatusCode)
			//	}
			default:
				require.NoError(t, err)
				require.NotNil(t, response)
				require.True(t, response.Data > 0)
			}
		})
	}
}
