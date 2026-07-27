---
title: "init() Functions Are Not Dead Code in ens-coincodec"
date: 2026-07-22
pr: TBD
area: [architecture, static-analysis]
files: [codec.go, ethereum.go, bitcoin.go, bitcoin_forks.go, bnb.go, cosmos.go, elrond.go, filecoin.go, iotex.go, nimiq.go, ontology.go, polkadot.go, solana.go, stellar.go, tezos.go, tron.go, xrp.go, zilliqa.go]
symptom: "sdd-knowledge --full reports 18 `init()` functions as dead-code candidates (non-exported, zero callers, unreachable)"
tags: [go, init, dead-code, false-positive, codec, registry]
summary: "All init() functions in this repo are Go package-initialization hooks that auto-register coin codecs — they are called by the Go runtime, not from source code, so static AST analysis always flags them as dead code. Never delete them."
---

## The Pattern / The Bug

`sdd-knowledge --full`'s dead-code detection correctly observes that every `init()` function in the repo has zero callers in the source code. It flags all 18 of them as "dead-code candidates."

This is a **false positive**. In Go, `func init()` is a reserved function signature that the runtime calls automatically when the package is first imported — no explicit call is ever written. The package-level `init()` in `ethereum.go`, for example, iterates over all EVM coins and registers `EtherToBytes`/`EtherToString` into the global dispatch maps. Without this, `ToBytes` and `ToString` would return "unhandled coin type" for every EVM chain.

## The Rule

**Never delete any `init()` function in this repo.** Each one silently registers one or more coin codecs. Deleting one breaks address encoding for those coins with no compile error, no test failure at import time, and only a runtime "unhandled coin type" error when that coin type is actually used.

## Detection

If a static analysis tool (sdd-knowledge, staticcheck, or a custom linter) flags an `init()` function as unused/dead:
1. Check whether it registers entries in `toBytesMap` or `toStringMap` in `codec.go`.
2. If yes: the tool is wrong. Suppress the warning; do not remove the function.

## Related

- `knowledge/patterns/init-registration.md` — full explanation of the registration pattern
- `knowledge/architecture/overview.md` — registry architecture
