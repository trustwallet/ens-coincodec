package coincodec

import (
	"errors"
	"github.com/wealdtech/go-slip44"
)

const (
	TronAddressLength = 21
	TronPrefix        = 0x41
)

func init() {
	toBytesMap[slip44.TRON] = TronDecodeToBytes
	toStringMap[slip44.TRON] = TronEncodeToString
}

// TronDecodeToBytes converts the input string to a byte array
func TronDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58ChecksumDecode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != TronAddressLength {
		return nil, errors.New("Invalid length")
	}
	// prefix check
	prefix := decoded[0]
	if prefix != TronPrefix {
		return nil, errors.New("Invalid prefix")
	}
	return decoded, nil
}

// TronEncodeToString converts the input byte array to a string representation of the Tron address.
func TronEncodeToString(bytes []byte) (string, error) {
	if len(bytes) != TronAddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	encoded := Base58ChecksumEncode(bytes, Base58DefaultAlphabet)
	return encoded, nil
}
