package coincodec

import (
	"github.com/btcsuite/btcutil/bech32"
	"github.com/pkg/errors"
)

func Bech32DecodeToBytes(input string, hrp string) ([]byte, error) {
	var empty []byte
	decoded, bytes, err := bech32.Decode(input)
	if err != nil {
		return empty, errors.Wrap(err, "decoding bech32 failed")
	}
	if hrp != decoded {
		return empty, errors.New("decoded hrp mismatch")
	}
	converted, err := bech32.ConvertBits(bytes, 5, 8, false)
	if err != nil {
		return empty, errors.Wrap(err, "converting bits failed")
	}
	return converted, nil
}

func Bech32EncodeToString(bytes []byte, hrp string) (string, error) {
	converted, err := bech32.ConvertBits(bytes, 8, 5, true)
	if err != nil {
		return "", errors.Wrap(err, "encoding bech32 failed")
	}
	return bech32.Encode(hrp, converted)
}
