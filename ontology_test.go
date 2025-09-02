package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/trustwallet/go-primitives/coin"
)

func TestOntologyEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal",
			input:  "ANDfjwrUroaVtvBguDtrWKRMyxFwvVwnZD",
			output: "1746b1a18af6b7c9f8a4602f9f73eeb3030f0c29b7",
		},
		{
			name:   "Normal2",
			input:  "AeicEjZyiXKgUeSBbYQHxsU1X3V5Buori5",
			output: "17fbacc8214765d457c8e3f2b5a1d3c4981a2e9d2a",
		},
		{
			name:   "Normal3",
			input:  "AYTxeseHT5khTWhtWX1pFFP1mbQrd4q1zz",
			output: "17b716d488862fedd488a4616cfc0068bb6a6c849f",
		},
		{
			name:  "Invalid Base58",
			input: "ANDfjwrUroaVtvBg+DtrWKRMyxFwvVwnZD",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Too short",
			input: "ANDfjwrUroaVtvBguDtrWKRMyxFwvVwn",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Invalid Base58",
			input: "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Too short",
			input: "ANDfjwrUroaVtvBguDtrWKRMy",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Correct length, but bad checksum",
			input: "ANDfjwrUroaVtvBguDtrWKRMyxFwvVwnZE",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Valid version, too short",
			input: "ANDfjwrUroaVtvBguDtrWKRMyxFwvV",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Valid base 58 and checksum, but version is bad",
			input: "GyscQcQtoyAiYSwGuELv1FjbqfXwwQPQ8Z",
			err:   errors.New("Invalid prefix"),
		},
		{
			name:  "Bad checksum",
			input: "AATxeseHT5khTWhtWX1pFFP1mbQrd4q1zz",
			err:   errors.New("Bad Base58 checksum"),
		},
	}

	RunTestsEncode(t, coin.ONTOLOGY, tests)
}

func TestOntologyDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Invalid decoded length"),
		},
		{
			name:  "Empty",
			input: "06a1a1a7f2ff4762",
			err:   errors.New("Invalid decoded length"),
		},
		{
			name:   "Good",
			input:  "1746b1a18af6b7c9f8a4602f9f73eeb3030f0c29b7",
			output: "ANDfjwrUroaVtvBguDtrWKRMyxFwvVwnZD",
		},
		{
			name:   "Good2",
			input:  "17fbacc8214765d457c8e3f2b5a1d3c4981a2e9d2a",
			output: "AeicEjZyiXKgUeSBbYQHxsU1X3V5Buori5",
		},
		{
			name:   "Good3",
			input:  "17b716d488862fedd488a4616cfc0068bb6a6c849f",
			output: "AYTxeseHT5khTWhtWX1pFFP1mbQrd4q1zz",
		},
		{
			name:  "Invalid Prefix",
			input: "415cd0fb0ab3ce40f3051414c604b27756e69e43db",
			err:   errors.New("Invalid prefix"),
		},
		{
			name:  "Invalid Prefix2",
			input: "4151b5659b685047f35498f763dce619c4720d2aa7",
			err:   errors.New("Invalid prefix"),
		},
	}

	RunTestsDecode(t, coin.ONTOLOGY, tests)
}
