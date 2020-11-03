package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestBitcoinEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "P2PKH",
			input:  "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
			output: "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac",
		},
		{
			name:   "P2SH",
			input:  "3Ai1JZ8pdJb2ksieUV8FsxSNVJCpoPi8W6",
			output: "a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1887",
		},
		{
			name:   "P2WPKH",
			input:  "BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
			output: "0014751e76e8199196d454941c45d1b3a323f1433bd6",
		},
		{
			name:   "P2WSH",
			input:  "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
			output: "00201863143c14c5166804bd19203356da136c985678cd4d27a1b8c6329604903262",
		},
		{
			name:  "Litecoin legacy",
			input: "LV7LV7Z4bWDEjYkfx9dQo6k6RjGbXsg6hS",
			err:   errors.New("decoding base58 and bech32 failed: Invalid prefix"),
		},
		{
			name:  "Litecoin segwit",
			input: "ltc1qytnqzjknvv03jwfgrsmzt0ycmwqgl0asjnaxwu",
			err:   errors.New("invalid hrp"),
		},
		{
			name:  "Wrong key hash",
			input: "bc1vehk7cnpwgz0ta92",
			err:   errors.New("wrong witness version"),
		},
		{
			name:  "Empty",
			input: "",
			err:   errors.New("invalid address"),
		},
		{
			name:  "Ethereum",
			input: "0x0102030405060708090a0b0c0d0e0f1011121314",
			err:   errors.New("decoding base58 and bech32 failed"),
		},
	}

	RunTestsEncode(t, slip44.BITCOIN, tests)
}

func TestBitcoinDecodeToString(t *testing.T) {
	script1 := "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac"
	script2 := "a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1887"
	script3 := "0014751e76e8199196d454941c45d1b3a323f1433bd6"
	script4 := "00201863143c14c5166804bd19203356da136c985678cd4d27a1b8c6329604903262"

	tests := []TestcaseDecode{
		{
			name:  "Empty",
			input: "",
			err:   errors.New("invalid data length"),
		},
		{
			name:  "Wrong script",
			input: "00140102",
			err:   errors.New("wrong script data"),
		},
		{
			name:   "P2PKH",
			input:  script1,
			output: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "3Ai1JZ8pdJb2ksieUV8FsxSNVJCpoPi8W6",
		},
		{
			name:   "P2WPKH",
			input:  script3,
			output: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
		},
		{
			name:   "P2WSH",
			input:  script4,
			output: "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
		},
	}

	RunTestsDecode(t, slip44.BITCOIN, tests)
}
