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

	"github.com/attestantio/go-starknet-client/types"
)

// EventsOpts are the options for events.
type EventsOpts struct {
	Common CommonOpts

	// FromBlock is the earliest block from which the events are obtained.
	// It can be a block number, block hash, or one of the special values "latest" or "pending".
	FromBlock string

	// ToBlock is the latest block from which the events are obtained.
	// It can be a block number, block hash, or one of the special values "latest" or "pending".
	ToBlock string

	// Address is the contract address from which the events are obtained.
	// If empty then events there is no address filter on returned events.
	Address *types.Address

	// Keys
	// If empty then there is no key filter on returned events.
	Keys []types.FieldElement

	// Limit is the maximum number of events to return.
	// This value must be provided.
	Limit uint32
}

// MarshalJSON marshals to a JSON representation.
func (o *EventsOpts) MarshalJSON() ([]byte, error) {
	filter := map[string]any{
		"to_block":   o.FromBlock,
		"chunk_size": o.Limit,
	}
	switch {
	case o.FromBlock == "latest" || o.FromBlock == "pending":
		filter["from_block"] = o.FromBlock
	case strings.HasPrefix(o.FromBlock, "0x"):
		filter["from_block"] = map[string]any{
			"block_hash": o.FromBlock,
		}
	default:
		fromBlock, err := strconv.ParseUint(o.FromBlock, 10, 64)
		if err != nil {
			return nil, errors.New("invalid from block")
		}
		filter["from_block"] = map[string]any{
			"block_number": fromBlock,
		}
	}
	switch {
	case o.ToBlock == "latest" || o.ToBlock == "pending":
		filter["to_block"] = o.ToBlock
	case strings.HasPrefix(o.ToBlock, "0x"):
		filter["to_block"] = map[string]any{
			"block_hash": o.ToBlock,
		}
	default:
		toBlock, err := strconv.ParseUint(o.ToBlock, 10, 64)
		if err != nil {
			return nil, errors.New("invalid to block")
		}
		filter["to_block"] = map[string]any{
			"block_number": toBlock,
		}
	}
	if o.Address != nil {
		filter["address"] = o.Address.String()
	}
	if len(o.Keys) > 0 {
		keys := make([]string, 0, len(o.Keys))
		for _, key := range o.Keys {
			keys = append(keys, key.String())
		}
		filter["keys"] = keys
	}

	eventsOpts := map[string]any{
		"filter": filter,
	}

	return json.Marshal(eventsOpts)
}

// String returns a string version of the structure.
func (o *EventsOpts) String() string {
	data, err := o.MarshalJSON()
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
