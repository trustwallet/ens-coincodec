package coincodec

import (
	"bytes"
	"encoding/base32"

	"github.com/pkg/errors"
	"github.com/trustwallet/go-primitives/coin"
)

const (
	versionByteAccountId byte = 0x30
)

func init() {
	toBytesMap[coin.STELLAR] = StellarDecodeToBytes
	toStringMap[coin.STELLAR] = StellarEncodeToString

	toBytesMap[coin.KIN] = StellarDecodeToBytes
	toStringMap[coin.KIN] = StellarEncodeToString
}

// StellarDecodeToBytes converts the input string to a byte array
func StellarDecodeToBytes(input string) ([]byte, error) {
	decoded, err := base32.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, errors.Wrap(err, "base32 decode error")
	}
	if decoded[0] != versionByteAccountId {
		return nil, errors.New("invalid version byte")
	}
	checksum := crc16(decoded[:len(decoded)-2])
	if !bytes.Equal(checksum, decoded[len(decoded)-2:]) {
		return nil, errors.New("wrong checksum")
	}
	// strip version byte and checksum
	return decoded[1 : len(decoded)-2], nil
}

// StellarEncodeToString converts the input byte array to a string representation of the Cosmos address.
func StellarEncodeToString(bytes []byte) (string, error) {
	data := []byte{versionByteAccountId}
	data = append(data, bytes...)
	checksum := crc16(data)
	data = append(data, checksum...)
	return base32.StdEncoding.EncodeToString(data), nil
}

func crc16(bytes []byte) []byte {
	// https://godoc.org/github.com/stellar/go/crc16
	crc := uint16(0x0000)
	polynomial := uint16(0x1021)
	for _, b := range bytes {
		for bitidx := byte(0); bitidx < 8; bitidx++ {
			bit := (b >> (7 - bitidx) & 1) == 1
			c15 := (crc >> 15 & 1) == 1
			crc <<= 1
			if c15 != bit {
				crc ^= polynomial
			}
		}
	}
	crc = crc & uint16(0xffff)
	checksum := []byte{byte(crc & 0x00ff), byte((crc >> 8) & 0x00ff)}
	return checksum
}
