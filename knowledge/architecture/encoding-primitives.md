---
title: "God-Node Deep-Dive: Base58 and Bech32 Encoding Primitives"
category: architecture
tags: [base58, bech32, address-encoding, god-nodes, primitives]
confidence: high
source: base58.go, base58_address.go, base58_checksum.go, bech32.go, bech32_address.go
updated: 2026-07-22
---

# God-Node Deep-Dive: Base58 and Bech32 Encoding Primitives

These five files are the most-depended-on symbols in the library. Every coin-specific codec either uses them directly or delegates through a thin wrapper.

## Base58 stack (3 levels)

### `base58.go` — raw encode/decode

`Base58Encode(b []byte, alphabet string) string` and `Base58Decode(b string, alphabet string) ([]byte, error)` implement the raw Base58 alphabet conversion with no checksum or versioning. The alphabet is a parameter — callers pass `Base58DefaultAlphabet` (the standard Bitcoin alphabet) or a custom one (XRP uses its own).

**Called by:** `Base58ChecksumDecode`, `SS58AddressDecodeToBytes`, `SolanaDecodeToBytes`.

### `base58_checksum.go` — checksum layer

`Base58ChecksumEncode` and `Base58ChecksumDecode` add a 4-byte double-SHA256 checksum (appended before encoding / verified on decode). This is the standard "Base58Check" format used by Bitcoin and most UTXO chains.

**Called by:** `Base58AddressDecodeToBytesPrefix`, `XRPDecodeToBytes`/`XRPEncodeToString` (with custom alphabet), `bitcoinEncodeToString`.

### `base58_address.go` — prefix validation layer

`Base58AddressDecodeToBytesPrefix(input, decodedSize, validPrefixes)` and `Base58AddressEncodeToStringPrefix(data, decodedSize, validPrefixes)` add version-byte prefix validation on top of checksum. The prefix is the coin's version byte (e.g., `0x00` for Bitcoin P2PKH, `0x1C 0xB8` for ZCash P2PKH). If the decoded data doesn't start with one of the valid prefixes, decode returns an error.

**God-node ranking:** `Base58AddressDecodeToBytesPrefix` has fan-in 4 (called by `bitcoinDecodeToBytes`, `OntologyDecodeToBytes`, `TezosDecodeToBytes`, `TronDecodeToBytes`).

## Bech32 stack (2 levels)

### `bech32.go` — raw Bech32

`Bech32DecodeToBytes` and `Bech32EncodeToString` wrap the `btcsuite/btcutil/bech32` library with a simpler interface: just pass `hrp` (human-readable part) and the payload bytes.

### `bech32_address.go` — keyhash validation layer

`Bech32AddressDecodeToBytes(hrp, input, keyhashLength)` and `Bech32AddressEncodeToString(hrp, bytes, keyhashLength)` add keyhash length validation on top of raw Bech32. `keyhashLength` is the expected payload byte length (20 for most chains via `Bech32DefaultKeyhashLength`, 32 for Elrond via `ElrondKeyhashLength`).

**God-node ranking:** `Bech32AddressDecodeToBytes` has fan-in 5 (BNB, Cosmos, Elrond, IoTex, Zilliqa all call it directly).

## Why these are the integration points

The cross-domain bridge table shows 14 of the 25 bridges involve these primitives. This is expected: the design intentionally factors encoding into reusable primitives so coin-specific files stay thin (< 50 lines each). Any change to the checksum or prefix logic affects all coins that use that primitive.

## See Also
- [adding a new coin](../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:strong -->
- [init registration](../patterns/init-registration.md) <!-- rel:related -->
- [overview](../build/overview.md) <!-- rel:related -->
- [strategy](../tests/strategy.md) <!-- rel:related -->
