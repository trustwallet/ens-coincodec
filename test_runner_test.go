package coincodec

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

// Some speical test for driving test_runner in error case without actually failing
func TestSpecialTests(t *testing.T) {
	err := RunTestEncode(slip44.BITCOIN, TestcaseEncode {
		name:   "Positive encode case but expecting excpetion",
		input:  "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		err:    errors.New("Error"),
	})
	if err == nil {
		t.Errorf("Expected error error mismatch")
	}
	
	err = RunTestEncode(slip44.BITCOIN, TestcaseEncode {
		name:  "Negative case but expecting output",
		input: "bc1vehk7cnpwgz0ta92",
		output: "0014751e76e8199196d454941c45d1b3a323f1433bd6",
	})
	if err == nil {
		t.Errorf("Expected error error mismatch")
	}

	err = RunTestDecode(slip44.BITCOIN, TestcaseDecode {
		name:   "Positive decode case but expecting excpetion",
		input:  "76a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac",
		err:    errors.New("Error"),
	})
	if err == nil {
		t.Errorf("Expected error error mismatch")
	}

	err = RunTestDecode(slip44.BITCOIN, TestcaseDecode {
		name:  "Negative case but expecting output",
		input: "00140102",
		output: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
	})
	if err == nil {
		t.Errorf("Expected error error mismatch")
	}

	err = RunTestDecode(slip44.BITCOIN, TestcaseDecode {
		name:  "Not hex",
		input: "THIS IS NOT A HEX",
		err:   errors.New("Preparation error"),
	})
	if err == nil {
		t.Errorf("Expected error for not hex")
	}
}
