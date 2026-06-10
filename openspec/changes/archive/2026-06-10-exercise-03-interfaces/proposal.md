## Why

Exercise 03 is the third hands-on exercise in the Go workshop for Java developers. After learning structs and methods in exercise 02, participants now tackle Go's most distinctive feature: implicit interface satisfaction. In Java, a class must explicitly declare `implements Foo`. In Go, if a type has the right methods, it satisfies the interface — no declaration needed. This is the biggest paradigm shift in Go's type system and unlocks the "accept interfaces, return structs" idiom that defines idiomatic Go.

## What Changes

- Create exercise 03 (`exercises/03_interfaces/`) with:
  - `interfaces.go` — interface definitions, stubbed types and functions for participants to implement
  - `interfaces_test.go` — tests covering implicit satisfaction, stdlib interfaces (fmt.Stringer, error), small interfaces, interface composition, type assertions, type switches, and the nil interface trap
- Follows the same pattern as exercises 01-02: stubbed implementations with `panic("TODO: implement")`, Java-comparison comments, three-tier difficulty
- Embedding challenge section (commented-out tests) for interface composition

## Capabilities

### New Capabilities
- `exercise-03`: Exercise covering implicit interface satisfaction, implementing stdlib interfaces (fmt.Stringer), small interface design (single-method interfaces), interface composition, accepting interfaces as function parameters, type assertions and type switches, and the nil interface pitfall

### Modified Capabilities

## Impact

- New files: `exercises/03_interfaces/interfaces.go`, `exercises/03_interfaces/interfaces_test.go`
- No existing code is affected
- Follows the directory convention from `project-setup` spec
