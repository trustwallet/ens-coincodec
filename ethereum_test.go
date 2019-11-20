// Copyright © 2019 Weald Technology Trading
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

package coincodec

import (
	"errors"
	"testing"

	"github.com/wealdtech/go-slip44"
)

func TestEtherToBytes(t *testing.T) {
	tests := []TestcaseEncode {
		{
			name:  "Empty",
			input: "",
			err:   errors.New("empty input"),
		},
		{
			name:  "Blank",
			input: "0x",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Short",
			input: "0x0102030405060708090a0b0c0d0e0f10111213",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Long",
			input: "0x0102030405060708090a0b0c0d0e0f101112131415",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Odd",
			input: "0x0102030405060708090a0b0c0d0e0f10111213141",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "InvalidChars",
			input: "0xabcdefghijklmnopqrstuvwxyzabcdefghijklmn",
			err:   errors.New("invalid hex string"),
		},
		{
			name:  "BadChecksum",
			input: "0x0102030405060708090A0b0c0d0e0f1011121314",
			err:   errors.New("invalid checksum"),
		},
		{
			name:   "Good",
			input:  "0x0102030405060708090a0B0c0d0e0f1011121314",
			output: "0102030405060708090a0b0c0d0e0f1011121314",
		},
	}

	RunTestsEncode(t, slip44.ETHER, tests)
}

func TestEtherToString(t *testing.T) {
	tests := []TestcaseDecode {
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Ethereum address must have 20 bytes"),
		},
		{
			name:  "Too short",
			input: "0102030405",
			err:   errors.New("Ethereum address must have 20 bytes"),
		},
		{
			name:   "Good",
			input: "0102030405060708090a0b0c0d0e0f1011121314",
			output: "0x0102030405060708090a0B0c0d0e0f1011121314",
		},
	}

	RunTestsDecode(t, slip44.ETHER, tests)
}
