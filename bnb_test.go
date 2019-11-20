package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestBNBEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode {
		{
			name:   "Normal",
			input:  "bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46h2",
			output: "40c2979694bbc961023d1d27be6fc4d21a9febe6",
		},
		{
			name:  "Testnet",
			input: "tbnb1683m5xkxa3p69yagrpdsvkp94xvx6utjf0cs9t",
			err:   errors.New("decoded hrp mismatch"),
		},
		{
			name:  "Wrong checksum",
			input: "bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46hy",
			err:   errors.New("decoding bech32 failed: checksum failed. Expected fn46h2, got fn46hy."),
		},
		{
			name:  "Invalid public key hash",
			input: "bnb1vehk7cnpwga5a9et",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
	}

	RunTestsEncode(t, slip44.BINANCE, tests)
}

func TestBNBDecodeToString(t *testing.T) {
	keyhash := "40c2979694bbc961023d1d27be6fc4d21a9febe6"
	tests := []TestcaseDecode {
		{
			name:  "Empty",
			input: "",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
		{
			name:  "Too short",
			input: "0102030405",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
		{
			name:   "Good",
			input:  keyhash,
			output: "bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46h2",
		},
	}

	RunTestsDecode(t, slip44.BINANCE, tests)
}
