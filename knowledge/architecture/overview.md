---
title: Architecture Overview
category: architecture
tags: [ens, coincodec, registry, eip-2304, go-library]
confidence: high
source: codec.go, README.md, ethereum.go, bitcoin.go
updated: 2026-07-22
---

# Architecture Overview

`ens-coincodec` is a Go library that implements **EIP-2304** — bidirectional conversion between the human-readable string form of a cryptocurrency address and its canonical binary form as stored in ENS resolver records.

## Responsibility

ENS stores multi-coin addresses as raw bytes in resolver contracts. When a caller resolves an address for a given coin type, they get bytes back and need to convert them to the displayable string form for that coin. Conversely, when setting an address, the string must be converted to bytes. This library owns that conversion for ~35 coin types.

## Public API (codec.go)

Two functions form the entire public interface:

```go
// ToBytes converts a coin's string address to its binary representation.
func ToBytes(input string, coinType uint32) ([]byte, error)

// ToString converts a coin's binary representation to its string address.
func ToString(input []byte, coinType uint32) (string, error)
```

`coinType` is the SLIP-44 numeric identifier (same standard used by BIP-44 derivation paths). The library uses `github.com/trustwallet/go-primitives/coin` for the canonical coin type constants.

## Registry Pattern

The library uses two package-level maps as its dispatch table:

```go
var toBytesMap = make(map[uint32]func(string) ([]byte, error))
var toStringMap = make(map[uint32]func([]byte) (string, error))
```

Every coin implementation registers its codec functions into these maps via a Go `init()` function that runs automatically when the package is imported. There is no explicit startup call required — importing `coincodec` is sufficient.

## Layered Architecture

```
Public API (codec.go)
    └── Per-coin init() registration
            ├── EVM coins (ethereum.go) — single codec for all ~49 EVM chains
            ├── Bitcoin (bitcoin.go) + Forks (bitcoin_forks.go) — factory pattern
            ├── Bech32 coins (bnb, cosmos, elrond, iotex, zilliqa)
            ├── Base58+prefix coins (tron, tezos, ontology)
            ├── Specialized codecs (filecoin, nimiq, stellar, solana, polkadot, xrp)
            └── Encoding primitives
                    ├── base58.go — raw Base58
                    ├── base58_address.go — Base58 + prefix validation
                    ├── base58_checksum.go — Base58 + double-SHA256 checksum
                    ├── bech32.go — raw Bech32
                    └── bech32_address.go — Bech32 + HRP + keyhash length validation
```

## Domains

| File group | Encoding scheme | Coins |
|---|---|---|
| `ethereum.go` | Hex + EIP-55 checksum | All EVM coins (ETH, BNB-ERC20, etc.) |
| `bitcoin.go` + `bitcoin_forks.go` | Base58Check / Bech32 (P2PKH/P2SH/witness) | BTC, LTC, DOGE, DASH, ZCash, RVN, DGB, VIA, MONA, XZC, QTUM |
| `bitcoin_cash.go` | Cashaddr (via btcutil) | BCH |
| `bech32_address.go` users | Bech32 + HRP | BNB, COSMOS, ELROND, IOTEX, ZILLIQA |
| `base58_address.go` users | Base58Check + prefix | TRON, TEZOS, ONTOLOGY |
| `ss58address.go` | SS58 (Blake2b checksum) | POLKADOT |
| `solana.go` | Raw Base58 | SOL |
| `stellar.go` | Base32 + CRC-16 | XLM |
| `xrp.go` | Base58Check (custom alphabet) | XRP |
| `filecoin.go` | Custom (type byte + payload + Base32 checksum) | FIL |
| `nimiq.go` | Custom (uppercase hex + spaces + modulo-97 checksum) | NIM |

## What the binary representation IS

For Bitcoin-family coins, the binary is the **scriptPubKey** (locking script), NOT the raw address bytes. For example, a P2PKH address decodes to `OP_DUP OP_HASH160 <20-byte hash> OP_EQUALVERIFY OP_CHECKSIG`. This matches ENS EIP-2304's intent: store the on-chain script, not just the hash.

For Ethereum/EVM coins the binary is the raw 20-byte address (no checksum, no `0x` prefix).

## See Also
- [init registration](../patterns/init-registration.md) <!-- rel:strong -->
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:strong -->
- [adding a new coin](../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [readme](../ci/workflows/readme.md) <!-- rel:strong -->
- [overview](../build/overview.md) <!-- rel:strong -->
