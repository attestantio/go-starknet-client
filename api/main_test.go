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

package api_test

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/attestantio/go-starknet-client/types"
	"github.com/rs/zerolog"
)

// timeout for tests.
var timeout = 60 * time.Second

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	if os.Getenv("JSONRPC_ADDRESS") != "" {
		os.Exit(m.Run())
	}
}

func mustUnmarshalAddress(input string) *types.Address {
	var address types.Address
	if err := json.Unmarshal([]byte(input), &address); err != nil {
		panic(err)
	}

	return &address
}
