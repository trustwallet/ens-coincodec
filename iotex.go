package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	hrpIOTEX = "io"
)

func init() {
	toBytesMap[slip44.IOTEX] = IoTexDecodeToBytes
	toStringMap[slip44.IOTEX] = IoTexEncodeToString
}

// IoTexDecodeToBytes converts the input string to a byte array
func IoTexDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpIOTEX, input)
	return bytes, err
}

// IoTexEncodeToString converts the input byte array to a string representation of the Cosmos address.
func IoTexEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpIOTEX, bytes)
	return output, err
}
