package coincodec

import (
	"errors"
)

var toBytesMap = make(map[uint32]func(string) ([]byte, error))
var toStringMap = make(map[uint32]func([]byte) (string, error))

// ToBytes converts the input string to a byte array for the given coin type.
func ToBytes(input string, coinType uint32) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("empty input")
	}
	f, exists := toBytesMap[coinType]
	if !exists {
		return nil, errors.New("unhandled coin type")
	}
	return f(input)
}

// ToString converts the input byte array to a string representation of the given coin type.
func ToString(input []byte, coinType uint32) (string, error) {
	if input == nil || len(input) == 0 {
		return "", errors.New("empty input")
	}
	f, exists := toStringMap[coinType]
	if !exists {
		return "", errors.New("unhandled coin type")
	}
	return f(input)
}
