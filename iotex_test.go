package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestIoTexDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j",
			output: "3f9c20bcec9de520d88d98cbe07ee7b5ded0dac4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IoTexDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("IoTexDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("IoTexDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestIoTexEncodeToString(t *testing.T) {
	keyhash, _ := hex.DecodeString("3f9c20bcec9de520d88d98cbe07ee7b5ded0dac4")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "Good",
			input:  keyhash,
			output: "io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IoTexEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("IoTexEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("IoTexEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
