## Context

This is the second exercise in a Go workshop for experienced Java developers. Exercise 01 covered types, variables, and functions. Participants now understand `:=`, zero values, multiple returns, pointers, and `iota` constants.

Exercise 02 introduces structs and methods — Go's replacement for Java classes. The key insight for Java devs: Go separates data (structs) from behavior (methods) and uses composition (embedding) instead of inheritance. This is the biggest conceptual shift in the workshop.

The exercise follows the established conventions: package under `exercises/02_structs_and_methods/`, stubbed `.go` file with `panic("TODO: implement")`, test file with Java-comparison comments, three-tier difficulty.

## Goals / Non-Goals

**Goals:**
- Teach struct definition, initialization, and field access
- Demonstrate factory functions as Go's replacement for constructors
- Clarify the critical difference between value and pointer receivers
- Show exported vs unexported visibility via capitalization
- Introduce struct embedding as Go's composition mechanism (replaces inheritance)

**Non-Goals:**
- Interfaces (exercise 03)
- Struct tags (JSON/DB tags — skipped for this workshop)
- Anonymous structs
- Generics on structs

## Decisions

### Domain model: BankAccount and Address/Person

Use a `BankAccount` struct for the core exercises (creation, methods, pointer receivers) and `Address`/`Person` structs for embedding. BankAccount is intuitive — deposit/withdraw naturally require pointer receivers (mutating state), and balance is a natural candidate for unexported fields with getter methods.

**Why over abstract examples:** Concrete domain models make the Java comparison immediate. Java devs think in terms of "I'd make this a class with private fields and public methods" — then they see the Go equivalent.

### Visibility via a separate internal detail

To demonstrate exported vs unexported, the `BankAccount` struct will have an unexported field (`balance`) with an exported getter method (`Balance()`). The tests will verify that the getter returns the correct value after mutations. We won't test unexported access from outside the package (that's a compile-time guarantee), but the comments will explain the parallel to Java's `private`.

### Value vs pointer receiver as an explicit "aha" test

Include a test that calls a value-receiver method that tries to mutate state, then asserts the original is unchanged. This is the single most important lesson for Java devs — in Java, method calls on objects always see the same reference. In Go, a value receiver gets a copy.

## Risks / Trade-offs

- **Embedding may feel like inheritance to Java devs** → Test comments explicitly call out "this is NOT inheritance — there is no polymorphism here." Exercise 03 (interfaces) will show what Go uses instead.
- **Unexported fields can't be tested from outside the package** → We're in the same package (`package structs_and_methods`), so tests can access unexported fields for setup. But the comments explain that external packages cannot.
