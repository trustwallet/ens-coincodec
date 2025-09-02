package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	OntologyAddressLength = 21
	OntologyVersion       = 0x17
)

var (
	ontologyPrefixes = [][]byte{[]byte{OntologyVersion}}
)

func init() {
	toBytesMap[coin.ONTOLOGY] = OntologyDecodeToBytes
	toStringMap[coin.ONTOLOGY] = OntologyEncodeToString
}

// OntologyDecodeToBytes converts the input string to a byte array
func OntologyDecodeToBytes(input string) ([]byte, error) {
	return Base58AddressDecodeToBytesPrefix(input, OntologyAddressLength, ontologyPrefixes)
}

// OntologyEncodeToString converts the input byte array to a string representation of the Ontology address.
func OntologyEncodeToString(bytes []byte) (string, error) {
	return Base58AddressEncodeToStringPrefix(bytes, OntologyAddressLength, ontologyPrefixes)
}
