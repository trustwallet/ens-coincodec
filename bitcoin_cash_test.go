package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestBitcoinCashEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode {
		{
			name:   "P2PKH",
			input:  "1BpEi6DfDAUFd7GtittLSdBeYJvcoaVggu",
			output: "76a91476a04053bda0a88bda5177b86a15c3b29f55987388ac",
		},
		{
			name:   "P2PKH CashAddr",
			input:  "bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a",
			output: "76a91476a04053bda0a88bda5177b86a15c3b29f55987388ac",
		},
		{
			name:   "P2SH",
			input:  "3CWFddi6m4ndiGyKqzYvsFYagqDLPVMTzC",
			output: "a91476a04053bda0a88bda5177b86a15c3b29f55987387",
		},
		{
			name:   "P2SH CashAddr",
			input:  "bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq",
			output: "a91476a04053bda0a88bda5177b86a15c3b29f55987387",
		},
		{
			name:  "Empty",
			input: "",
			err:   errors.New("empty input"),
		},
		{
			name:  "Testnet",
			input: "bchtest:qpm2qsznhks23z7629mms6s4cwef74vcwvqcw003ap",
			err:   errors.New("invalid hrp"),
		},
	}

	RunTestsEncode(t, slip44.BITCOIN_CASH, tests)
}

func TestBitcoinCashDecodeToString(t *testing.T) {
	script1 := "76a914cb481232299cd5743151ac4b2d63ae198e7bb0a988ac"
	script2 := "a914cb481232299cd5743151ac4b2d63ae198e7bb0a987"

	tests := []TestcaseDecode {
		{
			name:  "Empty",
			input: "",
			err:   errors.New("empty input"),
		},
		{
			name:  "Wrong script",
			input: "00140102",
			err:   errors.New("wrong script data"),
		},
		{
			name:   "P2PKH",
			input:  script1,
			output: "bitcoincash:qr95sy3j9xwd2ap32xkykttr4cvcu7as4y0qverfuy",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "bitcoincash:pr95sy3j9xwd2ap32xkykttr4cvcu7as4yc93ky28e",
		},
	}

	RunTestsDecode(t, slip44.BITCOIN_CASH, tests)
}
