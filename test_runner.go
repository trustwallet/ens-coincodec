package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"
	"strings"
)

type TestcaseEncode struct {
	name    string
	input   string // string
	output  string // encoded as hex string
	err     error
}

type TestcaseDecode struct {
	name    string
	input   string // encoded as hex string
	output  string // string
	err     error
}

func RunTestsEncode(t *testing.T, coinType uint32, tests []TestcaseEncode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBytes(tt.input, coinType)
			var goterror string = "(no error)"
			if err != nil { goterror = err.Error() }
			if tt.err != nil {
				if !strings.HasPrefix(goterror, err.Error()) {
				//if goterror != tt.err.Error() {
					t.Errorf("%v %v: ToBytes() error = %v, wantErr %v", coinType, tt.name, goterror, tt.err)
					return
				}
			} else {
				gothex := hex.EncodeToString(got)
				if !reflect.DeepEqual(gothex, tt.output) {
					t.Errorf("%v %v: ToBytes() = %v, want %v, err: %v", coinType, tt.name, gothex, tt.output, tt.err)
				}
			}
		})
	}
}

func RunTestsDecode(t *testing.T, coinType uint32, tests []TestcaseDecode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoded, err := hex.DecodeString(tt.input)
			if err != nil {
				t.Errorf("%v %v: Preparation error, input is not valid hex string err %v input %v", coinType, tt.name, err, tt.input)
				return
			}
			got, err := ToString(decoded, coinType)
			var goterror string = "(no error)"
			if err != nil { goterror = err.Error() }
			if tt.err != nil {
				if goterror != tt.err.Error() {
					t.Errorf("%v %v: ToString() error = %v, wantErr %v", coinType, tt.name, goterror, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("%v %v: ToString() = %v, want %v, err: %v", coinType, tt.name, got, tt.output, goterror)
				}
			}
		})
	}
}
