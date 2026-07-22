# Call Graph

<!-- sdd-knowledge-generated -->

> Deterministic call graph extracted via tree-sitter (no LLM). Direct calls are resolved by lexical scope + imports; **dynamic dispatch and ambiguous name matches are withheld** rather than guessed. The `## Calls` table is `EXTRACTED` (a single resolved target). Member calls (`obj.method()`) whose method name resolves to exactly one definition repo-wide are recovered separately under `## Inferred calls` (`INFERRED` — receiver type unverified, but a single plausible target).

## Calls

| Caller | Callee | Example args | Sites |
|--------|--------|--------------|-------|
| `Base58AddressDecodeToBytesPrefix` | `base58AddressIsValidDataPrefix(data: []byte, decodedSize: int, validPrefixes: [][]byte): error` | `(decoded, decodedSize, validPrefixes)` | 1 |
| `Base58AddressDecodeToBytesPrefix` | `Base58ChecksumDecode(b: string, alphabet: string): ([]byte, error)` | `(input, Base58DefaultAlphabet)` | 1 |
| `Base58AddressEncodeToStringPrefix` | `base58AddressIsValidDataPrefix(data: []byte, decodedSize: int, validPrefixes: [][]byte): error` | `(data, decodedSize, validPrefixes)` | 1 |
| `Base58AddressEncodeToStringPrefix` | `Base58ChecksumEncode(b: []byte, alphabet: string): string` | `(data, Base58DefaultAlphabet)` | 1 |
| `base58AddressIsValidDataPrefix` | `hasPrefix(data: []byte, prefix: []byte): bool` | `(data, prefix)` | 1 |
| `Base58ChecksumDecode` | `doubleSha256(b: []byte): []byte` | `(…)` | 1 |
| `Base58ChecksumDecode` | `Base58Decode(b: string, alphabet: string): ([]byte, error)` | `(b, alphabet)` | 1 |
| `Base58ChecksumEncode` | `doubleSha256(b: []byte): []byte` | `(b)` | 1 |
| `Base58ChecksumEncode` | `Base58Encode(b: []byte, alphabet: string): string` | `(b, alphabet)` | 1 |
| `Bech32AddressDecodeToBytes` | `Bech32DecodeToBytes(input: string, hrp: string): ([]byte, error)` | `(input, hrp)` | 1 |
| `Bech32AddressEncodeToString` | `Bech32EncodeToString(bytes: []byte, hrp: string): (string, error)` | `(bytes, hrp)` | 1 |
| `BitcoinCashDecodeToBytes` | `BitcoinDecodeToBytes(input: string): ([]byte, error)` | `(input)` | 1 |
| `BitcoinCashDecodeToBytes` | `buildP2PKHScript(bytes: []byte): []byte` | `(decoded)` | 1 |
| `BitcoinCashDecodeToBytes` | `buildP2SHScript(bytes: []byte): []byte` | `(decoded)` | 1 |
| `init` | `MakeBitcoinDecodeToBytes(config: *CoinConfig): func(string) ([]byte, error)` | `(…)` | 10 |
| `init` | `MakeBitcoinEncodeToString(config: *CoinConfig): func([]byte) (string, error)` | `(…)` | 10 |
| `bitcoinDecodeToBytes` | `Base58AddressDecodeToBytesPrefix(input: string, decodedSize: int, validPrefixes: [][]byte): ([]byte, error)` | `(input, …, [][]byte{config.P2PKHPrefix, config.P2SH)` | 1 |
| `BitcoinDecodeToBytes` | `bitcoinDecodeToBytes(input: string, config: *CoinConfig): ([]byte, error)` | `(input, …)` | 1 |
| `bitcoinDecodeToBytes` | `buildP2PKHScript(bytes: []byte): []byte` | `(data)` | 1 |
| `bitcoinDecodeToBytes` | `buildP2SHScript(bytes: []byte): []byte` | `(data)` | 1 |
| `bitcoinDecodeToBytes` | `buildWitnessScript(bytes: []byte): ([]byte, error)` | `(data)` | 1 |
| `bitcoinEncodeToString` | `Base58ChecksumEncode(b: []byte, alphabet: string): string` | `(…, Base58DefaultAlphabet)` | 2 |
| `BitcoinEncodeToString` | `bitcoinEncodeToString(input: []byte, config: *CoinConfig): (string, error)` | `(input, …)` | 1 |
| `bitcoinEncodeToString` | `replacePrefix(input: []byte, oldPrefix: []byte, newPrefix: []byte): []byte` | `(input, P2PKH_SCRIPT_PREFIX, config.P2PKHPrefix)` | 2 |
| `MakeBitcoinDecodeToBytes` | `bitcoinDecodeToBytes(input: string, config: *CoinConfig): ([]byte, error)` | `(input, config)` | 1 |
| `MakeBitcoinEncodeToString` | `bitcoinEncodeToString(input: []byte, config: *CoinConfig): (string, error)` | `(input, config)` | 1 |
| `BNBDecodeToBytes` | `Bech32AddressDecodeToBytes(hrp: string, input: string, keyhashLength: int): ([]byte, error)` | `(hrpBNB, input, Bech32DefaultKeyhashLength)` | 1 |
| `BNBEncodeToString` | `Bech32AddressEncodeToString(hrp: string, bytes: []byte, keyhashLength: int): (string, error)` | `(hrpBNB, bytes, Bech32DefaultKeyhashLength)` | 1 |
| `CosmosDecodeToBytes` | `Bech32AddressDecodeToBytes(hrp: string, input: string, keyhashLength: int): ([]byte, error)` | `(hrp, input, Bech32DefaultKeyhashLength)` | 1 |
| `CosmosEncodeToString` | `Bech32AddressEncodeToString(hrp: string, bytes: []byte, keyhashLength: int): (string, error)` | `(hrp, bytes, Bech32DefaultKeyhashLength)` | 1 |
| `ElrondDecodeToBytes` | `Bech32AddressDecodeToBytes(hrp: string, input: string, keyhashLength: int): ([]byte, error)` | `(hrpElrond, input, ElrondKeyhashLength)` | 1 |
| `ElrondEncodeToString` | `Bech32AddressEncodeToString(hrp: string, bytes: []byte, keyhashLength: int): (string, error)` | `(hrpElrond, bytes, ElrondKeyhashLength)` | 1 |
| `EtherToBytes` | `EtherToString(input: []byte): (string, error)` | `(output)` | 1 |
| `FilecoinDecodeToBytes` | `isValidString(input: string): error` | `(input)` | 1 |
| `FilecoinDecodeToBytes` | `parseType(ascii: byte): Type` | `(…)` | 1 |
| `FilecoinDecodeToBytes` | `payloadSize(t: Type): byte` | `(typ)` | 1 |
| `FilecoinEncodeToString` | `filecoinComputeChecksum(address: []byte): []byte` | `(data)` | 1 |
| `FilecoinEncodeToString` | `getType(raw: byte): Type` | `(…)` | 1 |
| `FilecoinEncodeToString` | `payloadSize(t: Type): byte` | `(t)` | 1 |
| `FilecoinEncodeToString` | `typeAscii(typ: Type): string` | `(t)` | 1 |
| `isValidBase32` | `filecoinComputeChecksum(address: []byte): []byte` | `(address)` | 1 |
| `isValidBase32` | `payloadSize(t: Type): byte` | `(t)` | 1 |
| `isValidString` | `isValidBase32(input: string, t: Type): error` | `(input, typ)` | 1 |
| `isValidString` | `isValidID(id: string): bool` | `(input)` | 1 |
| `isValidString` | `parseType(ascii: byte): Type` | `(…)` | 1 |
| `parseType` | `Type` | `(…)` | 1 |
| `IoTexDecodeToBytes` | `Bech32AddressDecodeToBytes(hrp: string, input: string, keyhashLength: int): ([]byte, error)` | `(hrpIOTEX, input, Bech32DefaultKeyhashLength)` | 1 |
| `IoTexEncodeToString` | `Bech32AddressEncodeToString(hrp: string, bytes: []byte, keyhashLength: int): (string, error)` | `(hrpIOTEX, bytes, Bech32DefaultKeyhashLength)` | 1 |
| `nimiqChecksum` | `nimiqToNumber(input: string): (string, error)` | `(…)` | 1 |
| `NimiqDecodeToBytes` | `nimiqChecksum(input: string): (uint8, error)` | `(input)` | 1 |
| `NimiqEncodeToString` | `nimiqChecksum(input: string): (uint8, error)` | `(noSpaces)` | 1 |
| `OntologyDecodeToBytes` | `Base58AddressDecodeToBytesPrefix(input: string, decodedSize: int, validPrefixes: [][]byte): ([]byte, error)` | `(input, OntologyAddressLength, ontologyPrefixes)` | 1 |
| `OntologyEncodeToString` | `Base58AddressEncodeToStringPrefix(data: []byte, decodedSize: int, validPrefixes: [][]byte): (string, error)` | `(bytes, OntologyAddressLength, ontologyPrefixes)` | 1 |
| `PolkadotDecodeToBytes` | `SS58AddressDecodeToBytes(input: string, network: byte): ([]byte, error)` | `(input, PolkadotNetwork)` | 1 |
| `PolkadotEncodeToString` | `SS58AddressEncodeToString(data: []byte, network: byte): (string, error)` | `(data, PolkadotNetwork)` | 1 |
| `SolanaDecodeToBytes` | `Base58Decode(b: string, alphabet: string): ([]byte, error)` | `(input, Base58DefaultAlphabet)` | 1 |
| `SolanaEncodeToString` | `Base58Encode(b: []byte, alphabet: string): string` | `(data, Base58DefaultAlphabet)` | 1 |
| `SS58AddressDecodeToBytes` | `Base58Decode(b: string, alphabet: string): ([]byte, error)` | `(input, Base58DefaultAlphabet)` | 1 |
| `SS58AddressDecodeToBytes` | `computeChecksum(data: []byte): []byte` | `(…)` | 1 |
| `SS58AddressEncodeToString` | `Base58Encode(b: []byte, alphabet: string): string` | `(checksummed, Base58DefaultAlphabet)` | 1 |
| `SS58AddressEncodeToString` | `computeChecksum(data: []byte): []byte` | `(networked)` | 1 |
| `StellarDecodeToBytes` | `crc16(bytes: []byte): []byte` | `(…)` | 1 |
| `StellarEncodeToString` | `crc16(bytes: []byte): []byte` | `(data)` | 1 |
| `RunTestDecode` | `errorString(err: error): string` | `(err)` | 1 |
| `RunTestEncode` | `errorString(err: error): string` | `(err)` | 1 |
| `RunTestsDecode` | `RunTestDecode(coinType: uint32, tt: TestcaseDecode): error` | `(coinType, tt)` | 1 |
| `RunTestsEncode` | `RunTestEncode(coinType: uint32, tt: TestcaseEncode): error` | `(coinType, tt)` | 1 |
| `TezosDecodeToBytes` | `Base58AddressDecodeToBytesPrefix(input: string, decodedSize: int, validPrefixes: [][]byte): ([]byte, error)` | `(input, TezosAddressLength, tzPrefixes)` | 1 |
| `TezosEncodeToString` | `Base58AddressEncodeToStringPrefix(data: []byte, decodedSize: int, validPrefixes: [][]byte): (string, error)` | `(bytes, TezosAddressLength, tzPrefixes)` | 1 |
| `TronDecodeToBytes` | `Base58AddressDecodeToBytesPrefix(input: string, decodedSize: int, validPrefixes: [][]byte): ([]byte, error)` | `(input, TronAddressLength, tronPrefixes)` | 1 |
| `TronEncodeToString` | `Base58AddressEncodeToStringPrefix(data: []byte, decodedSize: int, validPrefixes: [][]byte): (string, error)` | `(bytes, TronAddressLength, tronPrefixes)` | 1 |
| `XRPDecodeToBytes` | `Base58ChecksumDecode(b: string, alphabet: string): ([]byte, error)` | `(input, xrpAlphabet)` | 1 |
| `XRPEncodeToString` | `Base58ChecksumEncode(b: []byte, alphabet: string): string` | `(bytes, xrpAlphabet)` | 1 |
| `ZilliqaDecodeToBytes` | `Bech32AddressDecodeToBytes(hrp: string, input: string, keyhashLength: int): ([]byte, error)` | `(hrpZIL, input, Bech32DefaultKeyhashLength)` | 1 |
| `ZilliqaEncodeToString` | `Bech32AddressEncodeToString(hrp: string, bytes: []byte, keyhashLength: int): (string, error)` | `(hrpZIL, bytes, Bech32DefaultKeyhashLength)` | 1 |

