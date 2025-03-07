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
	"time"

	client "github.com/attestantio/go-starknet-client"
	"github.com/attestantio/go-starknet-client/jsonrpc"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		parameters []jsonrpc.Parameter
		location   string
		err        string
	}{
		{
			name: "Nil",
			err:  "no address specified",
		},
		{
			name: "AddressNil",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithLogLevel(zerolog.Disabled),
				jsonrpc.WithTimeout(5 * time.Second),
			},
			err: "no address specified",
		},
		{
			name: "TimeoutZero",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithLogLevel(zerolog.Disabled),
				jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
				jsonrpc.WithTimeout(0),
			},
			err: "no timeout specified",
		},
		{
			name: "AddressInvalid",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithLogLevel(zerolog.Disabled),
				jsonrpc.WithAddress(string([]byte{0x01})),
				jsonrpc.WithTimeout(5 * time.Second),
			},
			err: "invalid URL\nparse \"http://\\x01\": net/url: invalid control character in URL",
		},
		{
			name: "Good",
			parameters: []jsonrpc.Parameter{
				jsonrpc.WithLogLevel(zerolog.Disabled),
				jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
				jsonrpc.WithTimeout(5 * time.Second),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := jsonrpc.New(ctx, test.parameters...)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInterfaces(t *testing.T) {
	ctx := context.Background()
	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(5*time.Second),
	)
	require.NoError(t, err)

	assert.Implements(t, (*client.BlockHashAndNumberProvider)(nil), s)
	assert.Implements(t, (*client.BlockNumberProvider)(nil), s)
	assert.Implements(t, (*client.BlockProvider)(nil), s)
	assert.Implements(t, (*client.CallProvider)(nil), s)
	assert.Implements(t, (*client.ChainIDProvider)(nil), s)
	assert.Implements(t, (*client.EventsProvider)(nil), s)
	assert.Implements(t, (*client.NonceProvider)(nil), s)
	assert.Implements(t, (*client.ProtocolVersionProvider)(nil), s)
	assert.Implements(t, (*client.SpecVersionProvider)(nil), s)
	assert.Implements(t, (*client.SyncingProvider)(nil), s)
}
