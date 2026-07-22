---
title: Build and Dependencies
category: build
tags: [go-modules, dependencies, go1.22, btcsuite]
confidence: high
source: go.mod, .github/workflows/ci.yml
updated: 2026-07-22
---

# Build and Dependencies

## Module

`github.com/trustwallet/ens-coincodec`, Go 1.22.

## Key dependencies

| Module | Purpose |
|---|---|
| `github.com/trustwallet/go-primitives` | Coin type constants (`coin.BITCOIN`, `coin.ETHEREUM`, `IsEVM`, etc.) |
| `github.com/btcsuite/btcutil` | Bech32 encode/decode (used in bitcoin.go and bech32.go) |
| `github.com/cpacia/bchutil` | Bitcoin Cash cashaddr encoding |
| `github.com/pkg/errors` | Error wrapping with context (used in bitcoin.go) |
| `golang.org/x/crypto` | Keccak-256 for Ethereum EIP-55 checksum |

## CI

```
go test ./...
```

Single command, no build step — this is a library, not a binary. The CI workflow (`ci.yml`) runs tests and Codecov coverage upload on every push and PR.

## Install

```sh
go get github.com/trustwallet/ens-coincodec
```

## No binary output

This repo ships a Go package only. `go build ./...` succeeds (no syntax errors) but produces no binary artifacts.

## See Also
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:strong -->
- [overview](../architecture/overview.md) <!-- rel:strong -->
- [strategy](../tests/strategy.md) <!-- rel:strong -->
- [readme](../ci/workflows/readme.md) <!-- rel:strong -->
- [encoding primitives](../architecture/encoding-primitives.md) <!-- rel:strong -->
