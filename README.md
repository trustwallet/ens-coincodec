# ens-coincodec

[![Tag](https://img.shields.io/github/tag/trustwallet/ens-coincodec.svg)](https://github.com/trustwallet/ens-coincodec/releases/)
[![License](https://img.shields.io/github/license/trustwallet/ens-coincodec.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/trustwallet/ens-coincodec?status.svg)](https://godoc.org/github.com/trustwallet/ens-coincodec)
[![Build Status](https://travis-ci.org/trustwallet/ens-coincodec.svg?branch=master)](https://travis-ci.org/trustwallet/ens-coincodec)
[![codecov.io](https://img.shields.io/codecov/c/github/trustwallet/ens-coincodec.svg)](https://codecov.io/github/trustwallet/ens-coincodec)
[![Go Report](https://goreportcard.com/badge/github.com/trustwallet/ens-coincodec)](https://goreportcard.com/report/github.com/trustwallet/ens-coincodec)

Go utility library to provide movement between string and binary representation of multpile different cryptocurrency coin formats, mainly for ENS.


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
    "bytes"
    "encoding/hex"
    "errors"

    slip44 "github.com/wealdtech/go-slip44"
    cc "github.com/trustwallet/ens-coincodec"
)

func main() {
    bytes, err := cc.ToBytes("0x0102030405060708090A0b0c0d0e0f1011121314", slip44.ETHER)
    if err != nil {
        panic(err)
    }

    str, err := cc.ToString(bytes, slip44.ETHER)
    if err != nil {
        panic(err)
    }

    fmt.Printf("%s\n", str)
}
```

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/trustwallet/ens-coincodec/issues).

If you are adding a new coin type please try to follow the following rules:

  - use the existing `ether.go` and `ether_test.go` as templates
  - ensure you have 100% code coverage with your tests
  - try not to import large amounts of code; consider copying the relevant code rather than bringing in an entire project to use the address conversion functions

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd / Trust Wallet
