package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestSolanaEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal1",
			input:  "2gVkYWexTHR5Hb2aLeQN3tnngvWzisFKXDUPrgMHpdST",
			output: "18f9d8d877393bbbe8d697a8a2e52879cc7e84f467656d1cce6bab5a8d2637ec",
		},
		{
			name:   "Normal2",
			input:  "2bUBiBNZyD29gP1oV6de7nxowMLoDBtopMMTGgMvjG5m",
			output: "17b02c16bf792e54b606db6c2b10a24647a3e96215f5450186e183f57caaf0d0",
		},
		{
			name:   "Normal3",
			input:  "H4JcMPicKkHcxxDjkyyrLoQj7Kcibd9t815ak4UvTr9M",
			output: "ee93a4f66f8d16b819bb9beb9ffccdfcdc1412e87fee6a324c2a99a1e0e67148",
		},
		{
			name:   "Normal4",
			input:  "2gVkYWexTHR5Hb2aLeQN3tnngvWzisFKXDUPrgMHpdST",
			output: "18f9d8d877393bbbe8d697a8a2e52879cc7e84f467656d1cce6bab5a8d2637ec",
		},
		{
			name:   "Normal5",
			input:  "2bUBiBNZyD29gP1oV6de7nxowMLoDBtopMMTGgMvjG5m",
			output: "17b02c16bf792e54b606db6c2b10a24647a3e96215f5450186e183f57caaf0d0",
		},
		{
			name:  "Too short",
			input: "2gVkYWexTHR5Hb2aLeQN3tnngvWzisFKXDUPrgMHpd",
			err:   errors.New("Invalid length"),
		},
		{
			name:  "Invalid characters",
			input: "1a+-/=",
			err:   errors.New("Bad Base58 string"),
		},
		{
			name:  "Bitcoin",
			input: "1ES14c7qLb5CYhLMUekctxLgc1FV2Ti9DA",
			err:   errors.New("Invalid length"),
		},
		{
			name:  "Short from ensdomains/address-encoder",
			input: "TUrMmF9Gd4rzrXsQ34ui3Wou94E7HFuJQh",
			err:   errors.New("Invalid length"),
		},
	}

	RunTestsEncode(t, slip44.SOLANA, tests)
}

func TestSolanaDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:   "Good1",
			input:  "18f9d8d877393bbbe8d697a8a2e52879cc7e84f467656d1cce6bab5a8d2637ec",
			output: "2gVkYWexTHR5Hb2aLeQN3tnngvWzisFKXDUPrgMHpdST",
		},
		{
			name:   "Good2",
			input:  "17b02c16bf792e54b606db6c2b10a24647a3e96215f5450186e183f57caaf0d0",
			output: "2bUBiBNZyD29gP1oV6de7nxowMLoDBtopMMTGgMvjG5m",
		},
		{
			name:   "Good3",
			input:  "ee93a4f66f8d16b819bb9beb9ffccdfcdc1412e87fee6a324c2a99a1e0e67148",
			output: "H4JcMPicKkHcxxDjkyyrLoQj7Kcibd9t815ak4UvTr9M",
		},
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Too short",
			input: "18f9d8d877393bbb",
			err:   errors.New("Invalid decoded address length"),
		},
		{
			name:  "Too long",
			input: "18f9d8d877393bbbe8d697a8a2e52879cc7e84f467656d1cce6bab5a8d2637ec7915721bb3",
			err:   errors.New("Invalid decoded address length"),
		},
	}

	RunTestsDecode(t, slip44.SOLANA, tests)
}
