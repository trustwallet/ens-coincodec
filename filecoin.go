package coincodec

import (
	"bytes"
	"encoding/base32"
	"errors"
	"strconv"

	"github.com/trustwallet/go-primitives/coin"
	"golang.org/x/crypto/blake2b"
)

const (
	PREFIX                   byte   = 'f'
	checksumSize             byte   = 4
	BASE32_ALPHABET_FILECOIN string = "abcdefghijklmnopqrstuvwxyz234567"
)

var (
	base32Encoding = base32.NewEncoding(BASE32_ALPHABET_FILECOIN).WithPadding(base32.NoPadding)
)

type Type byte

const (
	TypeID = iota
	TypeSecp256k1
	TypeActor
	TypeBls
	TypeInvalid
)

func init() {
	toBytesMap[coin.FILECOIN] = FilecoinDecodeToBytes
	toStringMap[coin.FILECOIN] = FilecoinEncodeToString
}

// Attempts to get the type by ASCII.
func parseType(ascii byte) Type {
	if ascii >= '0' && ascii <= '3' {
		return Type(ascii - '0')
	}
	return TypeInvalid
}

func typeAscii(typ Type) string {
	return string(byte('0') + byte(typ))
}

func getType(raw byte) Type {
	switch raw {
	case 0:
		return TypeID
	case 1:
		return TypeSecp256k1
	case 2:
		return TypeActor
	case 3:
		return TypeBls
	default:
		return TypeInvalid
	}
}

// Returns the payload size (excluding any prefixes) of an address type.
// If the payload size is undefined/variable (e.g. ID)
// or the type is unknown, it returns zero.
func payloadSize(t Type) byte {
	switch t {
	case TypeSecp256k1:
		return 20
	case TypeActor:
		return 20
	case TypeBls:
		return 48
	}
	return 0
}

func isValidID(id string) bool {
	if len(id) > 22 {
		return false
	}
	for i := 2; i < len(id); i++ {
		if id[i] < '0' || id[i] > '9' {
			return false
		}
	}
	_, err := strconv.ParseUint(id[2:], 10, 0)
	return err == nil
}

func filecoinComputeChecksum(address []byte) []byte {
	hash, _ := blake2b.New(int(checksumSize), nil)
	hash.Write(address)
	sum := hash.Sum(nil)
	return sum
}

func isValidBase32(input string, t Type) error {
	// Check if valid Base32.
	size := payloadSize(t)
	var decoded []byte = make([]byte, size+checksumSize)
	_, err := base32Encoding.Decode(decoded, []byte(input[2:]))
	if err != nil {
		return errors.New("decoding base32 failed")
	}
	// Check size
	if len(decoded) != int(size+checksumSize) {
		return errors.New("Invalid size")
	}
	// Extract raw address
	var address []byte = make([]byte, 0)
	address = append(address, byte(t))
	address = append(address, decoded[:size]...)
	// Verify checksum
	shouldSum := filecoinComputeChecksum(address)
	if bytes.Compare(shouldSum[:], decoded[size:]) != 0 {
		return errors.New("Wrong checksum")
	}
	return nil
}

func isValidString(input string) error {
	if len(input) < 3 {
		return errors.New("Too short")
	}
	// Only main net addresses supported.
	if input[0] != PREFIX {
		return errors.New("Invalid network")
	}
	// Get address type.
	typ := parseType(input[1])
	if typ == TypeInvalid {
		return errors.New("Invalid type")
	}
	// ID addresses are special, they are just numbers.
	if typ == TypeID {
		if !isValidID(input) {
			return errors.New("Invalid ID")
		}
		return nil
	}
	return isValidBase32(input, typ)
}

// FilecoinDecodeToBytes converts the input string to a byte array
func FilecoinDecodeToBytes(input string) ([]byte, error) {
	if err := isValidString(input); err != nil {
		return nil, err
	}
	typ := parseType(input[1])
	var bytes []byte = make([]byte, 0)
	// First byte is type
	bytes = append(bytes, byte(typ))
	if typ == TypeID {
		id, _ := strconv.ParseUint(input[2:], 10, 0) // error checked in isValid
		for id >= 0x80 {
			bytes = append(bytes, byte(id)|byte(0x80))
			id >>= 7
		}
		bytes = append(bytes, byte(id))
		return bytes, nil
	}
	payloadSize := payloadSize(typ)
	var decoded []byte = make([]byte, payloadSize+checksumSize)
	_, err := base32Encoding.Decode(decoded, []byte(input[2:]))
	if err != nil {
		return nil, errors.New("Invalid address, base32 decoding failed")
	}
	bytes = append(bytes, decoded[:payloadSize]...)
	return bytes, nil
}

// FilecoinEncodeToString converts the input byte array to a string representation of the Solana address.
func FilecoinEncodeToString(data []byte) (string, error) {
	if len(data) < 2 {
		return "", errors.New("Data too short")
	}
	var t Type = getType(data[0])
	if t == TypeInvalid {
		return "", errors.New("Invalid type")
	}
	var s string
	// Main net address prefix
	s += string(PREFIX)
	// Address type prefix
	s += typeAscii(t)

	if t == TypeID {
		var id uint = 0
		shift := 0
		for i := 1; i < len(data); i++ {
			if data[i] < 0x80 {
				id |= uint(data[i]) << shift
				break
			} else {
				id |= (uint(data[i] & 0x7F)) << shift
				shift += 7
			}
		}
		s += strconv.FormatUint(uint64(id), 10)
		return s, nil
	}

	psize := payloadSize(t)
	if len(data) != 1+int(psize) {
		return "", errors.New("Invalid length")
	}
	// Base32 encoded body
	var toEncode []byte = make([]byte, 0)
	// Copy address payload without prefix
	toEncode = append(toEncode, data[1:psize+1]...)
	// Append checksum
	sum := filecoinComputeChecksum(data)
	toEncode = append(toEncode, sum...)
	s += base32Encoding.EncodeToString(toEncode)
	return s, nil
}
