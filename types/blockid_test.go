// Copyright © 2024 - 2026 Attestant Limited.
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

package types

import (
	"fmt"
	"testing"
)

func TestBlockIDFormat(t *testing.T) {
	tests := []struct {
		name   string
		block  BlockID
		format string
	}{
		{name: "latest-s", block: "latest", format: "%s"},
		{name: "latest-q", block: "latest", format: "%q"},
		{name: "latest-v", block: "latest", format: "%v"},
		{name: "latest-x", block: "latest", format: "%x"},
		{name: "latest-X", block: "latest", format: "%X"},
		{name: "latest-#x", block: "latest", format: "%#x"},
		{name: "latest-#X", block: "latest", format: "%#X"},
		{name: "latest-+s", block: "latest", format: "%+s"},
		{name: "latest-d", block: "latest", format: "%d"},
		{name: "pending-s", block: "pending", format: "%s"},
		{name: "pending-q", block: "pending", format: "%q"},
		{name: "pending-v", block: "pending", format: "%v"},
		{name: "pending-x", block: "pending", format: "%x"},
		{name: "pending-X", block: "pending", format: "%X"},
		{name: "pending-#x", block: "pending", format: "%#x"},
		{name: "pending-#X", block: "pending", format: "%#X"},
		{name: "pending-+s", block: "pending", format: "%+s"},
		{name: "pending-d", block: "pending", format: "%d"},
		{name: "hash-s", block: "0xabc123", format: "%s"},
		{name: "hash-q", block: "0xabc123", format: "%q"},
		{name: "hash-v", block: "0xabc123", format: "%v"},
		{name: "hash-x", block: "0xabc123", format: "%x"},
		{name: "hash-X", block: "0xabc123", format: "%X"},
		{name: "hash-#x", block: "0xabc123", format: "%#x"},
		{name: "hash-#X", block: "0xabc123", format: "%#X"},
		{name: "hash-+s", block: "0xabc123", format: "%+s"},
		{name: "hash-d", block: "0xabc123", format: "%d"},
		{name: "number-s", block: "12345", format: "%s"},
		{name: "number-q", block: "12345", format: "%q"},
		{name: "number-v", block: "12345", format: "%v"},
		{name: "number-x", block: "12345", format: "%x"},
		{name: "number-X", block: "12345", format: "%X"},
		{name: "number-#x", block: "12345", format: "%#x"},
		{name: "number-#X", block: "12345", format: "%#X"},
		{name: "number-+s", block: "12345", format: "%+s"},
		{name: "number-d", block: "12345", format: "%d"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fmt.Sprintf(tt.format, tt.block)
			want := fmt.Sprintf(tt.format, string(tt.block))
			if got != want {
				t.Fatalf("unexpected output: got %q want %q", got, want)
			}
		})
	}
}
