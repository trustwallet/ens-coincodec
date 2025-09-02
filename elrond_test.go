package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/trustwallet/go-primitives/coin"
)

func TestElrondEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Alice",
			input:  "erd1l453hd0gt5gzdp7czpuall8ggt2dcv5zwmfdf3sd3lguxseux2fsmsgldz",
			output: "fd691bb5e85d102687d81079dffce842d4dc328276d2d4c60d8fd1c3433c3293",
		},
		{
			name:   "Bob",
			input:  "erd1cux02zersde0l7hhklzhywcxk4u9n4py5tdxyx7vrvhnza2r4gmq4vw35r",
			output: "c70cf50b238372fffaf7b7c5723b06b57859d424a2da621bcc1b2f317543aa36",
		},
		{
			name:   "Carol",
			input:  "erd19nu5t7hszckwah5nlcadmk5rlchtugzplznskffpwecygcu0520s9tnyy0",
			output: "2cf945faf0162ceede93fe3addda83fe2ebe2041f8a70b2521767044638fa29f",
		},
		{
			name:   "ensdomains",
			input:  "erd1qdzvfpa7gqjsnfhdxhvcp2mlysc80uz60yjhxre3lwl00q0jd4nqgauy9q",
			output: "0344c487be402509a6ed35d980ab7f243077f05a7925730f31fbbef781f26d66",
		},
		// invalid
		{
			name:  "Empty",
			input: "",
			err:   errors.New("decoding bech32 failed"),
		},
		{
			name:  "Invalid",
			input: "foo",
			err:   errors.New("decoding bech32 failed"),
		},
		{
			name:  "Invalid 2",
			input: "10z9xdugayn528ksaesdwlhf006fw5sg2qmmm0h52fvxczwgesyvq5pwemr",
			err:   errors.New("decoding bech32 failed"),
		},
		{
			name:  "Invalid 3",
			input: "xerd10z9xdugayn528ksaesdwlhf006fw5sg2qmmm0h52fvxczwgesyvq5pwemr",
			err:   errors.New("decoding bech32 failed"),
		},
		{
			name:  "Invalid 4",
			input: "foo10z9xdugayn528ksaesdwlhf006fw5sg2qmmm0h52fvxczwgesyvq5pwemr",
			err:   errors.New("decoding bech32 failed"),
		},
		{
			name:  "Invalid Hex",
			input: "fd691bb5e85d102687d81079dffce842d4dc328276d2d4c60d8fd1c3433c3293",
			err:   errors.New("decoding bech32 failed"),
		},
	}

	RunTestsEncode(t, coin.ELROND, tests)
}

func TestElrondDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:   "Alice",
			input:  "fd691bb5e85d102687d81079dffce842d4dc328276d2d4c60d8fd1c3433c3293",
			output: "erd1l453hd0gt5gzdp7czpuall8ggt2dcv5zwmfdf3sd3lguxseux2fsmsgldz",
		},
		{
			name:   "Bob",
			input:  "c70cf50b238372fffaf7b7c5723b06b57859d424a2da621bcc1b2f317543aa36",
			output: "erd1cux02zersde0l7hhklzhywcxk4u9n4py5tdxyx7vrvhnza2r4gmq4vw35r",
		},
		{
			name:   "Carol",
			input:  "2cf945faf0162ceede93fe3addda83fe2ebe2041f8a70b2521767044638fa29f",
			output: "erd19nu5t7hszckwah5nlcadmk5rlchtugzplznskffpwecygcu0520s9tnyy0",
		},
		// invalid
		{
			name:  "Empty",
			input: "",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
		{
			name:  "Too short",
			input: "fd691bb5e85d102687d81079dffce842d4dc328276d2d4c60d8fd1c3433c32",
			err:   errors.New("A Bech32 address key hash must be 20 bytes"),
		},
	}

	RunTestsDecode(t, coin.ELROND, tests)
}
