## Why

Exercise 05 is the second exercise of Part 2 of the Go workshop. After learning error handling, participants now tackle Go's two most important collection types: slices and maps. Java developers rely heavily on `ArrayList` and `HashMap` — Go's equivalents behave differently in critical ways. Slices are backed by arrays with length/capacity semantics, and maps have no ordering guarantees and no `containsKey` — you use the comma-ok idiom instead. These are the data structures participants will use in virtually every Go program.

## What Changes

- Create exercise 05 (`exercises/05_collections/`) with:
  - `collections.go` — step-by-step instruction comments for participants to implement slice and map operations from scratch
  - `collections_test.go` — tests covering slice creation, append, slicing, len/cap, range loops, map creation, comma-ok lookups, delete, iteration, and combining slices with maps
- Follows the progressive unlock pattern: all tests commented out, participants uncomment step by step
- Participants define all functions themselves
- Domain: practical data manipulation — word counting, grouping, filtering — exercises that feel like real work
- Challenge section for combining slices and maps in a grouping function

## Capabilities

### New Capabilities
- `exercise-05`: Exercise covering slice literals, append, len/cap, sub-slicing, nil vs empty slices, range loops over slices, map literals, map insertion and lookup with comma-ok, delete, range over maps, and combining slices with maps in practical functions

### Modified Capabilities

## Impact

- New files: `exercises/05_collections/collections.go`, `exercises/05_collections/collections_test.go`
- No existing code is affected
- Follows the directory convention from `project-setup` spec
