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
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/wealdtech/go-slip44"
	"golang.org/x/crypto/sha3"
)

func init() {
	chains := []uint32{
		slip44.ETHER, slip44.ETHER_CLASSIC, slip44.GOCHAIN_GO,
		slip44.POA, slip44.CALLISTO, slip44.TOMO,
		slip44.THUNDERCORE, slip44.THETA, slip44.VECHAIN_TOKEN,
	}
	for _, c := range chains {
		toBytesMap[c] = EtherToBytes
		toStringMap[c] = EtherToString
	}
}

// EtherToBytes converts the input string to a byte array.
func EtherToBytes(input string) ([]byte, error) {
	input = strings.TrimPrefix(input, "0x")
	if len(input) != 40 {
		return nil, errors.New("Ethereum address must have 40 characters")
	}

	output, err := hex.DecodeString(input)
	if err != nil {
		return nil, errors.New("invalid hex string")
	}

	// Confirm checksum if present
	if strings.ToLower(input) != input && strings.ToUpper(input) != input {
		checksummed, err := EtherToString(output)
		if err != nil {
			return nil, errors.New("failed to validate checksum")
		}
		if checksummed[2:] != input {
			return nil, errors.New("invalid checksum")
		}
	}

	return output, nil
}

// EtherToString converts the input byte array to a string representation of the Ethereum address.
func EtherToString(input []byte) (string, error) {
	if len(input) != 20 {
		return "", errors.New("Ethereum address must have 20 bytes")
	}

	unchecksummed := hex.EncodeToString(input)
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return fmt.Sprintf("0x%s", result), nil
}
