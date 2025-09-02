package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	TezosAddressLength = 23
)

var (
	tz1Prefix  = []byte{6, 161, 159}
	tz2Prefix  = []byte{6, 161, 161}
	tz3Prefix  = []byte{6, 161, 164}
	tzPrefixes = [][]byte{tz1Prefix, tz2Prefix, tz3Prefix}
)

func init() {
	toBytesMap[coin.TEZOS] = TezosDecodeToBytes
	toStringMap[coin.TEZOS] = TezosEncodeToString
}

// TezosDecodeToBytes converts the input string to a byte array
func TezosDecodeToBytes(input string) ([]byte, error) {
	return Base58AddressDecodeToBytesPrefix(input, TezosAddressLength, tzPrefixes)
}

// TezosEncodeToString converts the input byte array to a string representation of the Tezos address.
func TezosEncodeToString(bytes []byte) (string, error) {
	return Base58AddressEncodeToStringPrefix(bytes, TezosAddressLength, tzPrefixes)
}
