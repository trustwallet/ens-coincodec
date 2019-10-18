# go-coincodec

[![Tag](https://img.shields.io/github/tag/wealdtech/go-coincodec.svg)](https://github.com/wealdtech/go-coincodec/releases/)
[![License](https://img.shields.io/github/license/wealdtech/go-coincodec.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/go-coincodec?status.svg)](https://godoc.org/github.com/wealdtech/go-coincodec)
[![Travis CI](https://img.shields.io/travis/wealdtech/go-coincodec.svg)](https://travis-ci.org/wealdtech/go-coincodec)
[![codecov.io](https://img.shields.io/codecov/c/github/wealdtech/go-coincodec.svg)](https://codecov.io/github/wealdtech/go-coincodec)

Go utility library to provide movement between string and binary representation of multpile different cryptocurrency coin formats.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-coincodec` is a standard Go module which can be installed with:

```sh
go get github.com/wealdtech/go-coincodec
```

## Usage

### Example

```go
import (
    "bytes"
    "encoding/hex"
    "errors"

    slip44 "github.com/wealdtech/go-slip44"
    cc "github.com/wealdtech/go-coincodec"
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

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/go-coincodec/issues).

If you are adding a new coin type please try to follow the following rules:

  - use the existing `ether.go` and `ether_test.go` as templates
  - ensure you have 100% code coverage with your tests
  - try not to import large amounts of code; consider copying the relevant code rather than bringing in an entire project to use the address conversion functions

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
