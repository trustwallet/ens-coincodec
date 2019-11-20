package coincodec

import (
	"github.com/wealdtech/go-slip44"
	"errors"
)

const (
	AddressLength = 23
)

func init() {
	toBytesMap[slip44.TEZOS] = TezosDecodeToBytes
	toStringMap[slip44.TEZOS] = TezosEncodeToString
}

// TezosDecodeToBytes converts the input string to a byte array
func TezosDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58Decode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != AddressLength {
		return nil, errors.New("Invalid length")
	}
	// no prefix check
	return decoded, nil
}

// TezosEncodeToString converts the input byte array to a string representation of the Cosmos address.
func TezosEncodeToString(bytes []byte) (string, error) {
	if len(bytes) != AddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	encoded := Base58Encode(bytes, Base58DefaultAlphabet)
	return encoded, nil
}
