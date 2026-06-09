## Why

Exercise 02 is the second hands-on exercise in the Go workshop for Java developers. After learning types and functions in exercise 01, participants now need to understand how Go models data and behavior without classes. This is the biggest paradigm shift for Java devs: no classes, no inheritance, no constructors — just structs, methods, composition via embedding, and visibility through capitalization.

## What Changes

- Create exercise 02 (`exercises/02_structs_and_methods/`) with:
  - `structs.go` — stubbed types, factory functions, and methods for participants to implement
  - `structs_test.go` — tests covering struct creation, methods, pointer vs value receivers, visibility, and embedding
- Tests follow the same three-tier structure (warm-up → core → think-about-it) with Java-comparison comments

## Capabilities

### New Capabilities
- `exercise-02`: Exercise covering struct definition and literals, factory functions (Go's replacement for constructors), methods with explicit receivers, pointer vs value receivers, exported/unexported fields and methods (visibility via capitalization), and struct embedding as composition (Go's replacement for inheritance)

### Modified Capabilities

## Impact

- New files: `exercises/02_structs_and_methods/structs.go`, `exercises/02_structs_and_methods/structs_test.go`
- No existing code is affected
- Follows the directory convention established by the `project-setup` spec
