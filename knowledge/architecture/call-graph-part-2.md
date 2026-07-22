---
category: architecture
confidence: low
documentType: explanation
scope: repo
contentHash: 9900fad35b60
tags: [architecture, domain]
source: architecture/call-graph.md
verified: 2026-07-22
splitPartIndex: 2
splitPartTotal: 2
canonical: false
synthetic: split-part
---

## Call Graph (Part 2)

### Cross-domain bridges

> Calls that cross a domain boundary — the integration points between subsystems. A sole link between two domains is the connection you'd least expect.

| From (domain) | To (domain) | Domain links |
|------|------|-----|
| `bitcoinDecodeToBytes` (bitcoin.go) | `Base58AddressDecodeToBytesPrefix` (base58_address.go) | 1 |
| `bitcoinEncodeToString` (bitcoin.go) | `Base58ChecksumEncode` (base58_checksum.go) | 1 |
| `Base58AddressDecodeToBytesPrefix` (base58_address.go) | `Base58ChecksumDecode` (base58_checksum.go) | 2 |
| `Base58AddressEncodeToStringPrefix` (base58_address.go) | `Base58ChecksumEncode` (base58_checksum.go) | 2 |
| `Base58ChecksumDecode` (base58_checksum.go) | `Base58Decode` (base58.go) | 2 |
| `Base58ChecksumEncode` (base58_checksum.go) | `Base58Encode` (base58.go) | 2 |
| `Bech32AddressDecodeToBytes` (bech32_address.go) | `Bech32DecodeToBytes` (bech32.go) | 2 |
| `Bech32AddressEncodeToString` (bech32_address.go) | `Bech32EncodeToString` (bech32.go) | 2 |
| `init` (bitcoin_forks.go) | `MakeBitcoinDecodeToBytes` (bitcoin.go) | 2 |
| `init` (bitcoin_forks.go) | `MakeBitcoinEncodeToString` (bitcoin.go) | 2 |
| `BNBDecodeToBytes` (bnb.go) | `Bech32AddressDecodeToBytes` (bech32_address.go) | 2 |
| `BNBEncodeToString` (bnb.go) | `Bech32AddressEncodeToString` (bech32_address.go) | 2 |
| `CosmosDecodeToBytes` (cosmos.go) | `Bech32AddressDecodeToBytes` (bech32_address.go) | 2 |
| `CosmosEncodeToString` (cosmos.go) | `Bech32AddressEncodeToString` (bech32_address.go) | 2 |
| `ElrondDecodeToBytes` (elrond.go) | `Bech32AddressDecodeToBytes` (bech32_address.go) | 2 |
| `ElrondEncodeToString` (elrond.go) | `Bech32AddressEncodeToString` (bech32_address.go) | 2 |
| `IoTexDecodeToBytes` (iotex.go) | `Bech32AddressDecodeToBytes` (bech32_address.go) | 2 |
| `IoTexEncodeToString` (iotex.go) | `Bech32AddressEncodeToString` (bech32_address.go) | 2 |
| `OntologyDecodeToBytes` (ontology.go) | `Base58AddressDecodeToBytesPrefix` (base58_address.go) | 2 |
| `OntologyEncodeToString` (ontology.go) | `Base58AddressEncodeToStringPrefix` (base58_address.go) | 2 |
| `PolkadotDecodeToBytes` (polkadot.go) | `SS58AddressDecodeToBytes` (ss58address.go) | 2 |
| `PolkadotEncodeToString` (polkadot.go) | `SS58AddressEncodeToString` (ss58address.go) | 2 |
| `SolanaDecodeToBytes` (solana.go) | `Base58Decode` (base58.go) | 2 |
| `SolanaEncodeToString` (solana.go) | `Base58Encode` (base58.go) | 2 |
| `SS58AddressDecodeToBytes` (ss58address.go) | `Base58Decode` (base58.go) | 2 |

### Possible duplicate entities (name variants)

> Symbol names that normalize identically — likely the same entity spelled inconsistently. Unify or distinguish in the docs.

- `BitcoinDecodeToBytes` / `bitcoinDecodeToBytes`
- `BitcoinEncodeToString` / `bitcoinEncodeToString`

### Dead-code candidates

> Non-exported symbols with no resolved callers, unreachable from any entry point. Static analysis cannot see dynamic dispatch — verify before removing.

- `init` (`bitcoin_cash.go`)
- `init` (`bitcoin_forks.go`)
- `init` (`bitcoin.go`)
- `init` (`bnb.go`)
- `init` (`cosmos.go`)
- `init` (`elrond.go`)
- `init` (`ethereum.go`)
- `init` (`filecoin.go`)
- `init` (`iotex.go`)
- `init` (`nimiq.go`)
- `init` (`ontology.go`)
- `init` (`polkadot.go`)
- `init` (`solana.go`)
- `init` (`stellar.go`)
- `init` (`tezos.go`)
- `init` (`tron.go`)
- `init` (`xrp.go`)
- `init` (`zilliqa.go`)

## See Also
- [adding a new coin](../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [init registration](../patterns/init-registration.md) <!-- rel:strong -->
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:related -->
- [readme](../ci/workflows/readme.md) <!-- rel:related -->
- [overview](../build/overview.md) <!-- rel:weak -->
