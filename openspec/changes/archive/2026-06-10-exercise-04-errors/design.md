## Context

Fourth exercise in the Go workshop for experienced Java developers. This is the first exercise of Part 2. Participants have completed exercises on types/functions, structs/methods, and interfaces. They understand Go's type system well. Now they need to learn Go's error handling philosophy — the most visible day-to-day difference from Java.

Java developers are deeply trained on exceptions: try/catch/finally, checked vs unchecked, throw/throws. Go rejects this entire model in favor of returning errors as values. This is not just syntax — it's a different philosophy about where error handling belongs in code flow.

## Goals / Non-Goals

**Goals:**
- Show that `error` is just an interface (`Error() string`) — connects to exercise 03
- Build the `if err != nil` reflex through repetition
- Teach creating errors with `errors.New` and `fmt.Errorf`
- Introduce sentinel errors (package-level error variables) as the Go equivalent of specific exception types
- Demonstrate custom error types (structs implementing `error`) for carrying context
- Teach error wrapping with `fmt.Errorf("...: %w", err)` and unwrapping with `errors.Is`
- Show `errors.As` for extracting custom error types from a chain

**Non-Goals:**
- `panic`/`recover` in depth (mention in comments, don't exercise — it's rare in production Go)
- Error handling patterns at scale (middleware, error groups)
- Third-party error libraries (pkg/errors, etc.)
- Logging or observability of errors

## Decisions

### Domain: string parsing functions

Use functions that parse strings into structured data — e.g., parsing "3:45" into hours and minutes, parsing "25" into a validated age, parsing "100USD" into an amount and currency. Parsing naturally produces errors (bad format, out of range, empty input) without needing I/O or external dependencies.

**Why over BankAccount continuation:** Fresh domain avoids repetition and gives more natural error variety. Bank operations have limited failure modes (just insufficient funds). Parsing has format errors, range errors, missing data — much richer for teaching.

**Why over calculator:** A calculator's main error is division by zero, which exercise 01 already covered. Parsing gives us more distinct error categories.

### Progressive unlock: all tests commented out, participants build everything

Follow the exercise 03 pattern. The `.go` file contains only instruction comments and suppressed imports. All tests are commented out. Participants uncomment one step at a time.

Steps build on each other:
1. **Basic errors** — write functions that return `(value, error)` using `errors.New` and `fmt.Errorf`
2. **Sentinel errors** — define package-level error variables, check with `errors.Is`
3. **Custom error types** — define a struct implementing the `error` interface
4. **Error wrapping** — wrap errors with `%w`, unwrap with `errors.Is` and `errors.As`
5. **Challenge** — `errors.As` to extract a custom error type from a wrapped chain

### Step 1: Basic error returns with parsing functions

Participants write two functions:
- `ParseAge(s string)` → parses a string to an int, returns error for non-numeric or out-of-range input
- `ParseHours(s string)` → parses "HH:MM" format, returns error for bad format

This builds the `(value, error)` return pattern and `if err != nil` muscle memory. Uses `strconv.Atoi` and `strings.Split` — both stdlib, no new dependencies.

### Step 2: Sentinel errors

Participants define sentinel errors like `var ErrEmpty = errors.New("empty input")` and `var ErrOutOfRange = errors.New("out of range")`. A `ValidateAge(age int)` function returns the appropriate sentinel. Tests check with `errors.Is`.

**Why sentinel errors early:** They're the Go equivalent of catching specific exception types in Java (`catch (IllegalArgumentException e)`). Java devs need to see this mapping clearly.

### Step 3: Custom error type

Participants define a `ParseError` struct with fields for what went wrong (e.g., the input string, a message). They implement `Error() string` on it. This ties back to exercise 03's lesson that interfaces are satisfied implicitly — `error` is just an interface.

### Step 4: Wrapping errors

Participants write a higher-level function that calls a lower-level one and wraps the error with `fmt.Errorf("context: %w", err)`. Tests use `errors.Is` to check through the wrap chain.

**Why keep `errors.As` for challenge:** `errors.Is` is simpler (compares values). `errors.As` requires understanding type assertion through an error chain — more advanced, better as a challenge.

### Challenge: errors.As

Participants use `errors.As` to extract a `*ParseError` from a wrapped error chain. This combines wrapping + custom types and completes the error handling toolkit.

## Risks / Trade-offs

- **Parsing domain feels disconnected** → Mitigated by keeping functions simple and focused on the error patterns, not the parsing logic itself. `strconv.Atoi` does the heavy lifting.
- **Too many steps for 50 minutes** → Steps 1-3 are the core; step 4 and challenge are stretch goals. The exercise is designed so stopping after step 3 still gives a solid understanding.
- **`errors.Is` vs `==` confusion** → Comments explain that `errors.Is` checks through wrapped chains while `==` only checks the top level. The wrapping step makes this concrete.
