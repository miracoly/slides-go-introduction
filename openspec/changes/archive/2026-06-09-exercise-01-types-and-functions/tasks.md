## 1. Project Setup

- [x] 1.1 Create `go.mod` at project root with `module go-introduction` and appropriate Go version
- [x] 1.2 Create `Taskfile.yml` with `test` task that runs `go test ./exercises/...` and supports filtering by exercise number via CLI args (e.g. `task test -- 01`)

## 2. Exercise 01 — Types & Functions

- [x] 2.1 Create `exercises/01_types_and_functions/types.go` with package declaration and all stubbed functions: `DeclareVariables`, `ZeroValues`, `Divide`, `Swap`, `ConvertTemperature`, `Increment`, `NewInt`, plus `Weekday` type with iota constants and `String()` method — all bodies using `panic("TODO: implement")`
- [x] 2.2 Create `exercises/01_types_and_functions/types_test.go` with tests for all functions, ordered warm-up → core → subtle, with Java-comparison comments on each test

## 3. Verification

- [x] 3.1 Run `task test` and verify all tests fail with "TODO: implement" panics (not compile errors)
- [x] 3.2 Implement all functions (temporary, to verify tests are correct), run `task test` and verify all tests pass, then revert implementations back to `panic("TODO: implement")`
