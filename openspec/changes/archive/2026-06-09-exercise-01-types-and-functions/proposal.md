## Why

This is the first hands-on exercise of a Go workshop for experienced Java developers. It introduces Go's type system, variable declarations, and functions — the most fundamental building blocks. Java devs already know these concepts, so the exercise focuses on what's *different*: short variable declarations, zero values, multiple return values, explicit type conversions, and pointers. Getting these basics right sets the foundation for all subsequent exercises.

## What Changes

- Add project-level `go.mod` to establish the Go module
- Add a `Taskfile.yml` with a task to run exercise tests
- Create exercise 01 (`exercises/01_types_and_functions/`) with:
  - `types.go` — stubbed function signatures for participants to implement
  - `types_test.go` — tests that verify each implementation, with Java-comparison comments
- Establish the directory structure convention used by all future exercises

## Capabilities

### New Capabilities
- `project-setup`: Go module initialization, Taskfile with test runner task, directory structure convention for exercises
- `exercise-01`: Exercise covering variable declarations (`:=`, `var`, zero values), multiple return values, constants with `iota`, explicit type conversions, and pointer basics

### Modified Capabilities

## Impact

- New files: `go.mod`, `Taskfile.yml`, `exercises/01_types_and_functions/types.go`, `exercises/01_types_and_functions/types_test.go`
- No existing code is affected (greenfield)
- Participants need Go installed (provided via Nix flake) and optionally `go-task`
