package coincodec

import "github.com/wealdtech/go-slip44"

const (
	hrpElrond           = "erd"
	ElrondKeyhashLength = 32
)

func init() {
	toBytesMap[slip44.ELROND] = ElrondDecodeToBytes
	toStringMap[slip44.ELROND] = ElrondEncodeToString
}

// ElrondDecodeToBytes converts the input string to a byte array
func ElrondDecodeToBytes(input string) ([]byte, error) {
	bytes, err := Bech32AddressDecodeToBytes(hrpElrond, input, ElrondKeyhashLength)
	return bytes, err
}

// ElrondEncodeToString converts the input byte array to a string representation of the Elrond address.
func ElrondEncodeToString(bytes []byte) (string, error) {
	output, err := Bech32AddressEncodeToString(hrpElrond, bytes, ElrondKeyhashLength)
	return output, err
}
