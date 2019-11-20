package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestTezosDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "tz1Yju7jmmsaUiG9qQLoYv35v5pHgnWoLWbt",
			output: "06a19f8fb5cea62d147c696afd9a93dbce962f4c8a9c91",
		},
		{
			name:   "Normal2",
			input:  "tz2PdGc7U5tiyqPgTSgqCDct94qd6ovQwP6u",
			output: "06a1a1a7f2ff4762f8f26aac80221d73be67709dea1d14",
		},
		{
			name:   "Normal3",
			input:  "tz3VEZ4k6a4Wx42iyev6i2aVAptTRLEAivNN",
			output: "06a1a461af383a78291ace2dea59d3da6c9a8b1cdb1b96",
		},
		{
			name:  "Invalid prefix, valid checksum",
			input: "NmH7tmeJUmHcncBDvpr7aJNEBk7rp5zYsB1qt",
			output: "a1069f88226e21ee0e4f0eda850d6d28c2ec992c3d9dfe",
		},
		{
			name:  "Valid prefix, invalid checksum",
			input: "tz1eZwq8b5cvE2bPKokatLkVMzkxz24z3AAAA",
			err:   errors.New("Bad Base58 checksum: [158 75 81 137 45 108 126 36 221 126 132 242 217 189 40 8 103 58 112 225 175 69 246 127 65 144 45 244 240 65 32 147] expected [52 175 46 191]"),
		},
		{
			name:  "Invalid prefix, invalid checksum",
			input: "1tzeZwq8b5cvE2bPKokatLkVMzkxz24zAAAAA",
			err:   errors.New("Bad Base58 checksum: [165 203 237 246 198 234 235 145 5 193 229 5 152 68 252 22 56 193 215 76 132 138 239 45 248 112 249 192 35 164 195 55] expected [57 103 234 47]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TezosDecodeToBytes(tt.input)
			if tt.err != nil {
				goterror := "(no error)"
				if err != nil { goterror = err.Error() }
				if goterror != tt.err.Error() {
					t.Errorf("TezosDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("TezosDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestTezosEncodeToString(t *testing.T) {
	keyhash, _ := hex.DecodeString("06a19f8fb5cea62d147c696afd9a93dbce962f4c8a9c91")
	keyhash2, _ := hex.DecodeString("06a1a1a7f2ff4762f8f26aac80221d73be67709dea1d14")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:  "Nil",
			input: nil,
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Empty",
			input: []byte{},
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:   "Good",
			input:  keyhash,
			output: "tz1Yju7jmmsaUiG9qQLoYv35v5pHgnWoLWbt",
		},
		{
			name:   "Good2",
			input:  keyhash2,
			output: "tz2PdGc7U5tiyqPgTSgqCDct94qd6ovQwP6u",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TezosEncodeToString(tt.input)
			if tt.err != nil {
				goterror := "(no error)"
				if err != nil { goterror = err.Error() }
				if goterror != tt.err.Error() {
					t.Errorf("TezosEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("TezosEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
