---
title: "init() Registration Pattern — NOT Dead Code"
category: patterns
tags: [go-init, registry, dead-code, auto-registration]
confidence: high
source: codec.go, ethereum.go, bitcoin.go, bitcoin_forks.go, bnb.go
updated: 2026-07-22
---

# init() Registration Pattern — NOT Dead Code

## The pattern

Every coin-specific file contains a `func init()` that registers its codec functions into the global maps in `codec.go`:

```go
// ethereum.go
func init() {
    for _, c := range coin.Coins {
        if coin.IsEVM(c.ID) {
            toBytesMap[uint32(c.ID)] = EtherToBytes
            toStringMap[uint32(c.ID)] = EtherToString
        }
    }
}
```

These `init()` functions are called automatically by the Go runtime when the package is first imported — no explicit call is needed. This is the idiomatic Go way to implement a self-registering plugin or codec system.

## Why the static analysis flags them as dead code

`sdd-knowledge --full` marks all `init()` functions as "dead-code candidates (non-exported, zero callers)." This is a **false positive**. The static AST analyzer correctly notes that `init()` is never explicitly called from source code — but it cannot know that Go's runtime calls it at package initialization time. `init()` functions in Go are reserved words that the runtime treats specially.

**Rule: never delete any `init()` function in this repo.** Each one registers one or more coin codecs; deleting it silently breaks those coins.

## EVM bulk registration

`ethereum.go`'s `init()` is unique: it loops over ALL coins in `go-primitives/coin.Coins` and registers `EtherToBytes`/`EtherToString` for every coin where `coin.IsEVM(c.ID)` returns true. This means:
- Adding a new EVM chain to `go-primitives` automatically adds support here with no code change
- All EVM addresses are encoded identically (raw 20-byte hex + EIP-55 checksum)
- The total number of registered coin types at runtime is much larger than the number of source files

## Adding a new coin

1. Create `<coin>.go` with `DecodeToBytes` and `EncodeToString` functions
2. Add an `init()` that registers them in `toBytesMap[coin.COINTYPE]` and `toStringMap[coin.COINTYPE]`
3. Add `<coin>_test.go` with full coverage using the `RunTests*` helpers from `test_runner.go`

## See Also
- [adding a new coin](../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [overview](../architecture/overview.md) <!-- rel:strong -->
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:related -->
- [bitcoin codec](../architecture/bitcoin-codec.md) <!-- rel:related -->
- [encoding primitives](../architecture/encoding-primitives.md) <!-- rel:related -->
