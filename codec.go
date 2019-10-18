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
)

var toBytesMap = make(map[uint32]func(string) ([]byte, error))
var toStringMap = make(map[uint32]func([]byte) (string, error))

// ToBytes converts the input string to a byte array for the given coin type.
func ToBytes(input string, coinType uint32) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("no input")
	}
	f, exists := toBytesMap[coinType]
	if !exists {
		return nil, errors.New("unhandled coin type")
	}
	return f(input)
}

// ToString converts the input byte array to a string representation of the given coin type.
func ToString(input []byte, coinType uint32) (string, error) {
	if len(input) == 0 {
		return "", errors.New("no input")
	}
	f, exists := toStringMap[coinType]
	if !exists {
		return "", errors.New("unhandled coin type")
	}
	return f(input)
}
