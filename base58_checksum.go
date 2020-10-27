package coincodec

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Base58Decode decodes a modified base58 string to a byte slice and checks checksum.
func Base58ChecksumDecode(b string, alphabet string) ([]byte, error) {
	val, err := Base58Decode(b, alphabet)
	if err != nil {
		return nil, err
	}
	if len(val) < 4 {
		return nil, fmt.Errorf("Base58 string too short, len %v", len(val))
	}

	// Check checksum
	checksum := doubleSha256(val[:len(val)-4])
	expected := val[len(val)-4:]
	if !bytes.Equal(checksum[0:4], expected) {
		return nil, fmt.Errorf("Bad Base58 checksum: %v expected %v", checksum, expected)
	}

	// strip checksum
	return val[:len(val)-4], nil
}

// Base58Encode encodes a byte slice to a modified base58 string.
func Base58ChecksumEncode(b []byte, alphabet string) string {
	checksum := doubleSha256(b)
	b = append(b, checksum[0:4]...)
	s := Base58Encode(b, alphabet)
	return s
}

func doubleSha256(b []byte) []byte {
	hasher := sha256.New()
	_, _ = hasher.Write(b)
	sha := hasher.Sum(nil)
	hasher.Reset()
	_, _ = hasher.Write(sha)
	return hasher.Sum(nil)
}
