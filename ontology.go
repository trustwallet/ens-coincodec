package coincodec

import (
	"errors"
	"github.com/wealdtech/go-slip44"
)

const (
	OntologyAddressLength = 21
	OntologyVersion = 0x17
)

func init() {
	toBytesMap[slip44.ONTOLOGY] = OntologyDecodeToBytes
	toStringMap[slip44.ONTOLOGY] = OntologyEncodeToString
}

// OntologyDecodeToBytes converts the input string to a byte array
func OntologyDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58ChecksumDecode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != OntologyAddressLength {
		return nil, errors.New("Invalid length")
	}
	// version check
	version := decoded[0]
	if (version != OntologyVersion) {
		return nil, errors.New("Invalid version")
	}
	return decoded, nil
}

// OntologyEncodeToString converts the input byte array to a string representation of the Ontology address.
func OntologyEncodeToString(bytes []byte) (string, error) {
	if len(bytes) != OntologyAddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	encoded := Base58ChecksumEncode(bytes, Base58DefaultAlphabet)
	return encoded, nil
}
