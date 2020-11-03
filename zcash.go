package coincodec

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

const (
	ZCashAddressLength    = 22
	ZCashStaticPrefixByte = 0x1C // 28
	ZCashPrefixByteP2pkh  = 0xB8 // 184
	ZCashPrefixByteP2sh   = 0xBD // 189
)

var (
	zcashPrefixP2pkh = []byte{ZCashStaticPrefixByte, ZCashPrefixByteP2pkh}
	zcashPrefixP2sh  = []byte{ZCashStaticPrefixByte, ZCashPrefixByteP2sh}
	zcashPrefixes    = [][]byte{zcashPrefixP2pkh, zcashPrefixP2sh}
)

func init() {
	toBytesMap[slip44.ZCASH] = ZCashDecodeToBytes
	toStringMap[slip44.ZCASH] = ZCashEncodeToString
}

// ZCashDecodeToBytes converts the input string to a byte array
func ZCashDecodeToBytes(input string) ([]byte, error) {
	data, err := Base58AddressDecodeToBytesPrefix(input, ZCashAddressLength, zcashPrefixes)
	if err != nil {
		return nil, err
	}
	if bytes.HasPrefix(data, zcashPrefixP2pkh) {
		return buildP2PKHScript(data[len(zcashPrefixP2pkh):]), nil
	}
	if bytes.HasPrefix(data, zcashPrefixP2sh) {
		return buildP2SHScript(data[len(zcashPrefixP2sh):]), nil
	}
	return nil, errors.New("Invalid prefix")
}

// ZCashEncodeToString converts the input byte array to a string representation of the ZCash address.
func ZCashEncodeToString(input []byte) (string, error) {
	if bytes.HasPrefix(input, P2PKH_SCRIPT_PREFIX) {
		return Base58ChecksumEncode(replacePrefix(input, P2PKH_SCRIPT_PREFIX, zcashPrefixP2pkh), Base58DefaultAlphabet), nil
	}
	if bytes.HasPrefix(input, P2SH_SCRIPT_PREFIX) {
		return Base58ChecksumEncode(replacePrefix(input, P2SH_SCRIPT_PREFIX, zcashPrefixP2sh), Base58DefaultAlphabet), nil
	}
	return "", errors.New("Invalid opcode bytes")
}
