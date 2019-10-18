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
	"bytes"
	"errors"
	"testing"
)

func TestEtherToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []byte
		err    error
	}{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Ethereum address must have 40 characters"),
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
			output: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := EtherToBytes(test.input)
			if test.err != nil {
				if err == nil {
					t.Fatalf("Missing expected error: expected %v", test.err)
				}
				if test.err.Error() != err.Error() {
					t.Fatalf("Unexpected error value: expected %v, received %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if bytes.Compare(test.output, output) != 0 {
					t.Fatalf("Unexpected output: expected %x, received %x", test.output, output)
				}
			}
		})

	}
}

func TestEtherToString(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:  "Nil",
			input: nil,
			err:   errors.New("Ethereum address must have 20 bytes"),
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   errors.New("Ethereum address must have 20 bytes"),
		},
		{
			name:   "Good",
			input:  []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14},
			output: "0x0102030405060708090a0B0c0d0e0f1011121314",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := EtherToString(test.input)
			if test.err != nil {
				if err == nil {
					t.Fatalf("Missing expected error: expected %v", test.err)
				}
				if test.err.Error() != err.Error() {
					t.Fatalf("Unexpected error value: expected %v, received %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if test.output != output {
					t.Fatalf("Unexpected output: expected %x, received %x", test.output, output)
				}
			}
		})

	}
}
