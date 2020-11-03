package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestZCashEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal",
			input:  "t1RygJmrLdNGgi98gUgEJDTVaELTAYWoMBy",
			output: "76a91458e71790e51ab7558c05a6067cfc4926aa8c44dd88ac",
		},
		{
			name:   "Normal2",
			input:  "t1TWk2mmvESDnE4dmCfT7MQ97ij6ZqLpNVU",
			output: "76a91469bf38acef973293c07f05c778eb1209748e8d5288ac",
		},
		{
			name:   "Normal3",
			input:  "t3RD6RFKhWSotNbPEY4Vw7Ku9QCfKkzrbBL",
			output: "a91448e71790e51ab7558c05a6067cfc4926aa8c44dd87",
		},
		{
			name:   "Normal4",
			input:  "t1Wg9uPPAfwhBWeRjtDPa5ZHNzyBx9rJVKY",
			output: "76a9148c6f453157897ce2e6de413f329d995fe0d8f90288ac",
		},
		{
			name:   "Normal5",
			input:  "t1gaySCXCYtXE3ygP38YuWtVZczsEbdjG49",
			output: "76a914f925b59e1dc043ad7f0b7e85ea05b06dfe83413888ac",
		},
		{
			name:  "Invalid Base58",
			input: "t1RygJmrLdNGgi98+UgEJDTVaELTAYWoMBy",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Too short",
			input: "t1RygJmrLdNGgi98gUgEJDTVaELTAYW",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Correct length, but bad checksum",
			input: "t1RygJmrLdNGgi98gUgEJDTVaELTAYWoMBz",
			err:   errors.New("Bad Base58 checksum"),
		},
		{
			name:  "Valid base58 but too short",
			input: "TJRyWwFs9wTFGZg3JbrVriFbNfCug5tDeC",
			err:   errors.New("Invalid decoded length"),
		},
		{
			name:  "Valid base 58 and checksum, but prefix is bad",
			input: "2NRbuP5YfzRNEa1RibT5kXay1VgvQHnydZY1",
			err:   errors.New("Invalid prefix"),
		},
	}

	RunTestsEncode(t, slip44.ZCASH, tests)
}

func TestZCashDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Invalid opcode bytes"),
		},
		{
			name:  "Short",
			input: "06a1a1a7f2ff4762",
			err:   errors.New("Invalid opcode bytes"),
		},
		{
			name:   "Good",
			input:  "76a91458e71790e51ab7558c05a6067cfc4926aa8c44dd88ac",
			output: "t1RygJmrLdNGgi98gUgEJDTVaELTAYWoMBy",
		},
		{
			name:   "Good2",
			input:  "76a91469bf38acef973293c07f05c778eb1209748e8d5288ac",
			output: "t1TWk2mmvESDnE4dmCfT7MQ97ij6ZqLpNVU",
		},
		{
			name:   "Good3",
			input:  "76a9148c6f453157897ce2e6de413f329d995fe0d8f90288ac",
			output: "t1Wg9uPPAfwhBWeRjtDPa5ZHNzyBx9rJVKY",
		},
		{
			name:   "Good5",
			input:  "76a914f925b59e1dc043ad7f0b7e85ea05b06dfe83413888ac",
			output: "t1gaySCXCYtXE3ygP38YuWtVZczsEbdjG49",
		},
	}

	RunTestsDecode(t, slip44.ZCASH, tests)
}
