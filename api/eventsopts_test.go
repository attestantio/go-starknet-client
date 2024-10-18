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
	"testing"

	"github.com/attestantio/go-starknet-client/api"
	"github.com/stretchr/testify/require"
)

// TestEventsOptsMarshalJSON tests JSON for EventsOpts.
func TestEventsOptsMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    *api.EventsOpts
		expected []byte
		err      string
	}{
		{
			name:     "Nil",
			expected: []byte(`null`),
		},
		{
			name: "Names",
			input: &api.EventsOpts{
				FromBlock: "latest",
				ToBlock:   "pending",
				Limit:     5,
			},
			expected: []byte(`{"filter":{"chunk_size":5,"from_block":"latest","to_block":"pending"}}`),
		},
		{
			name: "Numbers",
			input: &api.EventsOpts{
				FromBlock: "1",
				ToBlock:   "2",
				Limit:     5,
			},
			expected: []byte(`{"filter":{"chunk_size":5,"from_block":{"block_number":1},"to_block":{"block_number":2}}}`),
		},
		{
			name: "Hashes",
			input: &api.EventsOpts{
				FromBlock: "0x445152a69e628774b0f78a952e6f9ba0ffcda1374724b314140928fd2f31f4c",
				ToBlock:   "0x2e59a5adbdf53e00fd282a007b59771067870c1c7664ca7878327adfff398b4",
				Limit:     5,
			},
			expected: []byte(`{"filter":{"chunk_size":5,"from_block":{"block_hash":"0x445152a69e628774b0f78a952e6f9ba0ffcda1374724b314140928fd2f31f4c"},"to_block":{"block_hash":"0x2e59a5adbdf53e00fd282a007b59771067870c1c7664ca7878327adfff398b4"}}}`),
		},
		{
			name: "Address",
			input: &api.EventsOpts{
				FromBlock: "1",
				ToBlock:   "2",
				Limit:     5,
				Address:   mustUnmarshalAddress(`"0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"`),
			},
			expected: []byte(`{"filter":{"address":"0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7","chunk_size":5,"from_block":{"block_number":1},"to_block":{"block_number":2}}}`),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := json.Marshal(test.input)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, res)
			}
		})
	}
}