## Callers (reverse)

| Symbol | Called by |
|--------|-----------|
| `Base58AddressDecodeToBytesPrefix` | `OntologyDecodeToBytes`, `TezosDecodeToBytes`, `TronDecodeToBytes`, `bitcoinDecodeToBytes` |
| `Base58AddressEncodeToStringPrefix` | `OntologyEncodeToString`, `TezosEncodeToString`, `TronEncodeToString` |
| `base58AddressIsValidDataPrefix` | `Base58AddressDecodeToBytesPrefix`, `Base58AddressEncodeToStringPrefix` |
| `Base58ChecksumDecode` | `Base58AddressDecodeToBytesPrefix`, `XRPDecodeToBytes` |
| `Base58ChecksumEncode` | `Base58AddressEncodeToStringPrefix`, `XRPEncodeToString`, `bitcoinEncodeToString` |
| `Base58Decode` | `Base58ChecksumDecode`, `SS58AddressDecodeToBytes`, `SolanaDecodeToBytes` |
| `Base58Encode` | `Base58ChecksumEncode`, `SS58AddressEncodeToString`, `SolanaEncodeToString` |
| `Bech32AddressDecodeToBytes` | `BNBDecodeToBytes`, `CosmosDecodeToBytes`, `ElrondDecodeToBytes`, `IoTexDecodeToBytes`, `ZilliqaDecodeToBytes` |
| `Bech32AddressEncodeToString` | `BNBEncodeToString`, `CosmosEncodeToString`, `ElrondEncodeToString`, `IoTexEncodeToString`, `ZilliqaEncodeToString` |
| `Bech32DecodeToBytes` | `Bech32AddressDecodeToBytes` |
| `Bech32EncodeToString` | `Bech32AddressEncodeToString` |
| `bitcoinDecodeToBytes` | `BitcoinDecodeToBytes`, `MakeBitcoinDecodeToBytes` |
| `BitcoinDecodeToBytes` | `BitcoinCashDecodeToBytes` |
| `bitcoinEncodeToString` | `BitcoinEncodeToString`, `MakeBitcoinEncodeToString` |
| `buildP2PKHScript` | `BitcoinCashDecodeToBytes`, `bitcoinDecodeToBytes` |
| `buildP2SHScript` | `BitcoinCashDecodeToBytes`, `bitcoinDecodeToBytes` |
| `buildWitnessScript` | `bitcoinDecodeToBytes` |
| `computeChecksum` | `SS58AddressDecodeToBytes`, `SS58AddressEncodeToString` |
| `crc16` | `StellarDecodeToBytes`, `StellarEncodeToString` |
| `doubleSha256` | `Base58ChecksumDecode`, `Base58ChecksumEncode` |
| `errorString` | `RunTestDecode`, `RunTestEncode` |
| `EtherToString` | `EtherToBytes` |
| `filecoinComputeChecksum` | `FilecoinEncodeToString`, `isValidBase32` |
| `getType` | `FilecoinEncodeToString` |
| `hasPrefix` | `base58AddressIsValidDataPrefix` |
| `isValidBase32` | `isValidString` |
| `isValidID` | `isValidString` |
| `isValidString` | `FilecoinDecodeToBytes` |
| `MakeBitcoinDecodeToBytes` | `init` |
| `MakeBitcoinEncodeToString` | `init` |
| `nimiqChecksum` | `NimiqDecodeToBytes`, `NimiqEncodeToString` |
| `nimiqToNumber` | `nimiqChecksum` |
| `parseType` | `FilecoinDecodeToBytes`, `isValidString` |
| `payloadSize` | `FilecoinDecodeToBytes`, `FilecoinEncodeToString`, `isValidBase32` |
| `replacePrefix` | `bitcoinEncodeToString` |
| `RunTestDecode` | `RunTestsDecode` |
| `RunTestEncode` | `RunTestsEncode` |
| `SS58AddressDecodeToBytes` | `PolkadotDecodeToBytes` |
| `SS58AddressEncodeToString` | `PolkadotEncodeToString` |
| `Type` | `parseType` |
| `typeAscii` | `FilecoinEncodeToString` |

