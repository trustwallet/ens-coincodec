package coincodec

import (
	"errors"
	"testing"

	"github.com/wealdtech/go-slip44"
)

func TestEtherToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Blank",
			input: "0x",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Short",
			input: "0x0102030405060708090a0b0c0d0e0f10111213",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Long",
			input: "0x0102030405060708090a0b0c0d0e0f101112131415",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "Odd",
			input: "0x0102030405060708090a0b0c0d0e0f10111213141",
			err:   errors.New("Ethereum address must have 40 characters"),
		},
		{
			name:  "InvalidChars",
			input: "0xabcdefghijklmnopqrstuvwxyzabcdefghijklmn",
			err:   errors.New("invalid hex string"),
		},
		{
			name:  "BadChecksum",
			input: "0x0102030405060708090A0b0c0d0e0f1011121314",
			err:   errors.New("invalid checksum"),
		},
		{
			name:   "Good",
			input:  "0x0102030405060708090a0B0c0d0e0f1011121314",
			output: "0102030405060708090a0b0c0d0e0f1011121314",
		},
	}

	RunTestsEncode(t, slip44.ETHER, tests)
}

func TestEtherToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Ethereum address must have 20 bytes"),
		},
		{
			name:  "Too short",
			input: "0102030405",
			err:   errors.New("Ethereum address must have 20 bytes"),
		},
		{
			name:   "Good",
			input:  "0102030405060708090a0b0c0d0e0f1011121314",
			output: "0x0102030405060708090a0B0c0d0e0f1011121314",
		},
	}

	RunTestsDecode(t, slip44.ETHER, tests)
}
