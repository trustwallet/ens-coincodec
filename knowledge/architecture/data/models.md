# Data Models & Schemas

<!-- sdd-knowledge-generated -->

> Field-level shape of data models extracted via tree-sitter: TS interfaces / type aliases, Zod `z.object` schemas, Go/Rust/Swift structs, Kotlin data classes, and Python dataclasses. Scoped to domain data — UI views/props, view-models, design tokens (theme/style/colors), and constant/identifier namespaces are excluded. Deterministic, no LLM.

## CoinConfig

_struct · `bitcoin.go`:13_

| Field | Type | Optional |
|-------|------|----------|
| `P2PKHPrefix` | `[]byte` | no |
| `P2SHPPrefix` | `[]byte` | no |
| `HRP` | `string` | no |

## TestcaseDecode

_struct · `test_runner.go`:18_

| Field | Type | Optional |
|-------|------|----------|
| `name` | `string` | no |
| `input` | `string` | no |
| `output` | `string` | no |
| `err` | `error` | no |

## TestcaseEncode

_struct · `test_runner.go`:11_

| Field | Type | Optional |
|-------|------|----------|
| `name` | `string` | no |
| `input` | `string` | no |
| `output` | `string` | no |
| `err` | `error` | no |

## See Also
- [adding a new coin](../../guides/adding-a-new-coin.md) <!-- rel:strong -->
- [conventions](../../code-conventions/code-style/conventions.md) <!-- rel:strong -->
- [strategy](../../tests/strategy.md) <!-- rel:related -->
- [constitution](../../constitution.md) <!-- rel:related -->
- [readme](../../ci/workflows/readme.md) <!-- rel:related -->
