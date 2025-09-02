package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	TronAddressLength = 21
	TronPrefix        = 0x41
)

var (
	tronPrefixes = [][]byte{[]byte{TronPrefix}}
)

func init() {
	toBytesMap[coin.TRON] = TronDecodeToBytes
	toStringMap[coin.TRON] = TronEncodeToString
}

// TronDecodeToBytes converts the input string to a byte array
func TronDecodeToBytes(input string) ([]byte, error) {
	return Base58AddressDecodeToBytesPrefix(input, TronAddressLength, tronPrefixes)
}

// TronEncodeToString converts the input byte array to a string representation of the Tron address.
func TronEncodeToString(bytes []byte) (string, error) {
	return Base58AddressEncodeToStringPrefix(bytes, TronAddressLength, tronPrefixes)
}
