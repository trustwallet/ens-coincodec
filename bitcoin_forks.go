package coincodec

import "github.com/wealdtech/go-slip44"

var (
	LTC_CONFIG  = CoinConfig{P2PKHPrefix: 0x30, P2SHPPrefix: 0x32, HRP: "ltc"}
	DOGE_CONFIG = CoinConfig{P2PKHPrefix: 0x1E, P2SHPPrefix: 0x16}
	DASH_CONFIG = CoinConfig{P2PKHPrefix: 0x4C, P2SHPPrefix: 0x10}
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
}
