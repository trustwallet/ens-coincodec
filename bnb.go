package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	hrpBNB = "bnb"
)

func init() {
	toBytesMap[slip44.BINANCE] = BNBDecodeToBytes
	toStringMap[slip44.BINANCE] = BNBEncodeToString
}

// BNBDecodeToBytes converts the input string to a byte array
func BNBDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpBNB, input)
	return bytes, err
}

// BNBEncodeToString converts the input byte array to a string representation of the BNB address.
func BNBEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpBNB, bytes)
	return output, err
}
