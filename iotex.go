package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	IOTEX_MAINNET_HRP = "io"
)

func init() {
	toBytesMap[slip44.IOTEX] = IoTexDecodeToBytes
	toStringMap[slip44.IOTEX] = IoTexEncodeToString
}

// Converts the input string to a byte array
func IoTexDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(IOTEX_MAINNET_HRP, input)
	return bytes, err
}

// Converts the input byte array to a string representation of the Cosmos address.
func IoTexEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(IOTEX_MAINNET_HRP, bytes)
	return output, err
}
