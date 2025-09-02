package coincodec

import (
	"bytes"
	"errors"
	"testing"

	"github.com/trustwallet/go-primitives/coin"
)

func TestToBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		coinType uint32
		output   []byte
		err      error
	}{
		{
			name:     "Empty",
			input:    "",
			coinType: coin.ETHEREUM,
			err:      errors.New("empty input"),
		},
		{
			name:     "Unknown",
			input:    "unknown",
			coinType: 6543253,
			err:      errors.New("unhandled coin type"),
		},
		{
			name:     "Good",
			input:    "0x0102030405060708090a0B0c0d0e0f1011121314",
			coinType: coin.ETHEREUM,
			output:   []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := ToBytes(test.input, test.coinType)
			if test.err != nil {
				if err == nil {
					t.Fatalf("Missing expected error: expected %v", test.err)
				}
				if test.err.Error() != err.Error() {
					t.Fatalf("Unexpected error value: expected %v, received %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if !bytes.Equal(test.output, output) {
					t.Fatalf("Unexpected output: expected %x, received %x", test.output, output)
				}
			}
		})

	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		coinType uint32
		output   string
		err      error
	}{
		{
			name:     "Nil",
			input:    nil,
			coinType: coin.ETHEREUM,
			err:      errors.New("empty input"),
		},
		{
			name:     "Empty",
			input:    []byte{},
			coinType: coin.ETHEREUM,
			err:      errors.New("empty input"),
		},
		{
			name:     "Unknown",
			input:    []byte{0x01, 0x02, 0x03, 0x04},
			coinType: 6543253,
			err:      errors.New("unhandled coin type"),
		},
		{
			name:     "Good",
			input:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14},
			coinType: coin.ETHEREUM,
			output:   "0x0102030405060708090a0B0c0d0e0f1011121314",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := ToString(test.input, test.coinType)
			if test.err != nil {
				if err == nil {
					t.Fatalf("Missing expected error: expected %v", test.err)
				}
				if test.err.Error() != err.Error() {
					t.Fatalf("Unexpected error value: expected %v, received %v", test.err, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if test.output != output {
					t.Fatalf("Unexpected output: expected %x, received %x", test.output, output)
				}
			}
		})

	}
}
