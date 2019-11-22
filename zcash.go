package coincodec

import (
	"errors"
	"github.com/wealdtech/go-slip44"
)

const (
	ZCashAddressLength = 22
	ZCashStaticPrefix = 0x1c
)

func init() {
	toBytesMap[slip44.ZCASH] = ZCashDecodeToBytes
	toStringMap[slip44.ZCASH] = ZCashEncodeToString
}

// ZCashDecodeToBytes converts the input string to a byte array
func ZCashDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58ChecksumDecode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != ZCashAddressLength {
		return nil, errors.New("Invalid length")
	}
	// prefix check (t1, t3): first byte is constant, enforced; second has a few valid values, not checked here 
	staticPrefix := decoded[0]
	if staticPrefix != ZCashStaticPrefix {
		return nil, errors.New("Invalid static prefix")
	}
	return decoded, nil
}

// ZCashEncodeToString converts the input byte array to a string representation of the ZCash address.
func ZCashEncodeToString(bytes []byte) (string, error) {
	if len(bytes) != ZCashAddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	encoded := Base58ChecksumEncode(bytes, Base58DefaultAlphabet)
	return encoded, nil
}
