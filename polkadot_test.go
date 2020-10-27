package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestPolkadotEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal1",
			input:  "15KRsCq9LLNmCxNFhGk55s5bEyazKefunDxUH24GFZwsTxyu",
			output: "beff0e5d6f6e6e6d573d3044f3e2bfb353400375dc281da3337468d4aa527908",
		},
		{
			name:   "Normal2",
			input:  "15AeCjMpcSt3Fwa47jJBd7JzQ395Kr2cuyF5Zp4UBf1g9ony",
			output: "b84b605a51045e43740bc74db62d0077ff9d971e49d588f9b307915721bb3251",
		},
		{
			name:   "Normal3 from ensdomains/address-encoder",
			input:  "1FRMM8PEiWXYax7rpS6X4XZX1aAAxSWx1CrKTyrVYhV24fg",
			output: "0aff6865635ae11013a83835c019d44ec3f865145943f487ae82a8e7bed3a66b",
		},
		{
			name:  "Too short",
			input: "15KR",
			err:   errors.New("Base58 string too short"),
		},
		{
			name:  "Invalid characters",
			input: "1a+-/=",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Bad checksum",
			input: "15KRsCq9LLNmCxNFhGk55s5bEyazKefunDxUH24GFZwsTwoF",
			err:   errors.New("Invalid checksum"),
		},
		{
			name:  "Substrate ed25519",
			input: "5FqqU2rytGPhcwQosKRtW1E3ha6BJKAjHgtcodh71dSyXhoZ",
			err:   errors.New("Invalid network"),
		},
		{
			name:  "Bitcoin",
			input: "1ES14c7qLb5CYhLMUekctxLgc1FV2Ti9DA",
			err:   errors.New("Invalid length"),
		},
		{
			name:  "Kusama ed25519",
			input: "FHKAe66mnbk8ke8zVWE9hFVFrJN1mprFPVmD5rrevotkcDZ",
			err:   errors.New("Invalid network"),
		},
		{
			name:  "Kusama secp256k1",
			input: "FxQFyTorsjVsjjMyjdgq8w5vGx8LiA1qhWbRYcFijxKKchx",
			err:   errors.New("Invalid network"),
		},
		{
			name:  "Kusama sr25519",
			input: "EJ5UJ12GShfh7EWrcNZFLiYU79oogdtXFUuDDZzk7Wb2vCe",
			err:   errors.New("Invalid network"),
		},
	}

	RunTestsEncode(t, slip44.POLKADOT, tests)
}

func TestPolkadotDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:   "Good1",
			input:  "beff0e5d6f6e6e6d573d3044f3e2bfb353400375dc281da3337468d4aa527908",
			output: "15KRsCq9LLNmCxNFhGk55s5bEyazKefunDxUH24GFZwsTxyu",
		},
		{
			name:   "Good2",
			input:  "b84b605a51045e43740bc74db62d0077ff9d971e49d588f9b307915721bb3251",
			output: "15AeCjMpcSt3Fwa47jJBd7JzQ395Kr2cuyF5Zp4UBf1g9ony",
		},
		{
			name:   "Good3 from ensdomains/address-encoder",
			input:  "0aff6865635ae11013a83835c019d44ec3f865145943f487ae82a8e7bed3a66b",
			output: "1FRMM8PEiWXYax7rpS6X4XZX1aAAxSWx1CrKTyrVYhV24fg",
		},
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Too short",
			input: "06a1a1a7f2ff4762",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Too long",
			input: "06a1a1a7f2ff476205a51045e43740bc74db62d0077ff9d971e49d588f9b307915721bb3",
			err:   errors.New("Invalid decoded address length"),
		},
	}

	RunTestsDecode(t, slip44.POLKADOT, tests)
}
