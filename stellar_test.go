package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestStellarEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode {
		{
			name:   "Normal",
			input:  "GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER",
			output: "11b32750d877d03b29df85c0349b3ca5c6abf64786d773323c417f34df0b2fea",
		},
		{
			name:  "Ethereum",
			input: "0x0102030405060708090a0B0c0d0e0f1011121314",
			err:   errors.New("base32 decode error: illegal base32 data at input byte 0"),
		},
		{
			name:  "Version Tag",
			input: "ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY=",
			err:   errors.New("invalid version byte"),
		},
		{
			name:  "Checksum",
			input: "GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6V7DU",
			err:   errors.New("wrong checksum"),
		},
	}

	RunTestsEncode(t, slip44.STELLAR_LUMENS, tests)
}

func TestStellarDecodeToString(t *testing.T) {
	pubkey := "11b32750d877d03b29df85c0349b3ca5c6abf64786d773323c417f34df0b2fea"

	tests := []TestcaseDecode {
		{
			name:   "Good",
			input:  pubkey,
			output: "GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER",
		},
	}

	RunTestsDecode(t, slip44.STELLAR_LUMENS, tests)
}
