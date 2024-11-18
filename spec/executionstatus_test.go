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

func TestExecutionStatus(t *testing.T) {
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
			err:   "unrecognised execution status []",
		},
		{
			name:  "SUCCEEDED",
			input: []byte(`"SUCCEEDED"`),
		},
		{
			name:  "REVERTED",
			input: []byte(`"REVERTED"`),
		},
		{
			name:     "Succeeded",
			input:    []byte(`"succeeded"`),
			expected: []byte(`"SUCCEEDED"`),
		},
		{
			name:  "Unknown",
			input: []byte(`"unknown"`),
			err:   `unrecognised execution status "unknown"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var res spec.ExecutionStatus
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
