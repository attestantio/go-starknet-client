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
	"fmt"

	"github.com/attestantio/go-starknet-client/types"
)

// EventsOpts are the options for events.
type EventsOpts struct {
	Common CommonOpts

	// FromBlock is the earliest block from which the events are obtained.
	FromBlock types.BlockID

	// ToBlock is the latest block from which the events are obtained.
	// It can be a block number, block hash, or one of the special values "latest" or "pending".
	ToBlock types.BlockID

	// Address is the contract address from which the events are obtained.
	// If empty then events there is no address filter on returned events.
	Address *types.Address

	// Keys.
	// Each list corresponds to matching keys for a given location.
	// For example, [[a,b],[c,d] will return any event with either a or b in the first key field, and either c or d in the second
	// key field.  [[],[c,d]] will return any event with either c or d in the second key field regardless of what is in the first.
	// If empty then there is no key filter on returned events.
	Keys [][]types.FieldElement

	// Limit is the maximum number of events to return.
	// This value must be provided.
	Limit uint32
}

// MarshalJSON marshals to a JSON representation.
func (o *EventsOpts) MarshalJSON() ([]byte, error) {
	filter := map[string]any{
		"from_block": o.FromBlock,
		"to_block":   o.ToBlock,
		"chunk_size": o.Limit,
	}
	if o.Address != nil {
		filter["address"] = o.Address.String()
	}

	if len(o.Keys) > 0 {
		keys := make([][]string, 0, len(o.Keys))
		for _, keySet := range o.Keys {
			set := make([]string, 0, len(keySet))
			for _, key := range keySet {
				set = append(set, key.String())
			}

			keys = append(keys, set)
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
