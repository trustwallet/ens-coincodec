package coincodec

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/trustwallet/go-primitives/coin"
	"golang.org/x/crypto/sha3"
)

func init() {
	for _, c := range coin.Coins {
		if coin.IsEVM(c.ID) {
			toBytesMap[uint32(c.ID)] = EtherToBytes
			toStringMap[uint32(c.ID)] = EtherToString
		}
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
	_, _ = sha.Write([]byte(unchecksummed))
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
