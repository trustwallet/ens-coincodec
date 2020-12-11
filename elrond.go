package coincodec

const (
	hrpElrond           = "erd"
	ElrondKeyhashLength = 32
)

func init() {
	// 508 slip44.ELROND
	toBytesMap[508] = ElrondDecodeToBytes
	toStringMap[508] = ElrondEncodeToString
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
