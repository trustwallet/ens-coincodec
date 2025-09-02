package coincodec

import (
	"github.com/trustwallet/go-primitives/coin"
)

const (
	ZCashStaticPrefixByte = 0x1C // 28
	ZCashPrefixByteP2pkh  = 0xB8 // 184
	ZCashPrefixByteP2sh   = 0xBD // 189
)

var (
	zcashPrefixP2pkh = []byte{ZCashStaticPrefixByte, ZCashPrefixByteP2pkh}
	zcashPrefixP2sh  = []byte{ZCashStaticPrefixByte, ZCashPrefixByteP2sh}
)

var (
	configLTC   = CoinConfig{P2PKHPrefix: []byte{0x30}, P2SHPPrefix: []byte{0x32}, HRP: "ltc"}
	configMONA  = CoinConfig{P2PKHPrefix: []byte{0x32}, P2SHPPrefix: []byte{0x37}, HRP: "mona"}
	configQTUM  = CoinConfig{P2PKHPrefix: []byte{0x3A}, P2SHPPrefix: []byte{0x32}, HRP: "qc"}
	configVIA   = CoinConfig{P2PKHPrefix: []byte{0x47}, P2SHPPrefix: []byte{0x21}, HRP: "via"}
	configDGB   = CoinConfig{P2PKHPrefix: []byte{0x1E}, P2SHPPrefix: []byte{0x3F}, HRP: "dgb"}
	configDOGE  = CoinConfig{P2PKHPrefix: []byte{0x1E}, P2SHPPrefix: []byte{0x16}}
	configDASH  = CoinConfig{P2PKHPrefix: []byte{0x4C}, P2SHPPrefix: []byte{0x10}}
	configXZC   = CoinConfig{P2PKHPrefix: []byte{0x52}, P2SHPPrefix: []byte{0x07}}
	configRVN   = CoinConfig{P2PKHPrefix: []byte{0x3C}, P2SHPPrefix: []byte{0x7A}}
	configZCASH = CoinConfig{P2PKHPrefix: zcashPrefixP2pkh, P2SHPPrefix: zcashPrefixP2sh}
)

func init() {
	// LTC
	toBytesMap[coin.LITECOIN] = MakeBitcoinDecodeToBytes(&configLTC)
	toStringMap[coin.LITECOIN] = MakeBitcoinEncodeToString(&configLTC)

	// DOGE
	toBytesMap[coin.DOGE] = MakeBitcoinDecodeToBytes(&configDOGE)
	toStringMap[coin.DOGE] = MakeBitcoinEncodeToString(&configDOGE)

	// DASH
	toBytesMap[coin.DASH] = MakeBitcoinDecodeToBytes(&configDASH)
	toStringMap[coin.DASH] = MakeBitcoinEncodeToString(&configDASH)

	// VIA
	toBytesMap[coin.VIACOIN] = MakeBitcoinDecodeToBytes(&configVIA)
	toStringMap[coin.VIACOIN] = MakeBitcoinEncodeToString(&configVIA)

	// DGB
	toBytesMap[coin.DIGIBYTE] = MakeBitcoinDecodeToBytes(&configDGB)
	toStringMap[coin.DIGIBYTE] = MakeBitcoinEncodeToString(&configDGB)

	// MONA
	toBytesMap[coin.MONACOIN] = MakeBitcoinDecodeToBytes(&configMONA)
	toStringMap[coin.MONACOIN] = MakeBitcoinEncodeToString(&configMONA)

	// XZC
	toBytesMap[coin.FIRO] = MakeBitcoinDecodeToBytes(&configXZC)
	toStringMap[coin.FIRO] = MakeBitcoinEncodeToString(&configXZC)

	// RVN
	toBytesMap[coin.RAVENCOIN] = MakeBitcoinDecodeToBytes(&configRVN)
	toStringMap[coin.RAVENCOIN] = MakeBitcoinEncodeToString(&configRVN)

	// QTUM
	toBytesMap[coin.QTUM] = MakeBitcoinDecodeToBytes(&configQTUM)
	toStringMap[coin.QTUM] = MakeBitcoinEncodeToString(&configQTUM)

	// ZCASH
	toBytesMap[coin.ZCASH] = MakeBitcoinDecodeToBytes(&configZCASH)
	toStringMap[coin.ZCASH] = MakeBitcoinEncodeToString(&configZCASH)
}
