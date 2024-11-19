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

func TestEvents(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := jsonrpc.New(ctx,
		jsonrpc.WithLogLevel(zerolog.Disabled),
		jsonrpc.WithAddress(os.Getenv("JSONRPC_ADDRESS")),
		jsonrpc.WithTimeout(timeout),
	)
	require.NoError(t, err)

	response, err := s.Events(ctx, &api.EventsOpts{
		FromBlock: "1",
		ToBlock:   "latest",
		Limit:     10,
	})
	require.NoError(t, err)
	require.NotNil(t, response.Data)

	response2, err := s.Events(ctx, &api.EventsOpts{
		FromBlock: "300000",
		ToBlock:   "330000",
		Keys: [][]types.FieldElement{
			{
				strToFieldElement("0xc4a5eb3afec3e38cbe8f43f66c46bb0ca74ae6f10bfbd7c7f0f461d5cdb9f4"),
			},
			{
				strToFieldElement("0x6478e774beb00cd6b0714b0ad3ed3f80949b7e52b84f8f5792099d8990b7e26"),
			},
		},
		Limit: 10,
	})
	require.NoError(t, err)
	require.NotNil(t, response2.Data)
}
