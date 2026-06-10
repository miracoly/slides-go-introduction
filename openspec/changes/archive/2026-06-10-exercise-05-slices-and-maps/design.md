## Context

Fifth exercise in the Go workshop for experienced Java developers. This is the second exercise of Part 2. Participants have completed types/functions, structs/methods, interfaces, and error handling. They now need fluency with Go's core collection types.

Java developers live in `ArrayList<T>` and `HashMap<K,V>` — they know collections well. But Go's slices and maps have fundamental behavioral differences: slices have length AND capacity (no equivalent in Java), `append` may or may not allocate, sub-slicing shares underlying memory, maps use comma-ok instead of `containsKey`, and map iteration order is randomized.

## Goals / Non-Goals

**Goals:**
- Teach slice fundamentals: literals, append, len/cap, sub-slicing
- Show the nil vs empty slice distinction (both are valid, but behave differently with JSON later)
- Build fluency with `range` loops over slices and maps
- Teach map fundamentals: literals, insertion, comma-ok lookup, delete
- Demonstrate that map iteration order is not guaranteed
- Combine slices and maps in practical functions (word counting, grouping)

**Non-Goals:**
- `make` with pre-allocated capacity (optimization detail, not fundamentals)
- Multi-dimensional slices
- `sync.Map` or concurrent map access (belongs in concurrency exercise)
- Sorting (sort package is useful but not core to understanding collections)
- Generics with collections

## Decisions

### Step structure: slices first, then maps, then combine

Slices are more fundamental — they appear everywhere, including as map values. Teaching slices first means maps can use `[]string` as values naturally.

Steps:
1. **Slice basics** — create slices, use append, read len
2. **Slice operations** — sub-slicing, range loops, filtering
3. **Map basics** — create maps, insert, lookup with comma-ok, delete
4. **Practical map functions** — word counting with `map[string]int`
5. **Challenge** — combine slices and maps: group strings by a key function

### Domain: string manipulation functions

Use functions that operate on `[]string` and `map[string]int` / `map[string][]string`. These are the most common collection patterns in real Go code. Word counting, filtering, grouping — practical and universally understood.

**Why over numeric examples:** String slices and word counting feel like real work. Numeric examples (sum, max, etc.) are too trivial and don't exercise the full API surface (comma-ok, range with index+value, etc.).

### Step 1: Slice basics with Contains and Filter

Participants write `Contains(haystack []string, needle string) bool` and `FilterByPrefix(words []string, prefix string) []string`. These force using `range` and `append` — the two most common slice operations.

### Step 2: Map basics with WordCount

`WordCount(words []string) map[string]int` is a classic exercise that teaches map creation, zero-value behavior (incrementing a key that doesn't exist yet works because `int` zero value is `0`), and range over slices.

### Step 3: Comma-ok and map operations

`UniqueStrings(words []string) []string` uses a `map[string]bool` as a set — demonstrating the comma-ok idiom for checking existence. `Keys(m map[string]int) []string` iterates a map to extract keys.

### Step 4: Nil vs empty slice distinction

A test-only step (no function to write). Participants uncomment a test that demonstrates `var s []string` (nil) vs `s := []string{}` (empty, non-nil). Both have `len == 0`, but `s == nil` differs. This parallels the nil interface trap from exercise 03.

### Challenge: GroupBy function

`GroupBy(words []string, keyFn func(string) string) map[string][]string` combines everything: iterating slices, using maps with slice values, append to map values. The `func(string) string` parameter introduces function values naturally.

## Risks / Trade-offs

- **Too much ground to cover** → Steps 1-3 are the core (contains, filter, word count, unique). Step 4 is observational. Challenge is stretch. Stopping after step 3 gives solid working knowledge.
- **Map iteration order makes tests flaky** → Tests that check map contents must use element-by-element comparison, not string comparison. `Keys` function test sorts the result before comparing.
- **`make` omission** → Participants may wonder about pre-allocation. A brief comment mentions it exists but frames it as an optimization, not a fundamental.
