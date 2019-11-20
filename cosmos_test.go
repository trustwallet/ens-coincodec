package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestCosmosEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode {
		{
			name:   "Normal",
			input:  "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
			output: "bc2da90c84049370d1b7c528bc164bc588833f21",
		},
		{
			name:   "Normal2",
			input:  "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
			output: "6e436a571cec916167ba105160474b9c9cd132bd",
		},
		{
			name:  "Testnet",
			input: "cosmosvaloper1sxx9mszve0gaedz5ld7qdkjkfv8z992ax69k08",
			err:   errors.New("decoded hrp mismatch"),
		},
		{
			name:  "Wrong checksum",
			input: "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd03",
			err:   errors.New("decoding bech32 failed: checksum failed. Expected h6dd02, got h6dd03."),
		},
		{
			name:  "Invalid public key hash",
			input: "cosmos1vehk7cnpwgls32ra",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
	}

	RunTestsEncode(t, slip44.ATOM, tests)
}

func TestCosmosDecodeToString(t *testing.T) {
	keyhash := "bc2da90c84049370d1b7c528bc164bc588833f21"
	keyhash2 := "6e436a571cec916167ba105160474b9c9cd132bd"

	tests := []TestcaseDecode {
		{
			name:  "Empty",
			input: "",
			err:   errors.New("empty input"),
		},
		{
			name:  "Too short",
			input: "0102030405",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
		{
			name:   "Good",
			input:  keyhash,
			output: "cosmos1hsk6jryyqjfhp5dhc55tc9jtckygx0eph6dd02",
		},
		{
			name:   "Good2",
			input:  keyhash2,
			output: "cosmos1depk54cuajgkzea6zpgkq36tnjwdzv4afc3d27",
		},
	}

	RunTestsDecode(t, slip44.ATOM, tests)
}
