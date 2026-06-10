## Context

Third exercise in the Go workshop for experienced Java developers. Participants have completed exercises on types/functions and structs/methods. They understand structs, factory functions, pointer vs value receivers, and embedding. Now they need to understand Go's interface system — the mechanism that makes Go's type system truly powerful.

For Java devs, this is the biggest mental shift: interfaces are satisfied implicitly. There's no `implements` keyword. If your type has the methods, it satisfies the interface. This enables loose coupling, small interfaces, and the "accept interfaces, return structs" pattern.

## Goals / Non-Goals

**Goals:**
- Demonstrate implicit interface satisfaction (no `implements` keyword)
- Show the stdlib's small interface pattern (fmt.Stringer, error, io.Reader)
- Teach "accept interfaces, return structs" — the core Go design idiom
- Cover type assertions and type switches for extracting concrete types
- Expose the nil interface trap (an interface holding a nil pointer is not nil)

**Non-Goals:**
- io.Reader/Writer implementation (too detailed for this exercise, better in a dedicated I/O exercise)
- Generic interfaces (out of scope for this workshop)
- Empty interface `any` in depth (mention it, don't exercise it)
- Mocking/testing patterns with interfaces (would be great for an advanced workshop)

## Decisions

### Domain model: Shape interface with Circle and Rectangle

Use the classic `Shape` interface with `Area() float64` and `Perimeter() float64`. Java devs know this pattern from textbooks, making the comparison immediate. `Circle` and `Rectangle` implement it implicitly.

**Why over more creative examples:** The Shape example is universally understood. The teaching point isn't the domain — it's the *mechanism*. A familiar domain lets participants focus on the Go-specific behavior.

### Stringer interface as the "stdlib gateway"

Include implementing `fmt.Stringer` (the `String() string` method) on one of the shape types. This connects to exercise 02's `Summary()` pattern and shows that stdlib interfaces work the same way as user-defined ones.

### Describer interface for "accept interfaces" pattern

Define a `Describer` interface with a single `Describe() string` method. Then write a function `PrintDescription(d Describer) string` that accepts the interface. Multiple concrete types implement it, demonstrating polymorphism without inheritance.

### Type assertions and type switch as the "think about it" tier

Type assertions (`v, ok := i.(ConcreteType)`) and type switches are more advanced but critical for real-world Go. Place these in the third tier. Java devs will recognize the parallel to `instanceof` + cast.

### Nil interface trap as the final eye-opener

The nil interface gotcha is subtle and important. An interface value that holds a nil concrete pointer is itself *not nil*. This surprises everyone. Place it as the last test with thorough explanation.

### Challenge section for interface composition

Following the exercise 02 pattern, interface composition (`ReadWriter` = `Reader` + `Writer` style) is presented as a commented-out challenge. Participants define composed interfaces and uncomment tests.

## Risks / Trade-offs

- **Shape example feels academic** → Mitigated by connecting to Stringer (real stdlib) and the Describer pattern (real design idiom). The Shape is just the warm-up.
- **Nil interface trap is confusing without context** → Provide extensive comments explaining *why* it works this way (interface = type + value, both must be nil).
