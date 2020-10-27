package coincodec

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type TestcaseEncode struct {
	name   string
	input  string // string
	output string // encoded as hex string
	err    error
}

type TestcaseDecode struct {
	name   string
	input  string // encoded as hex string
	output string // string
	err    error
}

func errorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return "(no error)"
}

func RunTestsEncode(t *testing.T, coinType uint32, tests []TestcaseEncode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RunTestEncode(coinType, tt)
			if err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func RunTestEncode(coinType uint32, tt TestcaseEncode) error {
	testfun, _ := toBytesMap[coinType]

	got, err := testfun(tt.input)

	goterror := errorString(err)
	if tt.err != nil {
		if !strings.HasPrefix(goterror, tt.err.Error()) {
			return fmt.Errorf("%v %v: ToBytes() error = %v, wantErr %v", coinType, tt.name, goterror, tt.err)
		}
	} else {
		gothex := hex.EncodeToString(got)
		if !reflect.DeepEqual(gothex, tt.output) {
			return fmt.Errorf("%v %v: ToBytes() = %v, err: %v, want %v, err: %v", coinType, tt.name, gothex, goterror, tt.output, tt.err)
		}
	}
	return nil
}

func RunTestsDecode(t *testing.T, coinType uint32, tests []TestcaseDecode) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RunTestDecode(coinType, tt)
			if err != nil {
				t.Error(err.Error())
			}
		})
	}
}

func RunTestDecode(coinType uint32, tt TestcaseDecode) error {
	decoded, err := hex.DecodeString(tt.input)
	if err != nil {
		return fmt.Errorf("%v %v: Preparation error, input is not valid hex string err %v input %v", coinType, tt.name, err, tt.input)
	}
	testfun, _ := toStringMap[coinType]

	got, err := testfun(decoded)

	goterror := errorString(err)
	if tt.err != nil {
		if goterror != tt.err.Error() {
			return fmt.Errorf("%v %v: ToString() error = %v, wantErr %v", coinType, tt.name, goterror, tt.err)
		}
	} else {
		if got != tt.output {
			return fmt.Errorf("%v %v: ToString() = %v, want %v, err: %v", coinType, tt.name, got, tt.output, goterror)
		}
	}
	return nil
}
