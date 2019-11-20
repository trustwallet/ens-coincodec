package coincodec

import (
	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

const (
	xrpAlphabet = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
)

func init() {
	toBytesMap[slip44.RIPPLE] = XRPDecodeToBytes
	toStringMap[slip44.RIPPLE] = XRPEncodeToString
}

// XRPDecodeToBytes converts the input string to a byte array
func XRPDecodeToBytes(input string) ([]byte, error) {
	decoded, err := Base58Decode(input, xrpAlphabet)
	if err != nil {
		return nil, errors.Wrap(err, "base58 decode error")
	}
	return decoded, nil
}

// XRPEncodeToString converts the input byte array to a string representation of the XRP address.
func XRPEncodeToString(bytes []byte) (string, error) {
	return Base58Encode(bytes, xrpAlphabet), nil
}
