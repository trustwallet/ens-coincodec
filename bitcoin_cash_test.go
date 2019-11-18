package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestBitcoinCashDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "1BpEi6DfDAUFd7GtittLSdBeYJvcoaVggu",
			output: "76a91476a04053bda0a88bda5177b86a15c3b29f55987388ac",
		},
		{
			name:   "P2PKH CashAddr",
			input:  "bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a",
			output: "76a91476a04053bda0a88bda5177b86a15c3b29f55987388ac",
		},
		{
			name:   "P2SH",
			input:  "3CWFddi6m4ndiGyKqzYvsFYagqDLPVMTzC",
			output: "a91476a04053bda0a88bda5177b86a15c3b29f55987387",
		},
		{
			name:   "P2SH CashAddr",
			input:  "bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq",
			output: "a91476a04053bda0a88bda5177b86a15c3b29f55987387",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinCashDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("BitcoinDecodeToBytes() = %v, want %v, err: %v", hex.EncodeToString(got), tt.output, tt.err)
				}
			}
		})
	}
}

func TestBitcoinCashEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a914cb481232299cd5743151ac4b2d63ae198e7bb0a988ac")
	script2, _ := hex.DecodeString("a914cb481232299cd5743151ac4b2d63ae198e7bb0a987")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:  "Nil",
			input: nil,
			err:   errors.New("invalid data length"),
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   errors.New("invalid data length"),
		},
		{
			name:  "Wrong script",
			input: []byte{0x00, 0x14, 0x01, 0x2},
			err:   errors.New("wrong script data"),
		},
		{
			name:   "P2PKH",
			input:  script1,
			output: "bitcoincash:qr95sy3j9xwd2ap32xkykttr4cvcu7as4y0qverfuy",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "bitcoincash:pr95sy3j9xwd2ap32xkykttr4cvcu7as4yc93ky28e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinCashEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("BitcoinEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
