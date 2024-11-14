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

package jsonrpc

import (
	"context"

	client "github.com/attestantio/go-starknet-client"
	"github.com/attestantio/go-starknet-client/api"
)

// SpecVersion returns the version of the specification followed by the node.
func (s *Service) SpecVersion(ctx context.Context,
	opts *api.SpecVersionOpts,
) (
	*api.Response[string],
	error,
) {
	if err := s.assertIsActive(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	var data string
	if err := s.client.CallFor(&data, "starknet_specVersion"); err != nil {
		return nil, err
	}

	return &api.Response[string]{
		Data:     data,
		Metadata: map[string]any{},
	}, nil
}
