package coincodec

import (
	"fmt"
	"math/big"
	"strings"
)

const (
	// Default alphabet, used e.g. by Bitcoin
	Base58DefaultAlphabet = 
	   "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Base58Decode decodes a modified base58 string to a byte slice.
func Base58Decode(b string, alphabet string) ([]byte, error) {
	if len(b) < 5 {
		return nil, fmt.Errorf("Base58 string too short: %s", b)
	}
	answer := big.NewInt(0)
	j := big.NewInt(1)

	for i := len(b) - 1; i >= 0; i-- {
		tmp := strings.IndexAny(alphabet, string(b[i]))
		if tmp == -1 {
			return nil, fmt.Errorf("Bad Base58 string: %s", b)
		}
		idx := big.NewInt(int64(tmp))
		tmp1 := big.NewInt(0)
		tmp1.Mul(j, idx)

		answer.Add(answer, tmp1)
		j.Mul(j, bigRadix)
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabet[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen)
	copy(val[numZeros:], tmpval)

	return val, nil
}

// Base58Encode encodes a byte slice to a modified base58 string.
func Base58Encode(b []byte, alphabet string) string {
	if b == nil || len(b) == 0 {
		return ""
	}
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabet[0])
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
