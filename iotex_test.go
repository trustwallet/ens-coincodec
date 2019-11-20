package coincodec

import (
	"testing"

	"github.com/wealdtech/go-slip44"
)

func TestIoTexEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode {
		{
			name:   "Normal",
			input:  "io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j",
			output: "3f9c20bcec9de520d88d98cbe07ee7b5ded0dac4",
		},
	}

	RunTestsEncode(t, slip44.IOTEX, tests)
}

func TestIoTexDecodeToString(t *testing.T) {
	keyhash := "3f9c20bcec9de520d88d98cbe07ee7b5ded0dac4"

	tests := []TestcaseDecode {
		{
			name:   "Good",
			input:  keyhash,
			output: "io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j",
		},
	}

	RunTestsDecode(t, slip44.IOTEX, tests)
}
