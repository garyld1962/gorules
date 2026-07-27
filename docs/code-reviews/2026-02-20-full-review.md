# Code Review: gorules — Full Codebase Review

**Date:** 2026-02-20
**Scope:** All `.go` files (16 source, 11 test, 1 API server)
**Mode:** Full review (all files on master)
**Specialists:** Security, API Patterns, Test Quality, Code Quality
**Verdict:** **FAIL** (13 Blockers, 11 High, 10 Medium, 3 Low)

---

## Verdict Summary

| Severity | Count |
|----------|-------|
| Blocker  | 13    |
| High     | 11    |
| Medium   | 10    |
| Low      | 3     |
| **Total**| **37** |

---

## BLOCKER Findings

### B1. `isGreaterOrEqual` uses `>` instead of `>=`
- **File:** `utility.go:59`
- **Domain:** Code Quality
- **Description:** Returns `val1 > val2` — boundary case where values are equal returns `false` instead of `true`.
- **Impact:** Rules like `"age IsGreaterThanOrEqualTo '18'"` silently reject equal values.
- **Fix:** Change to `return val1 >= val2, nil`

### B2. `isLessOrEqual` uses `<` instead of `<=`
- **File:** `utility.go:74`
- **Domain:** Code Quality
- **Description:** Returns `val1 < val2` — same boundary-value bug as B1.
- **Fix:** Change to `return val1 <= val2, nil`

### B3. `isFalse` returns same result as `isTrue`
- **File:** `utility.go:106-113`
- **Domain:** Code Quality
- **Description:** Parses `target` to bool and returns `r` — should return `!r`.
- **Impact:** `isFalse("", "true")` returns `true` — operator is semantically inverted.
- **Fix:** Change line 112 to `return !r, nil`

### B4. `selectValue` recursive path broken — wrong delimiter + skips segment
- **File:** `utility.go:226-227`
- **Domain:** Code Quality
- **Description:** `strings.Join(propertyNames[2:], ",")` uses `","` instead of `"."`, and `[2:]` skips `propertyNames[1]` entirely.
- **Impact:** All nested JSON path evaluation (e.g., `"testobj.id"`) returns nil or wrong value. Corrupts every selector-based rule against nested data.
- **Fix:** Use `propertyNames[1:]` with `"."` delimiter, or rewrite iteratively.

### B5. `IsGreaterThan` missing from `operatorFuncList` — nil function panic
- **File:** `operator.go:61-63`, called from `expression.go:19-20`
- **Domain:** Security / Code Quality
- **Description:** `IsGreaterThan` is declared as a valid operator constant but has no entry in the dispatch map. Accessing it returns nil, and the immediate call panics.
- **Impact:** Any rule using `IsGreaterThan` crashes the server.
- **Fix:** Add `IsGreaterThan: isGreater` to `operatorFuncList` and add nil guard at call site.

### B6. `divide()` has no zero-division guard — runtime panic
- **File:** `math_utility.go:9-11`
- **Domain:** Security / Code Quality
- **Description:** Integer division by zero in Go panics at runtime. No guard exists.
- **Impact:** `|DIV 'x' '0'|` sent to `/expression` crashes the server.
- **Fix:** Add `if operand2 == 0 { return 0, errors.New("division by zero") }`

### B7. Unsafe `.(string)` type assertion on JSON values that may be `float64`
- **File:** `value.go:49`, `expression.go:94-97`
- **Domain:** Security / Code Quality
- **Description:** `Path.Evaluate` and `createConjuntionExprFromCollectionStmt` unconditionally assert `.(string)` on values from `selectValue`. Go's JSON decoder produces `float64` for numbers, `bool` for booleans.
- **Impact:** Any rule referencing a numeric JSON field (id, zip, quantity) panics. Test data itself contains `id: 25`, `zip5: 33076`.
- **Fix:** Use `fmt.Sprintf("%v", raw)` or type switch. Add comma-ok checks for all assertions.

