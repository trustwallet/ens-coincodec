package coincodec

import (
	"bytes"

	"github.com/pkg/errors"
)

func Base58AddressIsValidString(input string, decodedSize int, validPrefixes [][]byte) error {
	decoded, err := Base58ChecksumDecode(input, Base58DefaultAlphabet)
	if err != nil {
		return err
	}
	return base58AddressIsValidData(decoded, decodedSize, validPrefixes)
}

func hasPrefix(data []byte, prefix []byte) bool {
	if len(data) < len(prefix) {
		return false
	}
	return bytes.Equal(data[:len(prefix)], prefix)
}

func base58AddressIsValidData(data []byte, decodedSize int, validPrefixes [][]byte) error {
	if len(data) != decodedSize {
		return errors.New("Invalid decoded length")
	}
	for _, prefix := range validPrefixes {
		if hasPrefix(data, prefix) {
			return nil
		}
	}
	return errors.New("Invalid prefix")
}

func Base58AddressDecodeToBytes(input string, decodedSize int, validPrefixes [][]byte) ([]byte, error) {
	decoded, err := Base58ChecksumDecode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	err = base58AddressIsValidData(decoded, decodedSize, validPrefixes)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func Base58AddressEncodeToString(data []byte, decodedSize int, validPrefixes [][]byte) (string, error) {
	err := base58AddressIsValidData(data, decodedSize, validPrefixes)
	if err != nil {
		return "", err
	}
	encoded := Base58ChecksumEncode(data, Base58DefaultAlphabet)
	return encoded, nil
}
