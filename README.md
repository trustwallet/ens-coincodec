# ens-coincodec

[![Tag](https://img.shields.io/github/tag/trustwallet/ens-coincodec.svg)](https://github.com/trustwallet/ens-coincodec/releases/)
[![License](https://img.shields.io/github/license/trustwallet/ens-coincodec.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/trustwallet/ens-coincodec?status.svg)](https://godoc.org/github.com/trustwallet/ens-coincodec)
![CI](https://github.com/trustwallet/ens-coincodec/workflows/CI/badge.svg)
[![codecov.io](https://img.shields.io/codecov/c/github/trustwallet/ens-coincodec.svg)](https://codecov.io/github/trustwallet/ens-coincodec)
[![Go Report](https://goreportcard.com/badge/github.com/trustwallet/ens-coincodec)](https://goreportcard.com/report/github.com/trustwallet/ens-coincodec)

Go utility library to provide movement between string and binary representation of multpile different cryptocurrency coin formats, mainly for ENS, please checkout [EIP-2304](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2304.md) for details.


## Table of Contents

- [Supported Coins](#coins)
- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Coins

<a href="https://bitcoin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/0.png" width="32" /></a>
<a href="https://litecoin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/2.png" width="32" /></a>
<a href="https://dogecoin.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/3.png" width="32" /></a>
<a href="https://dash.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/5.png" width="32" /></a>
<a href="https://viacoin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/14.png" width="32" /></a>
<a href="https://www.digibyte.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/20.png" width="32" /></a>
<a href="https://monacoin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/22.png" width="32" /></a>
<a href="https://ethereum.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/60.png" width="32" /></a>
<a href="https://ethereumclassic.github.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/61.png" width="32" /></a>
<a href="https://cosmos.network/" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/118.png" width="32" /></a>
<a href="https://z.cash" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/133.png" width="32" /></a>
<a href="https://zcoin.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/136.png" width="32" /></a>
<a href="https://ripple.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/144.png" width="32" /></a>
<a href="https://bitcoincash.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/145.png" width="32" /></a>
<a href="https://stellar.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/148.png" width="32" /></a>
<a href="https://ravencoin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/175.png" width="32" /></a>
<a href="https://poa.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/178.png" width="32" /></a>
<a href="https://tron.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/195.png" width="32" /></a>
<a href="https://nimiq.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/242.png" width="32" /></a>
<a href="https://iotex.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/304.png" width="32" /></a>
<a href="https://zilliqa.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/313.png" width="32" /></a>
<a href="https://www.thetatoken.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/500.png" width="32" /></a>
<a href="https://binance.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/714.png" width="32" /></a>
<a href="https://vechain.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/818.png" width="32" /></a>
<a href="https://callisto.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/820.png" width="32" /></a>
<a href="https://tomochain.network" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/889.png" width="32" /></a>
<a href="https://thudercore.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1001.png" width="32" /></a>
<a href="https://ont.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1024.png" width="32" /></a>
<a href="https://tezos.com" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/1729.png" width="32" /></a>
<a href="https://kin.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/2017.png" width="32" /></a>
<a href="https://qtum.org" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/2301.png" width="32" /></a>
<a href="https://gochain.io" target="_blank"><img src="https://raw.githubusercontent.com/TrustWallet/tokens/master/coins/6060.png" width="32" /></a>

## Install

`ens-coincodec` is a standard Go module which can be installed with:

```sh
go get github.com/trustwallet/ens-coincodec
```

## Usage

### Example

```go
import (
	"fmt"

	cc "github.com/trustwallet/ens-coincodec"
	slip44 "github.com/wealdtech/go-slip44"
)

func main() {
	// Ethereum
	bytes, err := cc.ToBytes("0x314159265dD8dbb310642f98f50C066173C1259b", slip44.ETHER)
	// hex: 314159265dd8dbb310642f98f50c066173c1259b
	if err != nil {
		panic(err)
	}
	str, err := cc.ToString(bytes, slip44.ETHER)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ethereum: %s\n", str)

	// Bitcoin
	bytes, err = cc.ToBytes("bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4", slip44.BITCOIN)
	// script hash: 0014751e76e8199196d454941c45d1b3a323f1433bd6
	if err != nil {
		panic(err)
	}
	str, err = cc.ToString(bytes, slip44.BITCOIN)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bitcoin: %s\n", str)

	// BNB
	bytes, err = cc.ToBytes("bnb1grpf0955h0ykzq3ar5nmum7y6gdfl6lxfn46h2", slip44.BINANCE)
	// public key hash: 40c2979694bbc961023d1d27be6fc4d21a9febe6
	if err != nil {
		panic(err)
	}
	str, err = cc.ToString(bytes, slip44.BINANCE)
	if err != nil {
		panic(err)
	}

	fmt.Printf("BNB: %s\n", str)
}
```

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/trustwallet/ens-coincodec/issues).

If you are adding a new coin type please try to follow the following rules:

  - use the existing `ethereum.go` and `ethereum_test.go` as templates
  - ensure you have 100% code coverage with your tests
  - try not to import large amounts of code; consider copying the relevant code rather than bringing in an entire project to use the address conversion functions

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd / Trust Wallet
