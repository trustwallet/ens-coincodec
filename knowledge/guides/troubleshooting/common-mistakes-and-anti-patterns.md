---
category: guides
subcategory: troubleshooting
confidence: low
documentType: how-to
scope: org
contentHash: c1d2a4e3d386
tags: [anti-pattern, troubleshooting, faq]
source: (synthesized)
verified: 2026-07-22
synthetic: synthesized-faq
---

## Common Mistakes and Anti-Patterns

<!-- sdd-knowledge-synthesized -->

> Auto-generated from anti-patterns found across the knowledge base.

### Code-conventions

**From [Anti-Patterns (failed approaches)](../../code-conventions/anti-patterns-failed-approaches.md):**
<!-- Add failed approaches here. Each anti-pattern should include:
- **type**: anti-pattern
- **discovered**: YYYY-MM-DD

### Guides

**From [Adding a New Coin](../../guides/adding-a-new-coin.md):**
3. **Avoid large imports** — the README explicitly says to copy relevant code rather than importing a large project just for address conversion.

4. **Create `<coin>_test.go`** with 100% coverage using `RunTestsDecode`/`RunTestsEncode` from `test_runner.go`.

### Patterns

**From [init() Registration Pattern — NOT Dead Code](../../patterns/init-registration-pattern-not-dead-code.md):**
`sdd-knowledge --full` marks all `init()` functions as "dead-code candidates (non-exported, zero callers)." This is a **false positive**. The static AST analyzer correctly notes that `init()` is never explicitly called from source code — but it cannot know that Go's runtime calls it at package initialization time. `init()` functions in Go are reserved words that the runtime treats specially.

**Rule: never delete any `init()` function in this repo.** Each one registers one or more coin codecs; deleting it silently breaks those coins.

