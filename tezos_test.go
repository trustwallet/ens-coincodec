package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"
	"strings"

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
			name:  "Too short",
			input: "4cdW",
			err:   errors.New("Base58 string too short"),
		},
		{
			name:  "Too short, len 5 but decoded len is only 3",
			input: "21111",
			err:   errors.New("Base58 string too short"),
		},
		{
			name:  "Valid prefix, too short",
			input: "4cdWcRbbour5VffJdQuSCNG",
			err:   errors.New("Invalid length"),
		},
		{
			name:  "Invalid prefix, valid checksum",
			input: "NmH7tmeJUmHcncBDvpr7aJNEBk7rp5zYsB1qt",
			err:   errors.New("Invalid prefix"),
		},
		{
			name:  "Valid prefix, invalid checksum",
			input: "tz1eZwq8b5cvE2bPKokatLkVMzkxz24z3AAAA",
			err:   errors.New("Bad Base58 checksum: "),
		},
		{
			name:  "Invalid prefix, invalid checksum",
			input: "1tzeZwq8b5cvE2bPKokatLkVMzkxz24zAAAAA",
			err:   errors.New("Bad Base58 checksum: "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TezosDecodeToBytes(tt.input)
			if tt.err != nil {
				goterror := "(no error)"
				if err != nil { goterror = err.Error() }
				if !strings.HasPrefix(goterror, tt.err.Error()) {
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
	keyhash3, _ := hex.DecodeString("06a1a1a7f2ff4762")

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
			name:  "Empty",
			input: keyhash3,
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
