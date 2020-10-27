package coincodec

import (
	"github.com/wealdtech/go-slip44"
)

const (
	PolkadotNetwork = 0
)

func init() {
	toBytesMap[slip44.POLKADOT] = PolkadotDecodeToBytes
	toStringMap[slip44.POLKADOT] = PolkadotEncodeToString
}

// PolkadotDecodeToBytes converts the input string to a byte array
func PolkadotDecodeToBytes(input string) ([]byte, error) {
	return SS58AddressDecodeToBytes(input, PolkadotNetwork)
}

// PolkadotEncodeToString converts the input byte array to a string representation of the Polkadot address.
func PolkadotEncodeToString(data []byte) (string, error) {
	return SS58AddressEncodeToString(data, PolkadotNetwork)
}
