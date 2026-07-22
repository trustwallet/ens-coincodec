---
title: Code Conventions
category: code-conventions
tags: [go, naming, error-handling, style]
confidence: high
source: codec.go, ethereum.go, bitcoin.go, base58_address.go
updated: 2026-07-22
---

# Code Conventions

## Function naming

- Exported decode functions: `<Coin>DecodeToBytes(input string) ([]byte, error)`
- Exported encode functions: `<Coin>EncodeToString(input []byte) (string, error)`
- Unexported implementation: `<coin>DecodeToBytes(input, config)` — lowercase first letter, config parameter for reuse by factory/forks
- Factory constructors: `Make<Coin>DecodeToBytes(config)` returns a closure

## Error handling

Errors are returned as values — no panics. The library uses `github.com/pkg/errors` for wrapping (bitcoin.go) and the standard `errors` package for simple errors (ethereum.go, base58).

Wrapped errors include context: `errors.Wrapf(err, "decoding base58 and bech32 failed")` so callers can trace which encoding path failed.

## Constants and variables

Address prefixes are declared as package-level `var []byte` (or `[][]byte` for multi-prefix coins). Script opcode constants use `const`. Chain-specific configs use exported names when they may be useful to consumers (`P2PKH_SCRIPT_PREFIX`, `ZCashStaticPrefixByte`).

## No external address validation libraries

The README explicitly prefers copying small, relevant functions over importing large projects. The library implements its own Base58, Base58Check, Bech32 wrappers, SS58, Stellar CRC-16, and Nimiq modulo-97 rather than pulling in full chain SDKs.

## Test files mirror source files

Every `<coin>.go` has a corresponding `<coin>_test.go`. Tests use the shared `RunTestsDecode`/`RunTestsEncode` helpers for consistency.

## See Also
- [adding a new coin](../../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [bitcoin codec](../../architecture/bitcoin-codec.md) <!-- rel:strong -->
- [overview](../../architecture/overview.md) <!-- rel:strong -->
- [init registration](../../patterns/init-registration.md) <!-- rel:strong -->
- [encoding primitives](../../architecture/encoding-primitives.md) <!-- rel:strong -->
