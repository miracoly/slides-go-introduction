package collections

import (
	// Uncomment "testing" when you uncomment your first test.
	// Add "sort" when you reach Step 3, and "fmt" when you reach the challenge.
	// "fmt"
	// "sort"
	// "testing"
)

// ============================================================================
// Step 1: Slice basics — Contains and FilterByPrefix
//
// Implement Contains and FilterByPrefix in collections.go,
// then uncomment the tests below.
// ============================================================================

// // In Java: list.contains("java") — built into the collection.
// // In Go: there's no built-in. You write it with a range loop.
// // This is Go's philosophy: simple building blocks over large APIs.
// func TestContainsFound(t *testing.T) {
// 	got := Contains([]string{"go", "java", "rust"}, "java")
// 	if !got {
// 		t.Error("Contains should return true when element exists")
// 	}
// }

// func TestContainsNotFound(t *testing.T) {
// 	got := Contains([]string{"go", "java", "rust"}, "python")
// 	if got {
// 		t.Error("Contains should return false when element is missing")
// 	}
// }

// func TestContainsEmptySlice(t *testing.T) {
// 	got := Contains([]string{}, "anything")
// 	if got {
// 		t.Error("Contains should return false for empty slice")
// 	}
// }

// // In Java: list.stream().filter(s -> s.startsWith("go")).toList()
// // In Go: range + append. No streams, no lambdas needed.
// func TestFilterByPrefixSomeMatch(t *testing.T) {
// 	got := FilterByPrefix([]string{"go", "goroutine", "java", "golang"}, "go")
// 	want := []string{"go", "goroutine", "golang"}
//
// 	if len(got) != len(want) {
// 		t.Fatalf("FilterByPrefix: got %d elements, want %d", len(got), len(want))
// 	}
// 	for i := range want {
// 		if got[i] != want[i] {
// 			t.Errorf("FilterByPrefix[%d]: got %q, want %q", i, got[i], want[i])
// 		}
// 	}
// }

// func TestFilterByPrefixNoMatch(t *testing.T) {
// 	got := FilterByPrefix([]string{"java", "rust", "python"}, "go")
// 	if len(got) != 0 {
// 		t.Errorf("FilterByPrefix: got %d elements, want 0", len(got))
// 	}
// }

// func TestFilterByPrefixAllMatch(t *testing.T) {
// 	got := FilterByPrefix([]string{"go", "goroutine"}, "go")
// 	if len(got) != 2 {
// 		t.Errorf("FilterByPrefix: got %d elements, want 2", len(got))
// 	}
// }

// ============================================================================
// Step 2: Map basics — WordCount
//
// After Step 1 passes, implement WordCount in collections.go,
// then uncomment the tests below.
// ============================================================================

// // In Java: map.merge(word, 1, Integer::sum)
// // In Go:   counts[word]++ (zero value of int is 0, so this just works)
// //
// // This is one of Go's most elegant patterns: zero values make maps
// // usable without initialization of individual entries.
// func TestWordCountMixed(t *testing.T) {
// 	got := WordCount([]string{"go", "java", "go", "rust", "go"})
//
// 	if got["go"] != 3 {
// 		t.Errorf("WordCount[\"go\"]: got %d, want 3", got["go"])
// 	}
// 	if got["java"] != 1 {
// 		t.Errorf("WordCount[\"java\"]: got %d, want 1", got["java"])
// 	}
// 	if got["rust"] != 1 {
// 		t.Errorf("WordCount[\"rust\"]: got %d, want 1", got["rust"])
// 	}
// }

// func TestWordCountEmpty(t *testing.T) {
// 	got := WordCount([]string{})
// 	if len(got) != 0 {
// 		t.Errorf("WordCount(empty): got %d entries, want 0", len(got))
// 	}
// }

// func TestWordCountAllUnique(t *testing.T) {
// 	got := WordCount([]string{"a", "b", "c"})
// 	for _, k := range []string{"a", "b", "c"} {
// 		if got[k] != 1 {
// 			t.Errorf("WordCount[%q]: got %d, want 1", k, got[k])
// 		}
// 	}
// }

// ============================================================================
// Step 3: Comma-ok and map operations — UniqueStrings and Keys
//
// After Step 2 passes, implement UniqueStrings and Keys in collections.go,
// then uncomment the tests below.
// NOTE: Add "sort" to the import block at the top of this file.
// ============================================================================

// // In Java: new LinkedHashSet<>(list) preserves order and removes dupes.
// // In Go: you build it manually with a map[string]bool as a set.
// // The comma-ok pattern (_, ok := m[key]) tells you if a key exists.
// func TestUniqueStringsDuplicates(t *testing.T) {
// 	got := UniqueStrings([]string{"go", "java", "go", "rust", "java"})
// 	want := []string{"go", "java", "rust"}
//
// 	if len(got) != len(want) {
// 		t.Fatalf("UniqueStrings: got %d elements, want %d", len(got), len(want))
// 	}
// 	for i := range want {
// 		if got[i] != want[i] {
// 			t.Errorf("UniqueStrings[%d]: got %q, want %q", i, got[i], want[i])
// 		}
// 	}
// }

