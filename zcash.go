package coincodec

import (
	"errors"
	"github.com/wealdtech/go-slip44"
)

const (
	ZCashAddressLength = 22
	ZCashStaticPrefix = 0x1C // 28
	ZCashPrefixP2pkh = 0xB8 // 184
	ZCashPrefixP2sh = 0xBD // 189
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
	err = checkIsValid(decoded)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// ZCashEncodeToString converts the input byte array to a string representation of the ZCash address.
func ZCashEncodeToString(bytes []byte) (string, error) {
	err := checkIsValid(bytes)
	if err != nil {
		return "", err
	}
	encoded := Base58ChecksumEncode(bytes, Base58DefaultAlphabet)
	return encoded, nil
}

// Check for correct length and prefix
func checkIsValid(bytes []byte) (error) {
	if len(bytes) != ZCashAddressLength {
		return errors.New("Invalid length")
	}
	// prefix check: first byte is constant; second has two valid values 
	staticPrefix := bytes[0]
	prefix := bytes[1]
	if staticPrefix != ZCashStaticPrefix {
		return errors.New("Invalid static prefix")
	}
	if prefix != ZCashPrefixP2pkh && prefix != ZCashPrefixP2sh {
		return errors.New("Invalid prefix")
	}
	return nil
}
