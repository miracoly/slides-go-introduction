## Context

This is a Go workshop repository for experienced Java developers. The repo currently contains only a Nix flake providing Go tooling. We need to establish the project structure and create the first exercise covering Go's type system and functions — focusing on what surprises Java developers.

Participants will clone this repo, enter the Nix shell, and run tests with `go-task`. Each exercise lives in its own package under `exercises/`. The stubbed `.go` file has empty function bodies; the `_test.go` file contains tests with Java-comparison comments. Participants fill in implementations until all tests pass.

## Goals / Non-Goals

**Goals:**
- Establish a repeatable directory structure convention for all future exercises
- Provide a single `task test` command to run all exercise tests
- Create exercise 01 with ~7 test functions covering types, variables, functions, constants, conversions, and pointers
- Each test name and comment explicitly references the Java equivalent to anchor learning

**Non-Goals:**
- No `main.go` or runnable binary
- No solution files (solutions will live on a separate branch later)
- No CI/CD setup
- No coverage of generics, interfaces, or error handling (later exercises)

## Decisions

### Single Go module at project root

All exercises share one `go.mod` at the project root. Each exercise is a package under `exercises/01_types_and_functions/`, `exercises/02_structs_and_methods/`, etc.

**Why over separate modules per exercise:** Simpler for participants — one `go.mod`, one `go test ./...`. Separate modules would require `cd`-ing into each exercise and managing independent module files.

### go-task as test runner

A `Taskfile.yml` at the root provides `task test` (run all exercises) and `task test -- 01` (run a specific exercise by number prefix). go-task is already in the Nix flake.

**Why over Makefile:** go-task is already a dependency, its YAML format is more readable, and it supports cross-platform use.

### Stubbed functions with TODO comments

Each exercise's `.go` file provides the function signatures with return type info, but the body is a single `panic("TODO: implement")`. This way the code compiles (tests can run and fail) rather than requiring participants to also figure out signatures.

**Why panic over zero-value stubs:** A panic clearly signals "not implemented" in test output. Zero-value returns would silently pass some tests or give confusing failures.

### Test names reference Java equivalents

Test functions are named descriptively (e.g., `TestMultipleReturnValues`) and include a comment block that briefly explains the Java equivalent and what's different in Go. This anchors learning for the target audience.

### Three-tier difficulty within the exercise

Tests are ordered by difficulty:
1. **Warm-up** — trivial for any developer (variable declaration, basic function)
2. **Core** — the main concepts (multiple returns, zero values, constants with iota)
3. **Think about it** — subtle differences (pointer mutation, type conversion gotchas)

## Risks / Trade-offs

- **`panic("TODO")` crashes test runs early** → Each test calls functions independently, so one panic doesn't block others. Go's test runner reports each test separately.
- **Single package per exercise limits visibility testing** → Acceptable for exercise 01. Exercise 02 (structs) will introduce exported/unexported via a sub-package if needed.
- **Java comparisons in comments may feel patronizing** → Keep them brief and factual, not editorial. Frame as "In Java: X. In Go: Y." not "Go is better because...".
