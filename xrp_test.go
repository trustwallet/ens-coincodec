package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/trustwallet/go-primitives/coin"
)

func TestXRPEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal",
			input:  "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			output: "004b4e9c06f24296074f7bc48f92a97916c6dc5ea9",
		},
		{
			name:   "Normal2",
			input:  "X7qvLs7gSnNoKvZzNWUT2e8st17QPY64PPe7zriLNuJszeg",
			output: "05444b4e9c06f24296074f7bc48f92a97916c6dc5ea9000000000000000000",
		},
		{
			name:  "Short",
			input: "123",
			err:   errors.New("base58 decode error: Base58 string too short: 123"),
		},
		{
			name:  "Short with checksum",
			input: "1FuGcFfSmQMU2cMR",
			err:   errors.New("base58 decode error: Bad Base58 checksum"),
		},
		{
			name:  "Bitcoin",
			input: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
			err:   errors.New("base58 decode error: Bad Base58 checksum"),
		},
	}

	RunTestsEncode(t, coin.RIPPLE, tests)
}

func TestXRPDecodeToString(t *testing.T) {
	keyhash := "004b4e9c06f24296074f7bc48f92a97916c6dc5ea9"
	keyhash2 := "05444b4e9c06f24296074f7bc48f92a97916c6dc5ea9000000000000000000"

	tests := []TestcaseDecode{
		{
			name:   "Good",
			input:  keyhash,
			output: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		},
		{
			name:   "Good2",
			input:  keyhash2,
			output: "X7qvLs7gSnNoKvZzNWUT2e8st17QPY64PPe7zriLNuJszeg",
		},
	}

	RunTestsDecode(t, coin.RIPPLE, tests)
}