### B8. `os.Exit(1)` in library code kills entire process
- **File:** `rule_processor.go:35-38`
- **Domain:** Security / API
- **Description:** `RuleFromFile.Fetch()` calls `os.Exit(1)` when `./rulez.json` is missing. Bypasses all deferred functions and recovery middleware.
- **Impact:** One request for a missing rule file takes the server offline permanently. No recovery possible.
- **Fix:** Return `(string, error)` and propagate through the call chain.

### B9. `firstExpr` panics on empty slice — no bounds check
- **File:** `utility.go:436-438`
- **Domain:** Code Quality
- **Description:** `return list[0]` with no length check. Called from `rule.go:69` (`determineSeedAccum`).
- **Impact:** Evaluating a rule with zero expressions (empty string, standalone conjunction) panics.
- **Fix:** Add length guard and return error.

### B10. Unsafe type assertions in `NewRuleProcessor` — multiple panic paths
- **File:** `rule_processor.go:121-132`
- **Domain:** Security / API
- **Description:** Three bare type assertions: `input.([]string)` (JSON gives `[]interface{}`), `input.(string)` (fails for numbers/bools/nil). Called directly from `/rule` HTTP handler with user input.
- **Impact:** Any non-string `rule` field in API request crashes the server.
- **Fix:** Guard all type assertions with comma-ok pattern.

### B11. `panic()` used as error handling in 5+ API-reachable code paths
- **Files:** `expression.go:30,38,89`, `statement.go:87`, `utility.go:289`
- **Domain:** Security / Code Quality
- **Description:** Invalid operator names, conjunction strings, and malformed JSON all trigger `panic()` instead of returning errors. These paths are reachable from HTTP endpoints.
- **Impact:** Any invalid DSL string crashes the handler goroutine. Gin's recovery middleware may catch it, but internal state leaks and no meaningful error response is produced.
- **Fix:** Replace all `panic()` with error returns and propagate.

### B12. Silent failure on API binding error — no response sent
- **File:** `api/api.go:44-50, 54-61`
- **Domain:** API
- **Description:** Both POST handlers use `if c.BindJSON(&json) == nil { ... }` with no else branch. When binding fails, no response body is sent.
- **Impact:** Clients sending invalid JSON receive no response. API is unusable for error cases.
- **Fix:** Use `ShouldBindJSON` with explicit error response and `return`.

### B13. Nil function call for unrecognized math/string operators
- **File:** `value.go:100-101, 137-138`
- **Domain:** Security
- **Description:** `mathOperatorFuncList[m.operator]` and `stringOperatorFuncList[m.operator]` return nil for unrecognized operators. The nil function is immediately called, causing a panic.
- **Impact:** Sending an unknown operator in an expression crashes the server.
- **Fix:** Add nil/ok guard before calling the function.

---

## HIGH Findings

### H1. Tests always pass — hardcoded assertions ignore results
- **Files:** `evaluator_test.go:115,126`, `rule_processor_test.go:64,71`
- **Domain:** Test Quality
- **Description:** Four tests call production code, capture result in a variable, then assert on a hardcoded literal (`assert.False(t, false)`, `assert.True(t, true)`) instead of the result variable.
- **Impact:** Regressions in `EvaluateRules`, `RuleCollection.Process`, and `RuleObjectCollection.Process` are invisible. CI green signal is a lie.
- **Fix:** Replace with `assert.False(t, result)` / assertions on actual `n`.

### H2. `isNull`/`isNotNull` check wrong parameter
- **File:** `utility.go:87-95`
- **Domain:** Code Quality
- **Description:** Both check `target` instead of `value`. Semantic intent is to check the data field (value), not the rule operand (target).
- **Impact:** IsNull evaluations check the wrong operand — always returns based on target emptiness.
- **Fix:** Check `value` instead of `target`.

### H3. `in()` does substring match, not set membership
- **File:** `utility.go:16-18`
- **Domain:** Code Quality
- **Description:** Uses `strings.Contains(v, t)` — "US" matches "USA".
- **Fix:** Split on delimiter and check for exact membership.

