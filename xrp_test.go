package coincodec

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestXRPDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			output: "004b4e9c06f24296074f7bc48f92a97916c6dc5ea9",
		},
		{
			name:   "Normal2",
			input:  "X7qvLs7gSnNoKvZzNWUT2e8st17QPY64PPe7zriLNuJszeg",
			output: "05444b4e9c06f24296074f7bc48f92a97916c6dc5ea9000000000000000000",
		},
		{
			name:  "Short",
			input: "123",
			err:   errors.New("base58 decode error: Base58 string too short: 123"),
		},
		{
			name:  "Short with checksum",
			input: "1FuGcFfSmQMU2cMR",
			err:   errors.New("base58 decode error: Bad Base58 checksum"),
		},
		{
			name:  "Bitcoin",
			input: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
			err:   errors.New("base58 decode error: Bad Base58 checksum"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XRPDecodeToBytes(tt.input)
			if tt.err != nil {
				if !strings.HasPrefix(err.Error(), tt.err.Error()) {
					t.Errorf("XRPDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("XRPDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestXRPEncodeToString(t *testing.T) {
	keyhash, _ := hex.DecodeString("004b4e9c06f24296074f7bc48f92a97916c6dc5ea9")
	keyhash2, _ := hex.DecodeString("05444b4e9c06f24296074f7bc48f92a97916c6dc5ea9000000000000000000")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "Good",
			input:  keyhash,
			output: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		},
		{
			name:   "Good2",
			input:  keyhash2,
			output: "X7qvLs7gSnNoKvZzNWUT2e8st17QPY64PPe7zriLNuJszeg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := XRPEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("XRPEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("XRPEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
