package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestBitcoinDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
			output: "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac",
		},
		{
			name:   "P2SH",
			input:  "3Ai1JZ8pdJb2ksieUV8FsxSNVJCpoPi8W6",
			output: "a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1887",
		},
		{
			name:   "P2WPKH",
			input:  "BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
			output: "0014751e76e8199196d454941c45d1b3a323f1433bd6",
		},
		{
			name:   "P2WSH",
			input:  "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
			output: "00201863143c14c5166804bd19203356da136c985678cd4d27a1b8c6329604903262",
		},
		{
			name:  "Litecoin legacy",
			input: "LV7LV7Z4bWDEjYkfx9dQo6k6RjGbXsg6hS",
			err:   errors.New("invalid address prefix"),
		},
		{
			name:  "Litecoin segwit",
			input: "ltc1qytnqzjknvv03jwfgrsmzt0ycmwqgl0asjnaxwu",
			err:   errors.New("invalid hrp"),
		},
		{
			name:  "Wrong key hash",
			input: "bc1vehk7cnpwgz0ta92",
			err:   errors.New("wrong witness version"),
		},
		{
			name:  "Empty",
			input: "",
			err:   errors.New("invalid address"),
		},
		{
			name:  "Ethereum",
			input: "0x0102030405060708090a0b0c0d0e0f1011121314",
			err:   errors.New("decoding bech32 failed: invalid index of 1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinDecodeToBytes(tt.input)
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

func TestBitcoinEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac")
	script2, _ := hex.DecodeString("a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1887")
	script3, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	script4, _ := hex.DecodeString("00201863143c14c5166804bd19203356da136c985678cd4d27a1b8c6329604903262")

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
			output: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "3Ai1JZ8pdJb2ksieUV8FsxSNVJCpoPi8W6",
		},
		{
			name:   "P2WPKH",
			input:  script3,
			output: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
		},
		{
			name:   "P2WSH",
			input:  script4,
			output: "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinEncodeToString(tt.input)
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
