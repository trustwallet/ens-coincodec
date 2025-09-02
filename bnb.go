package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	hrpBNB = "bnb"
)

func init() {
	toBytesMap[coin.BINANCE] = BNBDecodeToBytes
	toStringMap[coin.BINANCE] = BNBEncodeToString
}

// BNBDecodeToBytes converts the input string to a byte array
func BNBDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpBNB, input, Bech32DefaultKeyhashLength)
	return bytes, err
}

// BNBEncodeToString converts the input byte array to a string representation of the BNB address.
func BNBEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpBNB, bytes, Bech32DefaultKeyhashLength)
	return output, err
}
