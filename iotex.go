package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	hrpIOTEX = "io"
)

func init() {
	toBytesMap[coin.IOTEX] = IoTexDecodeToBytes
	toStringMap[coin.IOTEX] = IoTexEncodeToString
}

// IoTexDecodeToBytes converts the input string to a byte array
func IoTexDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpIOTEX, input, Bech32DefaultKeyhashLength)
	return bytes, err
}

// IoTexEncodeToString converts the input byte array to a string representation of the Cosmos address.
func IoTexEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpIOTEX, bytes, Bech32DefaultKeyhashLength)
	return output, err
}
