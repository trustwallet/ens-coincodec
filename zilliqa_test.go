package coincodec

import (
	"testing"

	"github.com/trustwallet/go-primitives/coin"
)

func TestZilliqaEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal",
			input:  "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
			output: "7fccacf066a5f26ee3affc2ed1fa9810deaa632c",
		},
	}

	RunTestsEncode(t, coin.ZILLIQA, tests)
}

func TestZilliqaDecodeToString(t *testing.T) {
	keyhash := "7fccacf066a5f26ee3affc2ed1fa9810deaa632c"

	tests := []TestcaseDecode{
		{
			name:   "Good",
			input:  keyhash,
			output: "zil10lx2eurx5hexaca0lshdr75czr025cevqu83uz",
		},
	}

	RunTestsDecode(t, coin.ZILLIQA, tests)
}
