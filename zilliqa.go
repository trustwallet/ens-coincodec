package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	hrpZIL = "zil"
)

func init() {
	toBytesMap[slip44.ZILLIQA] = ZilliqaDecodeToBytes
	toStringMap[slip44.ZILLIQA] = ZilliqaEncodeToString
}

// ZilliqaDecodeToBytes converts the input string to a byte array
func ZilliqaDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpZIL, input)
	return bytes, err
}

// ZilliqaEncodeToString converts the input byte array to a string representation of the Cosmos address.
func ZilliqaEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpZIL, bytes)
	return output, err
}
