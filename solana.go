package coincodec

import (
	"errors"

	"github.com/wealdtech/go-slip44"
)

const (
	SolanaAddressLength = 32
)

func init() {
	toBytesMap[slip44.SOLANA] = SolanaDecodeToBytes
	toStringMap[slip44.SOLANA] = SolanaEncodeToString
}

// SolanaDecodeToBytes converts the input string to a byte array
func SolanaDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58Decode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != SolanaAddressLength {
		return nil, errors.New("Invalid length")
	}
	return decoded, nil
}

// SolanaEncodeToString converts the input byte array to a string representation of the Solana address.
func SolanaEncodeToString(data []byte) (string, error) {
	if len(data) != SolanaAddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	encoded := Base58Encode(data, Base58DefaultAlphabet)
	return encoded, nil
}
