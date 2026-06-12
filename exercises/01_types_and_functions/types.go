// Package typesandfunctions
//
// In this exercise, YOU define everything from scratch — functions, types,
// and constants. Work through the steps in order. Each step tells you what
// to build. After implementing, uncomment the matching tests in types_test.go
// and run: task test -- 01
package typesandfunctions

// Add imports as you need them:
// - "errors" for creating error values (Step 2)

// ============================================================================
// Step 1: Variable declarations
//
// Uncomment Step 1 tests in types_test.go, then implement below.
// ============================================================================
//
// In Java:
//
//	int i = 42; String s = "hello"; boolean b = true;
//
// In Go, use := to declare AND assign. The compiler infers the type:
//
//	i := 42       // int
//	s := "hello"  // string
//	b := true     // bool
//
// Go also has var declarations, where all types have a zero value:
//
//	var x int     → 0
//	var s string  → ""
//	var b bool    → false
//	var p *int    → nil
//
// YOUR TASK:
// 1. Define a function called DeclareVariables that returns (int, string, bool).
//    Use := to declare variables with values 42, "hello", true and return them.
//
// 2. Define a function called ZeroValues that returns (int, string, bool, *int).
//    Declare variables with var (no assignment!) and return them.
//    Go initializes all variables to their zero value automatically.

// ============================================================================
// Step 2: Multiple return values
//
// After Step 1 tests pass, uncomment Step 2 tests in types_test.go.
// ============================================================================
//
// In Java, a method can only return one value. To signal failure,
// you throw an exception — but the caller can't see it in the signature.
//
// In Go, functions return multiple values. The pattern (result, error)
// is THE idiomatic way to signal success or failure:
//
//	func DoSomething(input string) (int, error) {
//	    // return value, nil          → success
//	    // return 0, errors.New("..") → failure
//	}
//
// YOUR TASK:
// 1. Define a function called Divide that takes two ints (a, b) and returns
//    (int, error).
//    - If b is zero, return 0 and an error with message "division by zero".
//      Hint: use errors.New("division by zero")
//    - Otherwise, return a / b and nil.
//    NOTE: Add "errors" to your imports.

// ============================================================================
// Step 3: Custom types, constants and iota
//
// After Step 2 tests pass, uncomment Step 3 tests in types_test.go.
// ============================================================================
//
// In Java, you'd define an enum:
//
//	enum Weekday { MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY }
//	Weekday.MONDAY.ordinal() == 0
//
// Go has no enum keyword. Instead, define a custom type and use iota
// in a const block to auto-increment values:
//
//	type Direction int
//	const (
//	    North Direction = iota  // 0
//	    East                    // 1
//	    South                   // 2
//	    West                    // 3
//	)
//
// In Java, toString() is inherited from Object.
// In Go, implement the fmt.Stringer interface (just a String() string method)
// and fmt.Println, Sprintf, etc. will use it automatically:
//
//	// from the fmt package:
//	type Stringer interface {
//	    String() string
//	}
//
// YOUR TASK:
// 1. Define a type called Weekday based on int.
//    Hint: type Weekday int
// 2. Define constants Monday through Friday using iota in a const block.
//    Skip zero so Monday = 1 (use _ as the first constant).
//    Hint:
//      const (
//          _ Weekday = iota  // 0, skipped
//          Monday            // 1
//          ...
//      )
// 3. Add a String() method on Weekday that returns the day's name
//    (e.g. Monday → "Monday"). Hint: use a switch statement.

// ============================================================================
// Step 4: Explicit type conversions
//
// After Step 3 tests pass, uncomment Step 4 tests in types_test.go.
// ============================================================================
//
// In Java, widening conversions are implicit:
//
//	int celsius = 100;
//	double fahrenheit = celsius * 9.0 / 5 + 32;  // int → double automatically
//
// In Go, ALL type conversions must be explicit. This won't compile:
//
//	var celsius int = 100
//	var f float64 = celsius * 9.0 / 5 + 32  // ERROR: mismatched types
//
// You must write float64(celsius) to convert explicitly.
//
// YOUR TASK:
// 1. Define a function called ConvertTemperature that takes an int (celsius)
//    and returns a float64 (fahrenheit).
//    Formula: F = C * 9/5 + 32
//    Hint: convert celsius to float64 first, then do the math.

// ============================================================================
// Step 5: Pointers
//
// After Step 4 tests pass, uncomment Step 5 tests in types_test.go.
// ============================================================================
//
// In Java, all objects are references (hidden pointers), and primitives
// can't be passed by reference at all.
//
// In Go, pointers are EXPLICIT. You choose what gets passed by reference:
//
//	*int  — "a pointer to an int"
//	&x    — "the address of x"
//	*p    — "the value p points to" (dereference)
//
//	func Double(p *int) {
//	    *p = *p * 2
//	}
//	x := 5
//	Double(&x)  // x is now 10
//
// YOUR TASK:
// 1. Define a function called Increment that takes a *int (pointer to int)
//    and adds 1 to the value it points to. It returns nothing.
//
// 2. Define a function called NewInt that takes an int value and returns
//    a *int (pointer to int). Create the int and return its address.
//    In Java, this would be like Integer.valueOf() — auto-boxing.
//    In Go, you can return &localVar — the compiler moves it to the heap.
