package coincodec

import (
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
	return Base58AddressDecodeToBytes(input, ZCashAddressLength, zcashPrefixes)
}

// ZCashEncodeToString converts the input byte array to a string representation of the ZCash address.
func ZCashEncodeToString(bytes []byte) (string, error) {
	return Base58AddressEncodeToString(bytes, ZCashAddressLength, zcashPrefixes)
}
