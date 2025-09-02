package coincodec

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/trustwallet/go-primitives/coin"

	"github.com/cpacia/bchutil"
)

const (
	hrpBCH = "bitcoincash"
)

func init() {
	toBytesMap[coin.BITCOINCASH] = BitcoinCashDecodeToBytes
	toStringMap[coin.BITCOINCASH] = BitcoinCashEncodeToString
}

// BitcoinCashDecodeToBytes converts the input string to a byte array
func BitcoinCashDecodeToBytes(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid address")
	}
	// try cashaddr first
	decoded, hrp, addrType, err := bchutil.CheckDecodeCashAddress(input)
	if err != nil {
		// try base58
		return BitcoinDecodeToBytes(input)
	}
	if hrp != hrpBCH {
		return nil, errors.New("invalid hrp")
	}
	if addrType == bchutil.P2PKH {
		return buildP2PKHScript(decoded), nil
	} else if addrType == bchutil.P2SH {
		return buildP2SHScript(decoded), nil
	} else {
		return nil, errors.New("unknown address type")
	}
}

// BitcoinCashEncodeToString converts the input byte array to a string representation of the Bitcoin address.
func BitcoinCashEncodeToString(input []byte) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid data length")
	}
	if bytes.HasPrefix(input, P2PKH_SCRIPT_PREFIX) {
		address := bchutil.CheckEncodeCashAddress(input[3:len(input)-2], hrpBCH, bchutil.P2PKH)
		return fmt.Sprintf("%s:%s", hrpBCH, address), nil
	} else if bytes.HasPrefix(input, P2SH_SCRIPT_PREFIX) {
		address := bchutil.CheckEncodeCashAddress(input[2:len(input)-1], hrpBCH, bchutil.P2SH)
		return fmt.Sprintf("%s:%s", hrpBCH, address), nil
	}
	return "", errors.New("wrong script data")
}
