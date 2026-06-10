## ADDED Requirements

### Requirement: Basic error returns with ParseAge
The exercise SHALL include a function `ParseAge(s string) (int, error)` that parses a string into an age integer. Participants MUST implement this function themselves, guided by instruction comments.

#### Scenario: Valid age
- **WHEN** `ParseAge("25")` is called
- **THEN** it returns `25` and `nil` error

#### Scenario: Non-numeric input
- **WHEN** `ParseAge("abc")` is called
- **THEN** it returns `0` and a non-nil error with a message containing `"invalid age"`

#### Scenario: Empty string
- **WHEN** `ParseAge("")` is called
- **THEN** it returns `0` and a non-nil error

#### Scenario: Negative age
- **WHEN** `ParseAge("-5")` is called
- **THEN** it returns `0` and a non-nil error with a message containing `"invalid age"`

### Requirement: Sentinel errors
The exercise SHALL include sentinel error variables `ErrEmpty` and `ErrOutOfRange` defined at package level using `errors.New`. A function `ValidateAge(age int) error` SHALL return the appropriate sentinel error or `nil`.

#### Scenario: ErrEmpty sentinel
- **WHEN** `ParseAge("")` is called
- **THEN** the returned error satisfies `errors.Is(err, ErrEmpty)`

#### Scenario: Valid age passes validation
- **WHEN** `ValidateAge(25)` is called
- **THEN** it returns `nil`

#### Scenario: Age too low
- **WHEN** `ValidateAge(-1)` is called
- **THEN** it returns `ErrOutOfRange`, verifiable via `errors.Is(err, ErrOutOfRange)`

#### Scenario: Age too high
- **WHEN** `ValidateAge(200)` is called
- **THEN** it returns `ErrOutOfRange`, verifiable via `errors.Is(err, ErrOutOfRange)`

#### Scenario: Sentinel identity
- **WHEN** `ErrEmpty` and `ErrOutOfRange` are compared
- **THEN** they are not equal to each other

### Requirement: Custom error type
The exercise SHALL include a custom error struct `ParseError` with fields `Input string` and `Message string`. Participants MUST define this struct and implement the `Error() string` method on it, making it satisfy the `error` interface implicitly.

#### Scenario: ParseError implements error
- **WHEN** a `ParseError{Input: "abc", Message: "not a number"}` is created
- **THEN** its `Error()` method returns `"parse error: 'abc': not a number"`

#### Scenario: ParseError is assignable to error
- **WHEN** a `*ParseError` is assigned to an `error` variable
- **THEN** the assignment compiles and the error is non-nil

### Requirement: Function returning custom error type
The exercise SHALL include a function `ParseMonth(s string) (int, error)` that returns a `*ParseError` for invalid input. This demonstrates returning a custom error type where the caller sees the `error` interface.

#### Scenario: Valid month
- **WHEN** `ParseMonth("6")` is called
- **THEN** it returns `6` and `nil` error

#### Scenario: Non-numeric month
- **WHEN** `ParseMonth("june")` is called
- **THEN** it returns `0` and a `*ParseError` with `Input` equal to `"june"`

#### Scenario: Month out of range
- **WHEN** `ParseMonth("13")` is called
- **THEN** it returns `0` and a `*ParseError` with `Input` equal to `"13"` and a message indicating out of range

### Requirement: Error wrapping with fmt.Errorf %w
The exercise SHALL include a function `ParseUserAge(s string) (int, error)` that calls `ParseAge` and wraps any returned error using `fmt.Errorf("parsing user age: %w", err)`. Tests SHALL verify that `errors.Is` can find the original sentinel through the wrap chain.

#### Scenario: Wrapping preserves sentinel
- **WHEN** `ParseUserAge("")` is called
- **THEN** the returned error satisfies `errors.Is(err, ErrEmpty)`
- **THEN** the error message contains `"parsing user age:"`

#### Scenario: Wrapping valid input
- **WHEN** `ParseUserAge("30")` is called
- **THEN** it returns `30` and `nil`

#### Scenario: Wrapped error string contains context
- **WHEN** `ParseUserAge("abc")` is called
- **THEN** the error message starts with `"parsing user age:"` and contains the original error text

### Requirement: errors.As challenge
The exercise SHALL include a commented-out challenge section where participants use `errors.As` to extract a `*ParseError` from a wrapped error chain. A function `ParseBirthMonth(s string) (int, error)` SHALL call `ParseMonth` and wrap the error.

#### Scenario: errors.As extracts custom type
- **WHEN** `ParseBirthMonth("june")` is called and the error is unwrapped with `errors.As`
- **THEN** the extracted `*ParseError` has `Input` equal to `"june"`

#### Scenario: errors.As returns false for non-matching type
- **WHEN** an error created with `errors.New("simple")` is checked with `errors.As` for `*ParseError`
- **THEN** `errors.As` returns `false`

### Requirement: Progressive unlock test structure
The exercise SHALL use the same progressive unlock pattern as exercise 03: all tests commented out with instructions to uncomment step by step. Imports SHALL be commented out when not yet needed.

#### Scenario: Clean compilation
- **WHEN** the exercise files are compiled without any modifications
- **THEN** they compile without errors (commented-out tests do not affect compilation)

### Requirement: Participants define all types and functions
The exercise `.go` file SHALL contain only instruction comments and suppressed imports. Participants MUST define all sentinel errors, custom error types, and functions themselves. No structs, interfaces, functions, or error variables SHALL be pre-defined.

#### Scenario: Empty starting point
- **WHEN** a participant opens `errors.go`
- **THEN** they see instruction comments for each step but no pre-defined types, functions, or error variables (only suppressed imports)
