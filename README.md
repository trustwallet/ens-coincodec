# ens-coincodec

[![Tag](https://img.shields.io/github/tag/trustwallet/ens-coincodec.svg)](https://github.com/trustwallet/ens-coincodec/releases/)
[![License](https://img.shields.io/github/license/trustwallet/ens-coincodec.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/trustwallet/ens-coincodec?status.svg)](https://godoc.org/github.com/trustwallet/ens-coincodec)
[![Build Status](https://travis-ci.org/trustwallet/ens-coincodec.svg?branch=master)](https://travis-ci.org/trustwallet/ens-coincodec)
[![codecov.io](https://img.shields.io/codecov/c/github/trustwallet/ens-coincodec.svg)](https://codecov.io/github/trustwallet/ens-coincodec)
[![Go Report](https://goreportcard.com/badge/github.com/trustwallet/ens-coincodec)](https://goreportcard.com/report/github.com/trustwallet/ens-coincodec)

Go utility library to provide movement between string and binary representation of multpile different cryptocurrency coin formats, mainly for ENS, please checkout [EIP-2304](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2304.md) for details.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

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