// func TestUniqueStringsAlreadyUnique(t *testing.T) {
// 	got := UniqueStrings([]string{"a", "b", "c"})
// 	if len(got) != 3 {
// 		t.Errorf("UniqueStrings: got %d elements, want 3", len(got))
// 	}
// }

// func TestUniqueStringsEmpty(t *testing.T) {
// 	got := UniqueStrings([]string{})
// 	if len(got) != 0 {
// 		t.Errorf("UniqueStrings(empty): got %d elements, want 0", len(got))
// 	}
// }

// // Map iteration order is NOT guaranteed in Go. This test sorts the
// // result before comparing — otherwise it would be flaky.
// func TestKeysMultiple(t *testing.T) {
// 	got := Keys(map[string]int{"go": 3, "java": 1, "rust": 1})
// 	sort.Strings(got)
// 	want := []string{"go", "java", "rust"}
//
// 	if len(got) != len(want) {
// 		t.Fatalf("Keys: got %d elements, want %d", len(got), len(want))
// 	}
// 	for i := range want {
// 		if got[i] != want[i] {
// 			t.Errorf("Keys[%d]: got %q, want %q", i, got[i], want[i])
// 		}
// 	}
// }

// func TestKeysEmpty(t *testing.T) {
// 	got := Keys(map[string]int{})
// 	if len(got) != 0 {
// 		t.Errorf("Keys(empty): got %d elements, want 0", len(got))
// 	}
// }

// ============================================================================
// Step 4: Nil vs empty slice (just uncomment — no code to write)
//
// Read the comments carefully. This test demonstrates a gotcha, not a task.
// ============================================================================

// // THIS PARALLELS THE NIL INTERFACE TRAP FROM EXERCISE 03.
// //
// // In Java:
// //   List<String> a = null;              // null — NPE if you call methods
// //   List<String> b = new ArrayList<>();  // empty — safe to use
// //
// // In Go, both nil and empty slices are safe to use:
// //   var a []string     // nil — no backing array, but len/append/range all work
// //   b := []string{}    // empty — has a backing array (of size 0)
// //
// // The difference: a == nil is true, b == nil is false.
// // Both have len == 0. Both work with append and range.
// func TestNilVsEmptySlice(t *testing.T) {
// 	// A nil slice — declared but never initialized.
// 	var nilSlice []string
// 	if len(nilSlice) != 0 {
// 		t.Error("nil slice should have len 0")
// 	}
// 	if nilSlice != nil {
// 		t.Error("var []string should be nil")
// 	}
//
// 	// An empty slice — initialized to an empty literal.
// 	emptySlice := []string{}
// 	if len(emptySlice) != 0 {
// 		t.Error("empty slice should have len 0")
// 	}
// 	if emptySlice == nil {
// 		t.Error("[]string{} should NOT be nil")
// 	}
//
// 	// Both work perfectly fine with append.
// 	nilSlice = append(nilSlice, "works")
// 	emptySlice = append(emptySlice, "works")
// 	if len(nilSlice) != 1 || len(emptySlice) != 1 {
// 		t.Error("append should work on both nil and empty slices")
// 	}
// }

// ============================================================================
// CHALLENGE: GroupBy — combining slices and maps
//
// After all steps pass, implement GroupBy in collections.go,
// then uncomment the tests below.
// ============================================================================

// // In Java: Collectors.groupingBy — a single stream collector call.
// // In Go: a for loop with a map of slices. More explicit, equally readable.
// //
// // This function takes a key function — your first encounter with
// // function values as parameters. In Go, functions are first-class values,
// // just like in Java's functional interfaces (Function<T,R>).
// func TestGroupByFirstLetter(t *testing.T) {
// 	words := []string{"go", "goroutine", "java", "rust"}
// 	got := GroupBy(words, func(s string) string {
// 		return string(s[0])
// 	})
//
// 	// Check "g" group
// 	gGroup := got["g"]
// 	if len(gGroup) != 2 || gGroup[0] != "go" || gGroup[1] != "goroutine" {
// 		t.Errorf("GroupBy[\"g\"]: got %v, want [go goroutine]", gGroup)
// 	}
//
// 	// Check "j" group
// 	jGroup := got["j"]
// 	if len(jGroup) != 1 || jGroup[0] != "java" {
// 		t.Errorf("GroupBy[\"j\"]: got %v, want [java]", jGroup)
// 	}
//
// 	// Check "r" group
// 	rGroup := got["r"]
// 	if len(rGroup) != 1 || rGroup[0] != "rust" {
// 		t.Errorf("GroupBy[\"r\"]: got %v, want [rust]", rGroup)
// 	}
// }

// func TestGroupByLength(t *testing.T) {
// 	words := []string{"go", "java", "rust", "js"}
// 	got := GroupBy(words, func(s string) string {
// 		return fmt.Sprintf("%d", len(s))
// 	})
//
// 	twoGroup := got["2"]
// 	if len(twoGroup) != 2 || twoGroup[0] != "go" || twoGroup[1] != "js" {
// 		t.Errorf("GroupBy[\"2\"]: got %v, want [go js]", twoGroup)
// 	}
//
// 	fourGroup := got["4"]
// 	if len(fourGroup) != 2 || fourGroup[0] != "java" || fourGroup[1] != "rust" {
// 		t.Errorf("GroupBy[\"4\"]: got %v, want [java rust]", fourGroup)
// 	}
// }
