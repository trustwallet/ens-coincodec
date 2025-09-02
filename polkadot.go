package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	PolkadotNetwork = 0
)

func init() {
	toBytesMap[coin.POLKADOT] = PolkadotDecodeToBytes
	toStringMap[coin.POLKADOT] = PolkadotEncodeToString
}

// PolkadotDecodeToBytes converts the input string to a byte array
func PolkadotDecodeToBytes(input string) ([]byte, error) {
	return SS58AddressDecodeToBytes(input, PolkadotNetwork)
}

// PolkadotEncodeToString converts the input byte array to a string representation of the Polkadot address.
func PolkadotEncodeToString(data []byte) (string, error) {
	return SS58AddressEncodeToString(data, PolkadotNetwork)
}
