package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	hrpZIL = "zil"
)

func init() {
	toBytesMap[coin.ZILLIQA] = ZilliqaDecodeToBytes
	toStringMap[coin.ZILLIQA] = ZilliqaEncodeToString
}

// ZilliqaDecodeToBytes converts the input string to a byte array
func ZilliqaDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpZIL, input, Bech32DefaultKeyhashLength)
	return bytes, err
}

// ZilliqaEncodeToString converts the input byte array to a string representation of the Zilliqa address.
func ZilliqaEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpZIL, bytes, Bech32DefaultKeyhashLength)
	return output, err
}
