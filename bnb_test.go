package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestBNBDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46h2",
			output: "40c2979694bbc961023d1d27be6fc4d21a9febe6",
		},
		{
			name:  "Testnet",
			input: "tbnb1683m5xkxa3p69yagrpdsvkp94xvx6utjf0cs9t",
			err:   errors.New("decoded hrp mismatch"),
		},
		{
			name:  "Wrong checksum",
			input: "bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46hy",
			err:   errors.New("decoding bech32 failed: checksum failed. Expected fn46h2, got fn46hy."),
		},
		{
			name:  "Invalid public key hash",
			input: "bnb1vehk7cnpwga5a9et",
			err:   errors.New("BNB address key hash must be 20 bytes"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BNBDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BNBDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("BNBDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestBNBEncodeToString(t *testing.T) {
	keyhash, _ := hex.DecodeString("40c2979694bbc961023d1d27be6fc4d21a9febe6")
	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:  "Nil",
			input: nil,
			err:   errors.New("BNB address key hash must be 20 bytes"),
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   errors.New("BNB address key hash must be 20 bytes"),
		},
		{
			name:   "Good",
			input:  keyhash,
			output: "bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46h2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BNBEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BNBEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("BNBEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
