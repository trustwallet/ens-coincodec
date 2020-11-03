package coincodec

import "github.com/wealdtech/go-slip44"

var (
	configLTC  = CoinConfig{P2PKHPrefix: []byte{0x30}, P2SHPPrefix: []byte{0x32}, HRP: "ltc"}
	configMONA = CoinConfig{P2PKHPrefix: []byte{0x32}, P2SHPPrefix: []byte{0x37}, HRP: "mona"}
	configQTUM = CoinConfig{P2PKHPrefix: []byte{0x3A}, P2SHPPrefix: []byte{0x32}, HRP: "qc"}
	configVIA  = CoinConfig{P2PKHPrefix: []byte{0x47}, P2SHPPrefix: []byte{0x21}, HRP: "via"}
	configDGB  = CoinConfig{P2PKHPrefix: []byte{0x1E}, P2SHPPrefix: []byte{0x3F}, HRP: "dgb"}
	configDOGE = CoinConfig{P2PKHPrefix: []byte{0x1E}, P2SHPPrefix: []byte{0x16}}
	configDASH = CoinConfig{P2PKHPrefix: []byte{0x4C}, P2SHPPrefix: []byte{0x10}}
	configXZC  = CoinConfig{P2PKHPrefix: []byte{0x52}, P2SHPPrefix: []byte{0x07}}
	configRVN  = CoinConfig{P2PKHPrefix: []byte{0x3C}, P2SHPPrefix: []byte{0x7A}}
)

func init() {
	// LTC
	toBytesMap[slip44.LITECOIN] = MakeBitcoinDecodeToBytes(&configLTC)
	toStringMap[slip44.LITECOIN] = MakeBitcoinEncodeToString(&configLTC)

	// DOGE
	toBytesMap[slip44.DOGECOIN] = MakeBitcoinDecodeToBytes(&configDOGE)
	toStringMap[slip44.DOGECOIN] = MakeBitcoinEncodeToString(&configDOGE)

	// DASH
	toBytesMap[slip44.DASH] = MakeBitcoinDecodeToBytes(&configDASH)
	toStringMap[slip44.DASH] = MakeBitcoinEncodeToString(&configDASH)

	// VIA
	toBytesMap[slip44.VIACOIN] = MakeBitcoinDecodeToBytes(&configVIA)
	toStringMap[slip44.VIACOIN] = MakeBitcoinEncodeToString(&configVIA)

	// DGB
	toBytesMap[slip44.DIGIBYTE] = MakeBitcoinDecodeToBytes(&configDGB)
	toStringMap[slip44.DIGIBYTE] = MakeBitcoinEncodeToString(&configDGB)

	// MONA
	toBytesMap[slip44.MONACOIN] = MakeBitcoinDecodeToBytes(&configMONA)
	toStringMap[slip44.MONACOIN] = MakeBitcoinEncodeToString(&configMONA)

	// XZC
	toBytesMap[slip44.ZCOIN] = MakeBitcoinDecodeToBytes(&configXZC)
	toStringMap[slip44.ZCOIN] = MakeBitcoinEncodeToString(&configXZC)

	// RVN
	toBytesMap[slip44.RAVENCOIN] = MakeBitcoinDecodeToBytes(&configRVN)
	toStringMap[slip44.RAVENCOIN] = MakeBitcoinEncodeToString(&configRVN)

	// QTUM
	toBytesMap[slip44.QTUM] = MakeBitcoinDecodeToBytes(&configQTUM)
	toStringMap[slip44.QTUM] = MakeBitcoinEncodeToString(&configQTUM)
}
