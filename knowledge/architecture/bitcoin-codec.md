---
title: Bitcoin Multi-Format Codec
category: architecture
tags: [bitcoin, utxo, p2pkh, p2sh, bech32, witness, script, factory-pattern]
confidence: high
source: bitcoin.go, bitcoin_forks.go
updated: 2026-07-22
---

# Bitcoin Multi-Format Codec

`bitcoin.go` handles the most complex codec in the library because Bitcoin addresses encode a **scriptPubKey**, not just an address hash.

## Three address types, one codec

`bitcoinDecodeToBytes` handles all three P2PKH/P2SH/witness types automatically by trying formats in order:

1. **Base58Check** — tries `Base58AddressDecodeToBytesPrefix` first with P2PKH and P2SH prefixes
2. **Bech32 (SegWit)** — if Base58 fails and the coin config has an HRP, tries bech32 decode

On success, the binary output is the **scriptPubKey** (locking script), not just the hash:
- P2PKH → `OP_DUP OP_HASH160 <20-byte hash> OP_EQUALVERIFY OP_CHECKSIG`
- P2SH → `OP_HASH160 <20-byte hash> OP_EQUAL`
- Witness → `<version byte> <length byte> <data bytes>`

`bitcoinEncodeToString` does the reverse: detects the script prefix to determine format, then encodes accordingly.

## CoinConfig and the fork factory

`CoinConfig` holds three fields: `P2PKHPrefix`, `P2SHPPrefix`, and `HRP`. Different Bitcoin forks have different version bytes and different Bech32 prefixes. The factory functions `MakeBitcoinDecodeToBytes` and `MakeBitcoinEncodeToString` close over a `CoinConfig` and return a codec function — this is the entire mechanism for Bitcoin forks.

`bitcoin_forks.go` defines a `CoinConfig` for each fork and registers it in `init()`:

| Coin | P2PKH prefix | P2SH prefix | HRP |
|---|---|---|---|
| LTC | `0x30` | `0x32` | `ltc` |
| DOGE | `0x1E` | `0x16` | (none) |
| DASH | `0x4C` | `0x10` | (none) |
| ZCash | `0x1C 0xB8` | `0x1C 0xBD` | (none) |

ZCash uses a 2-byte prefix (`ZCashStaticPrefixByte` + type byte), which is why `validPrefixes` accepts `[][]byte` — this is the only case with a multi-byte version prefix.

## Bitcoin Cash is separate

`bitcoin_cash.go` re-uses `BitcoinDecodeToBytes` (via `bchutil`) but wraps it in its own cashaddr codec, which is registered directly. It is NOT in the `bitcoin_forks.go` factory pattern.

## Ambiguous naming: exported vs unexported

The call graph shows `BitcoinDecodeToBytes` (exported, thin wrapper) and `bitcoinDecodeToBytes` (unexported, implements logic). This split exists so:
- `BitcoinDecodeToBytes` can be the public API (used by `init()` for BTC)
- `bitcoinDecodeToBytes` can be reused by `MakeBitcoinDecodeToBytes` for forks with a config parameter

The static analysis flags these as "possible duplicates" because their names differ only by capitalization.

## See Also
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:strong -->
- [adding a new coin](../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [init registration](../patterns/init-registration.md) <!-- rel:strong -->
- [overview](../build/overview.md) <!-- rel:related -->
- [strategy](../tests/strategy.md) <!-- rel:weak -->
