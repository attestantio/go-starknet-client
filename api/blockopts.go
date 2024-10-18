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

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// BlockOpts are the options for blocks.
type BlockOpts struct {
	Common CommonOpts

	// Block is the ID of the block.
	// It can be a block number, block hash, or one of the special values "latest" or "pending".
	Block string
}

// MarshalJSON marshals to a JSON representation.
func (o *BlockOpts) MarshalJSON() ([]byte, error) {
	opts := map[string]any{}
	switch {
	case o.Block == "latest" || o.Block == "pending":
		opts["block_id"] = o.Block
	case strings.HasPrefix(o.Block, "0x"):
		opts["block_id"] = map[string]any{
			"block_hash": o.Block,
		}
	default:
		block, err := strconv.ParseUint(o.Block, 10, 64)
		if err != nil {
			return nil, errors.New("invalid from block")
		}
		opts["block_id"] = map[string]any{
			"block_number": block,
		}
	}

	return json.Marshal(opts)
}

// String returns a string version of the structure.
func (o *BlockOpts) String() string {
	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
