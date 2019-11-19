package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestCosmosDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
			output: "bc2da90c84049370d1b7c528bc164bc588833f21",
		},
		{
			name:   "Normal2",
			input:  "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
			output: "6e436a571cec916167ba105160474b9c9cd132bd",
		},
		{
			name:  "Testnet",
			input: "cosmosvaloper1sxx9mszve0gaedz5ld7qdkjkfv8z992ax69k08",
			err:   errors.New("decoded hrp mismatch"),
		},
		{
			name:  "Wrong checksum",
			input: "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd03",
			err:   errors.New("decoding bech32 failed: checksum failed. Expected h6dd02, got h6dd03."),
		},
		{
			name:  "Invalid public key hash",
			input: "cosmos1vehk7cnpwgls32ra",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CosmosDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("CosmosDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("CosmosDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestCosmosEncodeToString(t *testing.T) {
	keyhash, _ := hex.DecodeString("bc2da90c84049370d1b7c528bc164bc588833f21")
	keyhash2, _ := hex.DecodeString("6e436a571cec916167ba105160474b9c9cd132bd")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:  "Nil",
			input: nil,
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
		{
			name:   "Good",
			input:  keyhash,
			output: "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
		},
		{
			name:   "Good2",
			input:  keyhash2,
			output: "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CosmosEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("CosmosEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("CosmosEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
