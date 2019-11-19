package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	ZIL_MAINNET_HRP = "zil"
)

func init() {
	toBytesMap[slip44.ZILLIQA] = ZilliqaDecodeToBytes
	toStringMap[slip44.ZILLIQA] = ZilliqaEncodeToString
}

// Converts the input string to a byte array
func ZilliqaDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(ZIL_MAINNET_HRP, input)
	return bytes, err
}

// Converts the input byte array to a string representation of the Cosmos address.
func ZilliqaEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(ZIL_MAINNET_HRP, bytes)
	return output, err
}
