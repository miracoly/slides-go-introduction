## Why

Exercise 04 is the first exercise of Part 2 of the Go workshop. After learning Go's type system (types, structs, interfaces), participants now tackle error handling — the most visible cultural difference between Go and Java. Java uses exceptions with try/catch/finally and an invisible control flow. Go treats errors as values returned from functions. This exercise bridges that paradigm shift and gives participants the `if err != nil` muscle memory they'll use constantly.

## What Changes

- Create exercise 04 (`exercises/04_errors/`) with:
  - `errors.go` — step-by-step instruction comments for participants to implement error handling patterns from scratch
  - `errors_test.go` — tests covering error creation, the `error` interface, custom error types, error wrapping with `%w`, `errors.Is`/`errors.As`, and sentinel errors
- Follows the progressive unlock pattern from exercise 03: all tests commented out, participants uncomment step by step
- Participants define everything themselves: custom error types, sentinel errors, functions that return errors, wrapping logic
- Domain: a string parser that converts structured strings into data (natural failure modes, no I/O dependencies)
- Challenge section for `errors.As` with custom error types

## Capabilities

### New Capabilities
- `exercise-04`: Exercise covering error-as-values philosophy, the `error` interface, creating errors with `errors.New` and `fmt.Errorf`, sentinel errors, custom error types implementing `error`, error wrapping with `%w`, `errors.Is` for sentinel comparison, `errors.As` for custom type extraction, and the contrast with Java's exception model

### Modified Capabilities

## Impact

- New files: `exercises/04_errors/errors.go`, `exercises/04_errors/errors_test.go`
- No existing code is affected
- Follows the directory convention from `project-setup` spec