### H4. `isInArray` returns `true` for unhandled types
- **File:** `utility.go:209-213`
- **Domain:** Code Quality
- **Description:** The `default` case is empty, falls through to `return true, nil`.
- **Impact:** Any unrecognized type is falsely reported as "in the array".
- **Fix:** Return `false, error` for unknown types.

### H5. No middleware — no auth, CORS, rate limiting, structured logging
- **File:** `api/api.go:26`
- **Domain:** API / Security
- **Description:** `gin.Default()` provides only basic logger and recovery. The API is completely open.
- **Impact:** Any network client can submit arbitrary rule evaluations with no limits.
- **Fix:** Add appropriate middleware stack.

### H6. Hardcoded port `:8080` and file path `./rulez.json`
- **File:** `api/api.go:63`, `rule_processor.go:35`
- **Domain:** API
- **Description:** No configuration mechanism. Port and path are hardcoded.
- **Fix:** Read from environment variables with defaults.

### H7. No graceful shutdown
- **File:** `api/api.go:63`
- **Domain:** API
- **Description:** `router.Run()` blocks with no signal handling. SIGTERM kills in-flight requests.
- **Fix:** Use `http.Server` with `signal.Notify` and `srv.Shutdown()`.

### H8. Exported `var` functions can be overwritten by callers
- **Files:** `evaluator.go:29-31`, `parser.go:17-20`, `utility.go:295`
- **Domain:** Code Quality / Security
- **Description:** `DSLEvaluator`, `EvaluateRules`, `ParseDSL`, etc. are `var` not `func`. Any caller can nil them out.
- **Fix:** Convert to `func` declarations.

### H9. `isArray`/`isMap` panic on nil input
- **File:** `utility.go:146-151`
- **Domain:** Security
- **Description:** `reflect.TypeOf(nil).Kind()` panics. Called from `NewRuleProcessor` with user input.
- **Fix:** Add nil check before reflect call.

### H10. Unbounded recursion in `selectValue` + broken logic — stack overflow potential
- **File:** `utility.go:215-228`
- **Domain:** Security
- **Description:** The broken delimiter (`","`) and segment skipping create a recursion that doesn't properly reduce the path. Deep paths could stack overflow.
- **Fix:** Rewrite iteratively (also fixes B4).

### H11. No `go.mod` — legacy Glide, unmaintained Gin v1, deprecated deps
- **File:** `glide.yaml`
- **Domain:** API
- **Description:** No module system. `gopkg.in/gin-gonic/gin.v1` is from 2016, unpatched. `stretchr/stew` and `stretchr/signature` are unused/obsolete. `ioutil.ReadFile` deprecated.
- **Fix:** `go mod init`, update Gin, replace `stew` with `encoding/json`, remove `signature`.

---

## MEDIUM Findings

### M1. Empty test files — zero coverage for expression, string_slice, functionexecutor
- **Files:** `expression_test.go`, `string_slice_test.go`, `functionexecutor_test.go`
- **Domain:** Test Quality
- **Description:** All tests commented out or absent. Expression evaluation (the heart of the engine) has zero direct unit tests.
- **Fix:** Restore/write tests or delete empty files.

### M2. No error-path tests anywhere
- **Domain:** Test Quality
- **Description:** No test for IsGreaterThan dispatch, divide-by-zero, isFalse, isGreaterOrEqual boundary, or Path.Evaluate with non-string values. All documented bugs go undetected by tests.
- **Fix:** Add boundary and error-path test cases.

### M3. `TestFetchFromFile` depends on external file + `os.Exit` kills test process
- **File:** `rule_processor_test.go:55`
- **Domain:** Test Quality
- **Description:** Non-hermetic test. Missing `./rulez.json` causes `os.Exit(1)`, killing the entire `go test` process.
- **Fix:** Add skip guard; fix production code to return error.

### M4. 40+ `fmt.Println` debug statements in production code
- **Files:** Throughout all source files
- **Domain:** Code Quality / Security
- **Description:** Every rule evaluation logs full input, intermediate state, and results to stdout. Leaks business data, PII, and internal state.
- **Fix:** Remove all. Replace with `log/slog` at appropriate levels.

