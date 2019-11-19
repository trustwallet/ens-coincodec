package coincodec

import "github.com/wealdtech/go-slip44"

var (
	LTC_CONFIG  = CoinConfig{P2PKHPrefix: 0x30, P2SHPPrefix: 0x32, HRP: "ltc"}
	MONA_CONFIG = CoinConfig{P2PKHPrefix: 0x32, P2SHPPrefix: 0x37, HRP: "mona"}
	QTUM_CONFIG = CoinConfig{P2PKHPrefix: 0x3A, P2SHPPrefix: 0x32, HRP: "qc"}
	VIA_CONFIG  = CoinConfig{P2PKHPrefix: 0x47, P2SHPPrefix: 0x21, HRP: "via"}
	DGB_CONFIG  = CoinConfig{P2PKHPrefix: 0x1E, P2SHPPrefix: 0x3F, HRP: "dgb"}
	DOGE_CONFIG = CoinConfig{P2PKHPrefix: 0x1E, P2SHPPrefix: 0x16}
	DASH_CONFIG = CoinConfig{P2PKHPrefix: 0x4C, P2SHPPrefix: 0x10}
	XZC_CONFIG  = CoinConfig{P2PKHPrefix: 0x52, P2SHPPrefix: 0x07}
	RVN_CONFIG  = CoinConfig{P2PKHPrefix: 0x3C, P2SHPPrefix: 0x7A}
)

func init() {
	// LTC
	toBytesMap[slip44.LITECOIN] = MakeBitcoinDecodeToBytes(&LTC_CONFIG)
	toStringMap[slip44.LITECOIN] = MakeBitcoinEncodeToString(&LTC_CONFIG)

	// DOGE
	toBytesMap[slip44.DOGECOIN] = MakeBitcoinDecodeToBytes(&DOGE_CONFIG)
	toStringMap[slip44.DOGECOIN] = MakeBitcoinEncodeToString(&DOGE_CONFIG)

	// DASH
	toBytesMap[slip44.DASH] = MakeBitcoinDecodeToBytes(&DASH_CONFIG)
	toStringMap[slip44.DASH] = MakeBitcoinEncodeToString(&DASH_CONFIG)

	// VIA
	toBytesMap[slip44.VIACOIN] = MakeBitcoinDecodeToBytes(&VIA_CONFIG)
	toStringMap[slip44.VIACOIN] = MakeBitcoinEncodeToString(&VIA_CONFIG)

	// DGB
	toBytesMap[slip44.DIGIBYTE] = MakeBitcoinDecodeToBytes(&DGB_CONFIG)
	toStringMap[slip44.DIGIBYTE] = MakeBitcoinEncodeToString(&DGB_CONFIG)

	// MONA
	toBytesMap[slip44.MONACOIN] = MakeBitcoinDecodeToBytes(&MONA_CONFIG)
	toStringMap[slip44.MONACOIN] = MakeBitcoinEncodeToString(&MONA_CONFIG)

	// XZC
	toBytesMap[slip44.ZCOIN] = MakeBitcoinDecodeToBytes(&XZC_CONFIG)
	toStringMap[slip44.ZCOIN] = MakeBitcoinEncodeToString(&XZC_CONFIG)

	// RVN
	toBytesMap[slip44.RAVENCOIN] = MakeBitcoinDecodeToBytes(&RVN_CONFIG)
	toStringMap[slip44.RAVENCOIN] = MakeBitcoinEncodeToString(&RVN_CONFIG)

	// QTUM
	toBytesMap[slip44.QTUM] = MakeBitcoinDecodeToBytes(&QTUM_CONFIG)
	toStringMap[slip44.QTUM] = MakeBitcoinEncodeToString(&QTUM_CONFIG)
}
