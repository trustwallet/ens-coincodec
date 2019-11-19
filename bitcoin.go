package coincodec

import (
	"bytes"

	"github.com/btcsuite/btcutil/base58"
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

func MakeBitcoinDecodeToBytes(config *CoinConfig) func(string) ([]byte, error) {
	return func(input string) ([]byte, error) {
		return bitcoinDecodeToBytes(input, config)
	}
}

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
	bytes, version, err := base58.CheckDecode(input)
	if err != nil {
		if len(config.HRP) <= 0 {
			return nil, err
		}
		// try bech32
		decodedHrp, bytes, err := bech32.Decode(input)
		if err != nil {
			return nil, errors.Wrapf(err, "decoding bech32 failed")
		}
		if decodedHrp != config.HRP {
			return nil, errors.New("invalid hrp")
		}
		return buildWitnessScript(bytes)
	}
	// check data length
	if len(bytes) != btcKeyHashLenght {
		return nil, errors.New("invalid data length")
	}
	// check version byte
	if version == config.P2PKHPrefix {
		return buildP2PKHScript(bytes), nil
	} else if version == config.P2SHPPrefix {
		return buildP2SHScript(bytes), nil
	}
	return nil, errors.New("invalid address prefix")
}

func bitcoinEncodeToString(input []byte, config *CoinConfig) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid data length")
	}
	if bytes.HasPrefix(input, P2PKH_SCRIPT_PREFIX) {
		return base58.CheckEncode(input[3:len(input)-2], config.P2PKHPrefix), nil
	} else if bytes.HasPrefix(input, P2SH_SCRIPT_PREFIX) {
		return base58.CheckEncode(input[2:len(input)-1], config.P2SHPPrefix), nil
	} else if input[0] == btcWitnessVersion && len(input) > 2 && len(config.HRP) > 0 {
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
