package coincodec

import (
	"github.com/pkg/errors"
)

const (
	BECH32_KEYHASH_LENGTH = 20
)

var bech32KeyHashError = errors.New("A Bech32 address key hash must be 20 bytes")

// Bech32AddressDecodeToBytes converts the input string to a byte array
func Bech32AddressDecodeToBytes(hrp string, input string) ([]byte, error) {
	var empty []byte
	bytes, err := Bech32DecodeToBytes(input, hrp)
	if err != nil {
		return empty, err
	}
	if len(bytes) != BECH32_KEYHASH_LENGTH {
		return empty, bech32KeyHashError
	}
	return bytes, nil
}

// Bech32AddressEncodeToString converts the input byte array to a string representation of the bech32 address.
func Bech32AddressEncodeToString(hrp string, bytes []byte) (string, error) {
	if len(bytes) != BECH32_KEYHASH_LENGTH {
		return "", bech32KeyHashError
	}
	return Bech32EncodeToString(bytes, hrp)
}
