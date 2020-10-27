package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestTezosEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
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

	RunTestsEncode(t, slip44.TEZOS, tests)
}

func TestTezosDecodeToString(t *testing.T) {
	keyhash := "06a19f8fb5cea62d147c696afd9a93dbce962f4c8a9c91"
	keyhash2 := "06a1a1a7f2ff4762f8f26aac80221d73be67709dea1d14"
	keyhash3 := "06a1a1a7f2ff4762"

	tests := []TestcaseDecode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Short",
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

	RunTestsDecode(t, slip44.TEZOS, tests)
}
