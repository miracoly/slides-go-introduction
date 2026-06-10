## ADDED Requirements

### Requirement: Slice contains check
The exercise SHALL include a function `Contains(haystack []string, needle string) bool` that checks whether a string exists in a slice. Participants MUST implement this using a `range` loop.

#### Scenario: Element found
- **WHEN** `Contains([]string{"go", "java", "rust"}, "java")` is called
- **THEN** it returns `true`

#### Scenario: Element not found
- **WHEN** `Contains([]string{"go", "java", "rust"}, "python")` is called
- **THEN** it returns `false`

#### Scenario: Empty slice
- **WHEN** `Contains([]string{}, "anything")` is called
- **THEN** it returns `false`

### Requirement: Filter slice by prefix
The exercise SHALL include a function `FilterByPrefix(words []string, prefix string) []string` that returns a new slice containing only strings that start with the given prefix. Participants MUST use `append` to build the result.

#### Scenario: Some matches
- **WHEN** `FilterByPrefix([]string{"go", "goroutine", "java", "golang"}, "go")` is called
- **THEN** it returns `[]string{"go", "goroutine", "golang"}`

#### Scenario: No matches
- **WHEN** `FilterByPrefix([]string{"java", "rust", "python"}, "go")` is called
- **THEN** it returns an empty (possibly nil) slice with length 0

#### Scenario: All match
- **WHEN** `FilterByPrefix([]string{"go", "goroutine"}, "go")` is called
- **THEN** it returns `[]string{"go", "goroutine"}`

### Requirement: Word count with map
The exercise SHALL include a function `WordCount(words []string) map[string]int` that counts occurrences of each word. This demonstrates map creation and zero-value increment.

#### Scenario: Mixed words
- **WHEN** `WordCount([]string{"go", "java", "go", "rust", "go"})` is called
- **THEN** it returns a map where `"go"` maps to `3`, `"java"` to `1`, `"rust"` to `1`

#### Scenario: Empty slice
- **WHEN** `WordCount([]string{})` is called
- **THEN** it returns an empty map (length 0)

#### Scenario: All unique
- **WHEN** `WordCount([]string{"a", "b", "c"})` is called
- **THEN** each key maps to `1`

### Requirement: Unique strings using map as set
The exercise SHALL include a function `UniqueStrings(words []string) []string` that returns a slice of unique strings preserving first-occurrence order. This demonstrates using a map as a set with the comma-ok idiom.

#### Scenario: Duplicates removed
- **WHEN** `UniqueStrings([]string{"go", "java", "go", "rust", "java"})` is called
- **THEN** it returns `[]string{"go", "java", "rust"}` (first-occurrence order)

#### Scenario: Already unique
- **WHEN** `UniqueStrings([]string{"a", "b", "c"})` is called
- **THEN** it returns `[]string{"a", "b", "c"}`

#### Scenario: Empty input
- **WHEN** `UniqueStrings([]string{})` is called
- **THEN** it returns an empty (possibly nil) slice with length 0

### Requirement: Extract keys from map
The exercise SHALL include a function `Keys(m map[string]int) []string` that returns all keys from a map. Tests MUST sort the result before comparison since map iteration order is not guaranteed.

#### Scenario: Multiple keys
- **WHEN** `Keys(map[string]int{"go": 3, "java": 1, "rust": 1})` is called
- **THEN** it returns a slice containing `"go"`, `"java"`, `"rust"` (in any order)

#### Scenario: Empty map
- **WHEN** `Keys(map[string]int{})` is called
- **THEN** it returns an empty (possibly nil) slice with length 0

### Requirement: Nil vs empty slice distinction
The exercise SHALL include a test-only step (no function to write) demonstrating that a nil slice and an empty slice both have `len == 0` but differ in nil-ness. This parallels the nil interface trap from exercise 03.

#### Scenario: Nil slice behavior
- **WHEN** a slice is declared with `var s []string` (nil slice)
- **THEN** `len(s) == 0` is true AND `s == nil` is true

#### Scenario: Empty slice behavior
- **WHEN** a slice is created with `s := []string{}` (empty, non-nil)
- **THEN** `len(s) == 0` is true AND `s == nil` is false

#### Scenario: Append works on nil slice
- **WHEN** `append` is called on a nil slice
- **THEN** it works correctly (nil slice is a valid, empty slice for all operations)

### Requirement: GroupBy challenge
The exercise SHALL include a commented-out challenge section where participants implement `GroupBy(words []string, keyFn func(string) string) map[string][]string` that groups words by the result of applying a key function. This combines slices as map values with function values as parameters.

#### Scenario: Group by first letter
- **WHEN** `GroupBy([]string{"go", "goroutine", "java", "rust"}, func(s string) string { return string(s[0]) })` is called
- **THEN** it returns a map where `"g"` maps to `[]string{"go", "goroutine"}`, `"j"` to `[]string{"java"}`, `"r"` to `[]string{"rust"}`

#### Scenario: Group by length
- **WHEN** words are grouped by a function that returns the string length as a string
- **THEN** words of the same length are grouped together

### Requirement: Progressive unlock test structure
The exercise SHALL use the same progressive unlock pattern as exercises 03 and 04: all tests commented out with instructions to uncomment step by step. Imports SHALL be commented out when not yet needed.

#### Scenario: Clean compilation
- **WHEN** the exercise files are compiled without any modifications
- **THEN** they compile without errors

### Requirement: Participants define all functions
The exercise `.go` file SHALL contain only instruction comments and suppressed imports. Participants MUST define all functions themselves. No functions SHALL be pre-defined.

#### Scenario: Empty starting point
- **WHEN** a participant opens `collections.go`
- **THEN** they see instruction comments for each step but no pre-defined functions (only suppressed imports)
