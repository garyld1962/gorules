# gorules

[![CI](https://github.com/garyld1962/gorules/actions/workflows/ci.yml/badge.svg)](https://github.com/garyld1962/gorules/actions/workflows/ci.yml)

A Go business-rules engine from 2016 that evaluates a custom DSL (`ANY orderItems.weight IsGreaterThan '10' AND country IsEqualTo 'USA'`) against JSON data. It's here as the **"before" half of a regenerate-don't-patch case study**: in 2026 it was put through an AI multi-specialist audit, rehabilitated to a trustworthy test suite, then used as the specification for a clean-room Rust regeneration — **[rustrules](https://github.com/garyld1962/rustrules)** — rather than patched in place.

## The arc

1. **Audit (2026-02-20):** four AI specialists (security, API patterns, test quality, code quality) reviewed all 28 files — verdict **FAIL**, [37 findings: 13 blockers, 11 high, 10 medium, 3 low](docs/code-reviews/2026-02-20-full-review.md), including real evaluation bugs (`>=` implemented as `>`, silent type-coercion failures).
2. **Test rehab:** the suite grew **29 → 124 tests** so the engine's actual behavior — bugs included — was pinned down before anything was rewritten.
3. **Regeneration:** with behavior captured, [rustrules](https://github.com/garyld1962/rustrules) was generated clean-room in Rust — fixing 12+ of the audited bugs and adding 15 operators, verified by its own 121-test suite.

The audit trail lives in [`docs/code-reviews/`](docs/code-reviews/). The point of keeping this repo public, bugs and all, is that the case study is verifiable instead of anecdotal.

## Run it (60 seconds)

```bash
go test -v ./...   # 124 tests
```

The sample rules live in `evaluator_test.go`; the HTTP API server is in `api/`.

## License

MIT
