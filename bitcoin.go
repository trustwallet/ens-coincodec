package coincodec

import (
	"bytes"

	"github.com/pkg/errors"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/wealdtech/go-slip44"
)

// CoinConfig for Bitcoin and its forks
type CoinConfig struct {
	P2PKHPrefix byte
	P2SHPPrefix byte
	HRP         string
}

const (
	btcKeyHashLenght  = 0x14
	btcWitnessVersion = 0x00
)

const (
	OP_DUP         = 0x76
	OP_HASH160     = 0xa9
	OP_EQUALVERIFY = 0x88
	OP_CHECKSIG    = 0xac
	OP_EQUAL       = 0x87
)

var (
	configBTC           = CoinConfig{P2PKHPrefix: 0x00, P2SHPPrefix: 0x05, HRP: "bc"}
	P2PKH_SCRIPT_PREFIX = []byte{OP_DUP, OP_HASH160, btcKeyHashLenght}
	P2SH_SCRIPT_PREFIX  = []byte{OP_HASH160, btcKeyHashLenght}
)

func init() {
	toBytesMap[slip44.BITCOIN] = BitcoinDecodeToBytes
	toStringMap[slip44.BITCOIN] = BitcoinEncodeToString
}

// BitcoinDecodeToBytes converts the input string to a byte array
func BitcoinDecodeToBytes(input string) ([]byte, error) {
	return bitcoinDecodeToBytes(input, &configBTC)
}

// BitcoinEncodeToString converts the input byte array to a string representation of the Bitcoin address.
func BitcoinEncodeToString(input []byte) (string, error) {
	return bitcoinEncodeToString(input, &configBTC)
}

// MakeBitcoinDecodeToBytes takes a CoinConfig and returns a func to decode string to bytes
func MakeBitcoinDecodeToBytes(config *CoinConfig) func(string) ([]byte, error) {
	return func(input string) ([]byte, error) {
		return bitcoinDecodeToBytes(input, config)
	}
}

// MakeBitcoinEncodeToString takes a CoinConfig and returns a func to encode bytes to string
func MakeBitcoinEncodeToString(config *CoinConfig) func([]byte) (string, error) {
	return func(input []byte) (string, error) {
		return bitcoinEncodeToString(input, config)
	}
}

func bitcoinDecodeToBytes(input string, config *CoinConfig) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("invalid address")
	}
	// try base58 first
	bytes, err := Base58AddressDecodeToBytesPrefix(input, 20+1, [][]byte{[]byte{config.P2PKHPrefix}, []byte{config.P2SHPPrefix}})
	if err != nil {
		if len(config.HRP) <= 0 {
			return nil, err
		}
		// try bech32
		decodedHrp, bytes, err2 := bech32.Decode(input)
		if err2 != nil {
			return nil, errors.Wrapf(err2, "decoding base58 and bech32 failed: %v", err.Error())
		}
		if decodedHrp != config.HRP {
			return nil, errors.New("invalid hrp")
		}
		return buildWitnessScript(bytes)
	}
	var prefix byte
	if len(bytes) >= 1 {
		prefix = bytes[0]
		bytes = bytes[1:]
	}
	// check data length
	if len(bytes) != btcKeyHashLenght {
		return nil, errors.New("invalid data length")
	}
	// check prefix byte
	if prefix == config.P2PKHPrefix {
		return buildP2PKHScript(bytes), nil
	}
	if prefix == config.P2SHPPrefix {
		return buildP2SHScript(bytes), nil
	}
	return nil, errors.New("Invalid address prefix")
}

func replacePrefix(input []byte, oldPrefix []byte, newPrefix []byte) []byte {
	if !bytes.HasPrefix(input, oldPrefix) {
		return input
	}
	var withVersion []byte
	withVersion = append(withVersion, newPrefix...)
	withVersion = append(withVersion, input[len(oldPrefix):len(input)-len(oldPrefix)+1]...)
	return withVersion
}

func bitcoinEncodeToString(input []byte, config *CoinConfig) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid data length")
	}
	if bytes.HasPrefix(input, P2PKH_SCRIPT_PREFIX) {
		return Base58ChecksumEncode(replacePrefix(input, P2PKH_SCRIPT_PREFIX, []byte{config.P2PKHPrefix}), Base58DefaultAlphabet), nil
	}
	if bytes.HasPrefix(input, P2SH_SCRIPT_PREFIX) {
		return Base58ChecksumEncode(replacePrefix(input, P2SH_SCRIPT_PREFIX, []byte{config.P2SHPPrefix}), Base58DefaultAlphabet), nil
	}
	if input[0] == btcWitnessVersion && len(input) > 2 && len(config.HRP) > 0 {
		if int(input[1]) != len(input)-2 {
			return "", errors.New("wrong script data")
		}
		converted, err := bech32.ConvertBits(input[2:], 8, 5, true)
		if err != nil {
			return "", errors.Wrap(err, "converting bits failed")
		}
		data := []byte{btcWitnessVersion}
		data = append(data, converted...)
		return bech32.Encode(config.HRP, data)
	}
	return "", errors.New("Invalid bytes")
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
	if bytes[0] != btcWitnessVersion {
		return nil, errors.New("wrong witness version")
	}
	converted, err := bech32.ConvertBits(bytes[1:], 5, 8, false)
	if err != nil {
		return nil, errors.Wrap(err, "converting bits failed")
	}
	script := []byte{btcWitnessVersion, byte(len(converted))}
	script = append(script, converted...)
	return script, nil
}
