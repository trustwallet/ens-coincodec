package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	BNB_MAINNET_HRP = "bnb"
)

func init() {
	toBytesMap[slip44.BINANCE] = BNBDecodeToBytes
	toStringMap[slip44.BINANCE] = BNBEncodeToString
}

// BNBDecodeToBytes converts the input string to a byte array
func BNBDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(BNB_MAINNET_HRP, input)
	return bytes, err
}

// BNBEncodeToString converts the input byte array to a string representation of the BNB address.
func BNBEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(BNB_MAINNET_HRP, bytes)
	return output, err
}
