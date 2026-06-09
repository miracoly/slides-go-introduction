## ADDED Requirements

### Requirement: Short variable declaration
The exercise SHALL include a function that demonstrates Go's `:=` short variable declaration syntax, contrasting with Java's explicit type declarations.

#### Scenario: Declare and return variables
- **WHEN** `DeclareVariables()` is called
- **THEN** it returns an `int`, a `string`, and a `bool` declared using `:=`
- **THEN** the returned values are `42`, `"hello"`, and `true`

### Requirement: Zero values
The exercise SHALL include a function demonstrating that Go initializes all variables to their zero value — unlike Java where uninitialized local variables cause a compile error.

#### Scenario: Return zero values for all basic types
- **WHEN** `ZeroValues()` is called
- **THEN** it returns the zero values: `0` (int), `""` (string), `false` (bool), and `nil` (pointer to any struct)

### Requirement: Multiple return values
The exercise SHALL include a function demonstrating Go's native multiple return values — a concept that in Java requires a wrapper class or `Pair`.

#### Scenario: Successful division
- **WHEN** `Divide(10, 3)` is called
- **THEN** it returns `3` (integer division result) and `nil` (no error)

#### Scenario: Division by zero
- **WHEN** `Divide(10, 0)` is called
- **THEN** it returns `0` and a non-nil error with message `"division by zero"`

### Requirement: Named return values
The exercise SHALL include a function using named return values, showing how Go can name its return parameters and use a bare `return`.

#### Scenario: Swap two integers
- **WHEN** `Swap(1, 2)` is called
- **THEN** it returns `2, 1` using named return values

### Requirement: Constants and iota
The exercise SHALL include typed constants using `iota` for sequential values, analogous to Java's `enum` ordinals.

#### Scenario: Weekday constants
- **WHEN** the package is loaded
- **THEN** constants `Monday` through `Friday` exist as type `Weekday` (an `int` type)
- **THEN** `Monday` equals `1`, `Tuesday` equals `2`, through `Friday` equals `5`

#### Scenario: Weekday string representation
- **WHEN** `Monday.String()` is called
- **THEN** it returns `"Monday"`

### Requirement: Explicit type conversion
The exercise SHALL include a function requiring explicit numeric type conversion, highlighting that Go has no implicit widening — unlike Java which silently widens `int` to `long`.

#### Scenario: Convert between numeric types
- **WHEN** `ConvertTemperature(100)` is called with an `int` parameter (Celsius)
- **THEN** it returns `212.0` as a `float64` (Fahrenheit), performing explicit int-to-float conversion

### Requirement: Pointer basics
The exercise SHALL include functions demonstrating pointer usage, since Go has explicit pointers whereas Java hides references for objects.

#### Scenario: Increment via pointer
- **WHEN** `Increment(p)` is called with a `*int` pointer to a variable holding `5`
- **THEN** the pointed-to variable now holds `6`

#### Scenario: Create and return a pointer
- **WHEN** `NewInt(42)` is called
- **THEN** it returns a `*int` pointing to a value of `42`
