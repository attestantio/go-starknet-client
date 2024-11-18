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

package types

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// PublicKeyLength is the length of a startknet public key.
const PublicKeyLength = 32

// PublicKey is a starknet public key
//
//nolint:recvcheck
type PublicKey [PublicKeyLength]byte

var zeroPublicKey = PublicKey{}

// IsZero returns true if the public key is zero.
func (a *PublicKey) IsZero() bool {
	return bytes.Equal(a[:], zeroPublicKey[:])
}

// String returns the string representation of the public key.
func (a *PublicKey) String() string {
	res := hex.EncodeToString(a[:])
	// Leading 0s not allowed...
	res = strings.TrimLeft(res, "0")
	// ...unless that's all there was.
	if len(res) == 0 {
		res = "0"
	}

	return "0x" + res
}

// Format formats the address.
func (a *PublicKey) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, a.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, a[:])
	default:
		fmt.Fprintf(state, "%"+format, a[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *PublicKey) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("address missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid address prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid address suffix")
	}

	// Ensure that there are an even number of characters.
	bytesStr := string(input[3 : len(input)-1])
	if len(bytesStr)%2 == 1 {
		bytesStr = "0" + bytesStr
	}

	val, err := hex.DecodeString(bytesStr)
	if err != nil {
		return errors.New("invalid address")
	}
	copy(a[len(a)-len(val):], val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (a PublicKey) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}
