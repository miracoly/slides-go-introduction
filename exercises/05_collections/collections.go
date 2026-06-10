// Package collections
//
// In this exercise, YOU define everything from scratch — functions that
// operate on slices and maps. Work through the steps in order.
// Each step tells you what to build. After implementing, uncomment the
// matching tests in collections_test.go and run: task test -- 05
package collections

import (
	"strings"
)

// ============================================================================
// Step 1: Slice basics — Contains and FilterByPrefix
//
// Uncomment Step 1 tests in collections_test.go, then implement below.
// ============================================================================
//
// In Java, you'd use ArrayList and streams:
//
//	list.contains("java");                              // contains
//	list.stream().filter(s -> s.startsWith("go")).toList(); // filter
//
// In Go, there is no built-in contains or filter. Slices are simpler —
// they're just backed by arrays. You iterate with a for-range loop:
//
//	for i, v := range slice {
//	    // i is the index, v is the value
//	}
//
// To build a new slice, start with nil and use append:
//
//	var result []string
//	result = append(result, "item")
//
// YOUR TASK:
// 1. Write a function called Contains that takes a slice of strings
//    (haystack) and a single string (needle), and returns true if the
//    needle is found in the haystack.
//    Hint: use a range loop to iterate the slice
// 2. Write a function called FilterByPrefix that takes a slice of strings
//    and a prefix string, and returns a new slice containing only the
//    strings that start with the prefix.
//    Hint: use strings.HasPrefix to check, and append to build the result
// 3. Uncomment Step 1 tests, run: task test -- 05

// ============================================================================
// Step 2: Map basics — WordCount
//
// After Step 1 tests pass, uncomment Step 2 tests in collections_test.go.
// ============================================================================
//
// In Java, you'd use HashMap:
//
//	Map<String, Integer> counts = new HashMap<>();
//	for (String w : words) {
//	    counts.merge(w, 1, Integer::sum);
//	}
//
// In Go, you create a map with make or a literal, then use index syntax:
//
//	m := make(map[string]int)   // or map[string]int{}
//	m["key"] = 42               // set
//	v := m["key"]               // get (returns zero value if missing)
//	m["key"]++                  // increment (works even if key is new!)
//
// The zero value of int is 0, so m["newkey"]++ starts from 0.
// No need for containsKey or getOrDefault — zero values handle it.
//
// YOUR TASK:
// 1. Write a function called WordCount that takes a slice of strings and
//    returns a map from each word to the number of times it appears.
//    Hint: create a map, range over the words, and increment the count
// 2. Uncomment Step 2 tests, run: task test -- 05

// ============================================================================
// Step 3: Comma-ok and map operations — UniqueStrings and Keys
//
// After Step 2 tests pass, uncomment Step 3 tests in collections_test.go.
// ============================================================================
//
// In Java:
//
//	if (map.containsKey("key")) { ... }
//	Set<String> keys = map.keySet();
//
// In Go, there is no containsKey. Instead, you use the comma-ok pattern:
//
//	v, ok := m["key"]   // ok is true if key exists, false if not
//	if ok { ... }
//
// This is the same comma-ok pattern from type assertions (exercise 03)
// and will appear again with channel receives (exercise 06).
//
// To get all keys, you range over the map:
//
//	for k, v := range m {
//	    // k is the key, v is the value
//	}
//
// WARNING: Map iteration order is NOT guaranteed in Go. Unlike Java's
// LinkedHashMap, Go randomizes map iteration on purpose.
//
// YOUR TASK:
// 1. Write a function called UniqueStrings that takes a slice of strings
//    and returns a new slice with duplicates removed, preserving the order
//    of first occurrence.
//    Hint: use a map[string]bool as a "set". Before appending a word,
//    check if it's already in the set using the comma-ok pattern.
// 2. Write a function called Keys that takes a map[string]int and returns
//    a slice of all the keys.
//    Hint: range over the map and append each key
// 3. Uncomment Step 3 tests, run: task test -- 05

// ============================================================================
// Step 4: Nil vs empty slice (just uncomment — no code to write)
//
// After Step 3 tests pass, uncomment Step 4 tests in collections_test.go.
// ============================================================================
//
// In Java, null and empty list are clearly different:
//
//	List<String> a = null;              // null
//	List<String> b = new ArrayList<>(); // empty
//
// In Go, slices have a similar but subtler distinction:
//
//	var a []string        // nil slice — no underlying array
//	b := []string{}       // empty slice — has an array, but it's empty
//
// Both have len == 0. Both work with append, range, etc.
// But a == nil is true, b == nil is false.
//
// This is the slice equivalent of the nil interface trap (exercise 03).
// Read the test carefully. No code to write — just understand it.

// ============================================================================
// CHALLENGE: GroupBy — combining slices and maps
//
// After all steps pass, uncomment the challenge tests in collections_test.go.
// ============================================================================
//
// In Java, you'd use Collectors.groupingBy:
//
//	Map<String, List<String>> groups = words.stream()
//	    .collect(Collectors.groupingBy(w -> String.valueOf(w.charAt(0))));
//
// In Go, you build it yourself — it's just a map with slice values:
//
//	groups := map[string][]string{}
//	groups["key"] = append(groups["key"], "value")
//
// YOUR TASK:
// 1. Write a function called GroupBy that takes a slice of strings and a
//    key function (func(string) string), and returns a map from each key
//    to the slice of strings that produced that key.
//    - For each word, call the key function to get the group key
//    - Append the word to the slice for that key in the map
// 2. Uncomment the challenge tests, run: task test -- 05

// Keep the compiler happy — these are used in later steps.
var _ = strings.HasPrefix
