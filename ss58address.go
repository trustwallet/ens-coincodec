package coincodec

import (
	"bytes"
	"errors"

	"golang.org/x/crypto/blake2b"
)

const (
	ss58AddressLength       = 32
	ss58AddressChecksumSize = 2
	ss58AddressTotalLength  = ss58AddressLength + 1 + ss58AddressChecksumSize // incl. network byte and checksum
)

var (
	ss58Prefix = []byte("SS58PRE")
)

func computeChecksum(data []byte) []byte {
	prefixed := append(ss58Prefix, data...)
	hashres := blake2b.Sum512(prefixed)
	checksum := hashres[:ss58AddressChecksumSize]
	return checksum
}

// SS58AddressDecodeToBytes converts the input string to a byte array
func SS58AddressDecodeToBytes(input string, network byte) ([]byte, error) {
	decoded, err := Base58Decode(input, Base58DefaultAlphabet)
	if err != nil {
		return nil, err
	}
	if len(decoded) != ss58AddressTotalLength {
		return nil, errors.New("Invalid length")
	}
	// check network
	networkActual := decoded[0]
	if networkActual != network {
		return nil, errors.New("Invalid network")
	}
	checksum := computeChecksum(decoded[:len(decoded)-ss58AddressChecksumSize])
	// compare checksum
	if !bytes.Equal(checksum, decoded[len(decoded)-ss58AddressChecksumSize:]) {
		return nil, errors.New("Invalid checksum")
	}
	// strip network, checksum
	return decoded[1 : len(decoded)-2], nil
}

// SS58AddressEncodeToString converts the input byte array to a string representation of the address.
func SS58AddressEncodeToString(data []byte, network byte) (string, error) {
	if len(data) != ss58AddressLength {
		return "", errors.New("Invalid decoded address length")
	}
	networked := []byte{}
	networked = append(networked, network)
	networked = append(networked, data...)
	checksum := computeChecksum(networked)
	checksummed := append(networked, checksum...)
	encoded := Base58Encode(checksummed, Base58DefaultAlphabet)
	return encoded, nil
}
