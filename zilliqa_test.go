package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestZilliqaDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
			output: "7fccacf066a5f26ee3affc2ed1fa9810deaa632c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZilliqaDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("ZilliqaDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("ZilliqaDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestZilliqaEncodeToString(t *testing.T) {
	keyhash, _ := hex.DecodeString("7fccacf066a5f26ee3affc2ed1fa9810deaa632c")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "Good",
			input:  keyhash,
			output: "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZilliqaEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("ZilliqaEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("ZilliqaEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