## Graph analytics

- Call edges: **75** across **98** symbols
- Reachable from entry points (exported symbols + routes): **80**
- Dependency cycles: **0**
- Cross-domain bridges: **25**
- Dead-code candidates (non-exported, zero callers, unreachable): **18**

### Most-coupled symbols (god-node ranking)

| Symbol | Fan-in | Fan-out | Degree |
|--------|--------|---------|--------|
| `Base58AddressDecodeToBytesPrefix` | 4 | 2 | 6 |
| `Bech32AddressDecodeToBytes` | 5 | 1 | 6 |
| `Bech32AddressEncodeToString` | 5 | 1 | 6 |
| `bitcoinDecodeToBytes` | 2 | 4 | 6 |
| `Base58AddressEncodeToStringPrefix` | 3 | 2 | 5 |
| `Base58ChecksumEncode` | 3 | 2 | 5 |
| `Base58ChecksumDecode` | 2 | 2 | 4 |
| `bitcoinEncodeToString` | 2 | 2 | 4 |
| `FilecoinEncodeToString` | 0 | 4 | 4 |
| `isValidString` | 1 | 3 | 4 |
| `base58AddressIsValidDataPrefix` | 2 | 1 | 3 |
| `Base58Decode` | 3 | 0 | 3 |
| `Base58Encode` | 3 | 0 | 3 |
| `BitcoinCashDecodeToBytes` | 0 | 3 | 3 |
| `FilecoinDecodeToBytes` | 0 | 3 | 3 |

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

