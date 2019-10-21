package coincodec

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/wealdtech/go-slip44"
)

const (
	BTC_MAINNET_HRP = "bc"

	BTC_P2PKH_PREFIX   = 0x00
	BTC_P2SH_Prefix    = 0x05
	BTC_KEYHASH_LENGTH = 0x14

	BTC_WITNESS_VERSION = 0x00
)

const (
	OP_DUP         = 0x76
	OP_HASH160     = 0xa9
	OP_EQUALVERIFY = 0x88
	OP_CHECKSIG    = 0xac
	OP_EQUAL       = 0x87
)

var (
	P2PKH_SCRIPT_PREFIX = []byte{OP_DUP, OP_HASH160, BTC_KEYHASH_LENGTH}
	P2SH_SCRIPT_PREFIX  = []byte{OP_HASH160, BTC_KEYHASH_LENGTH}
)

func init() {
	toBytesMap[slip44.BITCOIN] = BitcoinDecodeToBytes
	toStringMap[slip44.BITCOIN] = BitcoinEncodeToString
}

// BitcoinDecodeToBytes converts the input string to a byte array
func BitcoinDecodeToBytes(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid address")
	}
	// try base58 first
	bytes, version, err := base58.CheckDecode(input)
	if err != nil {
		// try bech32
		hrp, bytes, err := bech32.Decode(input)
		if err != nil {
			return nil, errors.Wrapf(err, "decoding bech32 failed")
		}
		if hrp != BTC_MAINNET_HRP {
			return nil, errors.New("invalid hrp")
		}
		return buildWitnessScript(bytes)
	} else {
		// check data length
		if len(bytes) != BTC_KEYHASH_LENGTH {
			return nil, errors.New("invalid data length")
		}
		// check version byte
		if version == BTC_P2PKH_PREFIX {
			return buildP2PKHScript(bytes), nil
		} else if version == BTC_P2SH_Prefix {
			return buildP2SHScript(bytes), nil
		}
		return nil, errors.New("invalid address prefix")
	}
}

// BitcoinEncodeToString converts the input byte array to a string representation of the Bitcoin address.
func BitcoinEncodeToString(input []byte) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid data length")
	}
	if bytes.HasPrefix(input, P2PKH_SCRIPT_PREFIX) {
		return base58.CheckEncode(input[3:len(input)-2], BTC_P2PKH_PREFIX), nil
	} else if bytes.HasPrefix(input, P2SH_SCRIPT_PREFIX) {
		return base58.CheckEncode(input[2:len(input)-1], BTC_P2SH_Prefix), nil
	} else if input[0] == BTC_WITNESS_VERSION && len(input) > 2 {
		if int(input[1]) != len(input)-2 {
			return "", errors.New("wrong script data")
		}
		converted, err := bech32.ConvertBits(input[2:], 8, 5, true)
		if err != nil {
			return "", errors.Wrap(err, "converting bits failed")
		}
		data := []byte{BTC_WITNESS_VERSION}
		data = append(data, converted...)
		return bech32.Encode(BTC_MAINNET_HRP, data)
	}
	return "", errors.New("invalid bytes")
}

func buildP2PKHScript(bytes []byte) []byte {
	var script []byte
	suffix := []byte{OP_EQUALVERIFY, OP_CHECKSIG}

	script = append(script, P2PKH_SCRIPT_PREFIX...)
	script = append(script, bytes...)
	script = append(script, suffix...)
	return script
}

func buildP2SHScript(bytes []byte) []byte {
	var script []byte
	script = append(script, P2SH_SCRIPT_PREFIX...)
	script = append(script, bytes...)
	script = append(script, OP_EQUAL)
	return script
}

func buildWitnessScript(bytes []byte) ([]byte, error) {
	if bytes[0] != BTC_WITNESS_VERSION {
		return nil, errors.New("wrong witness version")
	}
	converted, err := bech32.ConvertBits(bytes[1:], 5, 8, false)
	if err != nil {
		return nil, errors.Wrap(err, "converting bits failed")
	}
	script := []byte{BTC_WITNESS_VERSION, byte(len(converted))}
	script = append(script, converted...)
	return script, nil
}
