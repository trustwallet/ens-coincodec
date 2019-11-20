package coincodec

import (
	"errors"
	"bytes"
	"github.com/wealdtech/go-slip44"
)

const (
	AddressLength = 23
)

var (
	tz1Prefix = []byte{6, 161, 159}
	tz2Prefix = []byte{6, 161, 161}
	tz3Prefix = []byte{6, 161, 164}
)

func init() {
	toBytesMap[slip44.TEZOS] = TezosDecodeToBytes
	toStringMap[slip44.TEZOS] = TezosEncodeToString
}

// TezosDecodeToBytes converts the input string to a byte array
func TezosDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58ChecksumDecode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != AddressLength {
		return nil, errors.New("Invalid length")
	}
	// prefix check
	prefix := decoded[0:3]
	if (!bytes.Equal(prefix, tz1Prefix) &&
		!bytes.Equal(prefix, tz2Prefix) &&
		!bytes.Equal(prefix, tz3Prefix)) {
		return nil, errors.New("Invalid prefix")
	}
	return decoded, nil
}

// TezosEncodeToString converts the input byte array to a string representation of the Cosmos address.
func TezosEncodeToString(bytes []byte) (string, error) {
	if len(bytes) != AddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	encoded := Base58ChecksumEncode(bytes, Base58DefaultAlphabet)
	return encoded, nil
}
