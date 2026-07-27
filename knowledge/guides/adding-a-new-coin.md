---
title: Adding a New Coin
category: guides
tags: [contributing, new-coin, add-support, codec]
confidence: high
source: README.md, ethereum.go, bitcoin.go, test_runner.go
updated: 2026-07-22
---

# Adding a New Coin

## Step-by-step

1. **Create `<coin>.go`** — use `ethereum.go` as the template for simple coins. The file needs:
   - `func init()` — registers `DecodeToBytes`/`EncodeToString` into the global maps
   - `func <Coin>DecodeToBytes(input string) ([]byte, error)`
   - `func <Coin>EncodeToString(input []byte) (string, error)`

2. **Pick the right encoding primitive** based on the coin's address format:
   - Ethereum/EVM → `EtherToBytes`/`EtherToString` (EIP-55 hex)
   - Bitcoin-fork → use `MakeBitcoinDecodeToBytes`/`MakeBitcoinEncodeToString` with a `CoinConfig`
   - Bech32 chain (Cosmos ecosystem) → `Bech32AddressDecodeToBytes`/`Bech32AddressEncodeToString`
   - Base58Check chain (Tron-like) → `Base58AddressDecodeToBytesPrefix`/`Base58AddressEncodeToStringPrefix`

3. **Avoid large imports** — the README explicitly says to copy relevant code rather than importing a large project just for address conversion.

4. **Create `<coin>_test.go`** with 100% coverage using `RunTestsDecode`/`RunTestsEncode` from `test_runner.go`.

## EVM chains are automatic

New EVM chains added to `github.com/trustwallet/go-primitives/coin.Coins` are **automatically supported** because `ethereum.go`'s `init()` iterates over all `IsEVM` coins. No code change in this repo is required for a new EVM chain.

## Bitcoin forks

Add a `CoinConfig` to `bitcoin_forks.go` and register it in the existing `init()`. The P2PKH prefix, P2SH prefix, and HRP (for SegWit) are the only chain-specific values.

## See Also
- [init registration](../patterns/init-registration.md) <!-- rel:strong -->
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:strong -->
- [overview](../architecture/overview.md) <!-- rel:strong -->
- [encoding primitives](../architecture/encoding-primitives.md) <!-- rel:strong -->
- [call graph](../architecture/call-graph.md) <!-- rel:strong -->
