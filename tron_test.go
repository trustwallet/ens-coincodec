package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestTronEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal",
			input:  "TJRyWwFs9wTFGZg3JbrVriFbNfCug5tDeC",
			output: "415cd0fb0ab3ce40f3051414c604b27756e69e43db",
		},
		{
			name:   "Normal2",
			input:  "THRF3GuPnvvPzKoaT8pJex5XHmo8NNbCb3",
			output: "4151b5659b685047f35498f763dce619c4720d2aa7",
		},
		{
			name:  "Invalid Base58",
			input: "THRF3GuPnvvPzK+aT8pJex5XHmo8NNbCb3",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Too short",
			input: "THRF3GuPnvvPzT8pJex5XHmo8NNbCb3",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Invalid Base58",
			input: "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Too short",
			input: "175tWpb8K1S7NmH4Zx6rewF9WQrcZv245W",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Correct length, but bad checksum",
			input: "THRF3GuPnvvPzKoaT8pJex5XHmo8NNbCc3",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Valid prefix, too short",
			input: "THRF3GuPnvvPzK8pJex5XHmo8NNbCb3",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Valid base 58 and checksum, but prefix is bad",
			input: "Ls2KmCVFo43EBdSeutXQ6hu1Jb5pDhC7RL",
			err:   errors.New("Invalid prefix"),
		},
	}

	RunTestsEncode(t, slip44.TRON, tests)
}

func TestTronDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Empty",
			input: "06a1a1a7f2ff4762",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:   "Good",
			input:  "415cd0fb0ab3ce40f3051414c604b27756e69e43db",
			output: "TJRyWwFs9wTFGZg3JbrVriFbNfCug5tDeC",
		},
		{
			name:   "Good2",
			input:  "4151b5659b685047f35498f763dce619c4720d2aa7",
			output: "THRF3GuPnvvPzKoaT8pJex5XHmo8NNbCb3",
		},
	}

	RunTestsDecode(t, slip44.TRON, tests)
}
