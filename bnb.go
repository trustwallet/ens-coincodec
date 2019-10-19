package coincodec

import (
	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

var bnbMainnetHRP = "bnb"
var bnbKeyHashLength = 20
var bnbKeyHashError = errors.New("BNB address key hash must be 20 bytes")

func init() {
	toBytesMap[slip44.BINANCE] = BNBDecodeToBytes
	toStringMap[slip44.BINANCE] = BNBEncodeToString
}

// BNBDecodeToBytes converts the input string to a byte array
func BNBDecodeToBytes(input string) ([]byte, error) {
	var empty []byte
	bytes, err := Bech32DecodeToBytes(input, bnbMainnetHRP)
	if err != nil {
		return empty, err
	}
	if len(bytes) != bnbKeyHashLength {
		return empty, bnbKeyHashError
	}
	return bytes, nil
}

// BNBEncodeToString converts the input byte array to a string representation of the BNB address.
func BNBEncodeToString(bytes []byte) (string, error) {
	if len(bytes) != bnbKeyHashLength {
		return "", bnbKeyHashError
	}
	return Bech32EncodeToString(bytes, bnbMainnetHRP)
}
