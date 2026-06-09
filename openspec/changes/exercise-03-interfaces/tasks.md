## 1. Exercise 03 — Interfaces

- [x] 1.1 Create `exercises/03_interfaces/interfaces.go` with: `Shape` interface (`Area`, `Perimeter`), `Circle` and `Rectangle` structs with stubbed methods, `String()` on Circle (fmt.Stringer), `Describer` interface, stubbed `Describe()` on both types, `TotalArea([]Shape)` function, `DescribeShape(Shape)` function — all bodies using `panic("TODO: implement")`. Include commented-out `ShapeDescriber` composed interface as challenge section.
- [x] 1.2 Create `exercises/03_interfaces/interfaces_test.go` with tests ordered warm-up (implicit satisfaction, Stringer) → core (TotalArea, Describer, type assertion, type switch) → think-about-it (nil interface trap). Include commented-out tests for interface composition challenge. Java-comparison comments throughout.

## 2. Verification

- [x] 2.1 Run `task test -- 03` and verify all tests fail with "TODO: implement" panics (not compile errors)
- [x] 2.2 Implement all functions temporarily, verify all tests pass, then revert to `panic("TODO: implement")`
