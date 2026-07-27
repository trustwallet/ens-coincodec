---
title: Test Infrastructure
category: tests
tags: [testing, test-runner, table-driven, coverage]
confidence: high
source: test_runner.go, test_runner_test.go, ethereum_test.go, bitcoin_test.go
updated: 2026-07-22
---

# Test Infrastructure

## test_runner.go helpers

The library provides shared test helpers in `test_runner.go` to keep per-coin test files thin and consistent:

```go
type TestcaseDecode struct {
    Input    string
    Expected string  // hex-encoded expected bytes
    Error    error   // expected error (nil means expect success)
}

type TestcaseEncode struct {
    Input    string  // hex-encoded input bytes
    Expected string  // expected string output
    Error    error
}

func RunTestsDecode(coinType uint32, testcases []TestcaseDecode) error
func RunTestsEncode(coinType uint32, testcases []TestcaseEncode) error
```

Each coin's `_test.go` file declares a slice of test cases and passes them to these helpers. This ensures every coin is tested through the same public `ToBytes`/`ToString` API path.

## 100% coverage requirement

The README states: "ensure you have 100% code coverage with your tests." The CI workflow enforces this via `go test -cover`. Each coin file is typically < 80 lines, so achieving full coverage is straightforward with a handful of test cases per format variant.

## Error path testing

Test cases with a non-nil `Error` field validate rejection of malformed inputs. Every coin has at least one error case.

## CI

The `.github/workflows/ci.yml` runs `go test ./...` on push/PR. Codecov is integrated for coverage tracking (see README badge).

## See Also
- [overview](../build/overview.md) <!-- rel:strong -->
- [readme](../ci/workflows/readme.md) <!-- rel:related -->
- [conventions](../code-conventions/code-style/conventions.md) <!-- rel:related -->
- [adding a new coin](../guides/adding-a-new-coin.md) <!-- rel:related -->
- [overview](../architecture/overview.md) <!-- rel:related -->