### M5. `ConjunctionExpression.Add` uses value receiver
- **File:** `expression.go:68-71`
- **Domain:** Code Quality
- **Description:** Value receiver means `append` may not modify the original. Error-prone if return value is ignored.
- **Fix:** Use pointer receiver.

### M6. Exported dead code and empty files
- **Files:** `utility.go:392` (`RRecursiveEncodeWithPipe`), `utility.go:394-407` (`recursiveSpaceEncode`), `functionexecutor.go`
- **Domain:** Code Quality
- **Description:** Double-R typo export, unused function, empty file.
- **Fix:** Remove all.

### M7. 12+ spelling errors in identifiers and comments
- **Files:** Multiple (see handoff.md for full list)
- **Domain:** Code Quality
- **Description:** `spiltWithSpace`, `accumlator`, `conjuntion`, `linetoWorkOn`, `lengt`, `arry`, `delimitter`, `implemente`, wrong function names in comments.
- **Fix:** Rename all. Note: renaming exported identifiers is a breaking API change.

### M8. `MathExpression.String()` panics when operand is a Path
- **File:** `value.go:106-108, 128-130`
- **Domain:** Code Quality
- **Description:** Calls `Evaluate` with `make([]interface{}, 0)` as data. If operand is a `Path`, it asserts `.(map[string]interface{})` on a slice and panics.
- **Fix:** Return descriptive string without evaluating.

### M9. `RuleMessage.Rule` and `ExpressionMessage.Expression` typed as `interface{}`
- **File:** `api/api.go:14,21`
- **Domain:** API
- **Description:** No type enforcement before dispatch. `binding:"required"` only checks presence.
- **Fix:** Type `Expression` as `string`; validate `Rule` type before calling `NewRuleProcessor`.

### M10. Errors silently discarded throughout evaluation chain
- **Files:** `api/api.go:55`, `evaluator.go:22`, `value.go:96-97`
- **Domain:** API
- **Description:** `result, _ := ...Evaluate(...)` discards errors everywhere. Divide-by-zero, type mismatches, invalid paths all produce zero/empty and appear as successful.
- **Fix:** Propagate errors and return HTTP 422/500.

---

## LOW Findings

### L1. `fmt.Println` in test files pollutes output
- **Files:** `value_test.go` (8), `rule_processor_test.go` (2), `evaluator_test.go` (2), `utility_test.go` (2)
- **Fix:** Remove or replace with `t.Logf`.

### L2. `isMatch` has arguments reversed
- **File:** `utility.go:24-31`
- **Description:** Compiles regex from `value` (data) instead of `target` (pattern). Currently unwired.
- **Fix:** Swap: compile from `target`, match against `value`.

### L3. Missing test files for `selector.go`, `rule.go`, `math_utility.go`
- **Description:** No dedicated unit tests for core evaluation/rule reduction logic.
- **Fix:** Add targeted unit tests for `reduceExpressions`, `determineSeedAccum`, selector dispatch.

---

## Priority Fix Order

**Wave 1 — Critical correctness bugs (Exercise 1):**
B1-B4 (wrong operators, broken recursion), B5-B6 (missing dispatch, div-by-zero), H1 (broken tests)

**Wave 2 — Crash/panic elimination (Exercise 3):**
B7-B13 (all panic vectors: type assertions, os.Exit, panic calls), H9-H10 (nil input, stack overflow)

**Wave 3 — Infrastructure modernization (Exercise 2):**
H11 (go.mod), M4 (remove fmt.Println), M6 (dead code), M7 (spelling)

**Wave 4 — API hardening:**
H5 (middleware), H6-H7 (config, graceful shutdown), M9-M10 (type safety, error propagation)

**Wave 5 — Test rehabilitation:**
M1-M3 (empty tests, error paths, file dependency), L1, L3

**Wave 6 — Remaining cleanup:**
H2-H4, H8 (wrong params, var→func), M5, M8, L2
