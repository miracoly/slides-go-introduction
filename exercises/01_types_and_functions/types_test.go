package typesandfunctions

import (
	// Uncomment "testing" when you uncomment your first test.
	// "testing"
)

// ============================================================================
// Step 1: Variable declarations
//
// Implement DeclareVariables and ZeroValues in types.go,
// then uncomment the tests below.
// ============================================================================

// // In Java:  int i = 42; String s = "hello"; boolean b = true;
// // In Go:    i := 42      s := "hello"        b := true
// //
// // The := operator declares AND assigns. The compiler infers the type.
// // You never need to write the type when the right-hand side makes it obvious.
// func TestDeclareVariables(t *testing.T) {
// 	i, s, b := DeclareVariables()
//
// 	if i != 42 {
// 		t.Errorf("int: got %d, want 42", i)
// 	}
// 	if s != "hello" {
// 		t.Errorf("string: got %q, want \"hello\"", s)
// 	}
// 	if b != true {
// 		t.Errorf("bool: got %v, want true", b)
// 	}
// }

// // In Java, uninitialized local variables cause a compile error:
// //
// //	int x;
// //	System.out.println(x); // ERROR: variable x might not have been initialized
// //
// // In Go, every type has a defined zero value. Variables declared with var
// // (without assignment) are automatically set to that zero value.
// //
// //	var x int    → 0
// //	var s string → ""
// //	var b bool   → false
// //	var p *int   → nil
// func TestZeroValues(t *testing.T) {
// 	i, s, b, p := ZeroValues()
//
// 	if i != 0 {
// 		t.Errorf("int zero value: got %d, want 0", i)
// 	}
// 	if s != "" {
// 		t.Errorf("string zero value: got %q, want \"\"", s)
// 	}
// 	if b != false {
// 		t.Errorf("bool zero value: got %v, want false", b)
// 	}
// 	if p != nil {
// 		t.Errorf("pointer zero value: got %v, want nil", p)
// 	}
// }

// ============================================================================
// Step 2: Multiple return values and error handling
//
// After Step 1 passes, implement Divide in types.go.
// Then uncomment the tests below.
// NOTE: Add "errors" to your imports in types.go.
// ============================================================================

// // In Java, a method can only return one value. To return multiple values,
// // you'd create a wrapper class, use a Pair<A,B>, or throw an exception.
// //
// // In Go, functions natively return multiple values. The pattern
// //
// //	(result, error)
// //
// // is THE idiomatic way to signal success or failure. No exceptions needed.
// func TestDivide(t *testing.T) {
// 	result, err := Divide(10, 3)
// 	if err != nil {
// 		t.Fatalf("Divide(10, 3): unexpected error: %v", err)
// 	}
// 	if result != 3 {
// 		t.Errorf("Divide(10, 3): got %d, want 3", result)
// 	}
// }

// func TestDivideByZero(t *testing.T) {
// 	result, err := Divide(10, 0)
// 	if err == nil {
// 		t.Fatal("Divide(10, 0): expected error, got nil")
// 	}
// 	if err.Error() != "division by zero" {
// 		t.Errorf("Divide(10, 0): error message: got %q, want \"division by zero\"", err.Error())
// 	}
// 	if result != 0 {
// 		t.Errorf("Divide(10, 0): result: got %d, want 0", result)
// 	}
// }

// ============================================================================
// Step 3: Custom types, constants and iota
//
// After Step 2 passes, define Weekday type, constants, and String() method
// in types.go. Then uncomment the tests below.
// ============================================================================

// // In Java, you'd define an enum:
// //
// //	enum Weekday { MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY }
// //	Weekday.MONDAY.ordinal() == 0
// //
// // Go has no enum keyword. Instead, you define a type and use iota in a const
// // block to auto-increment values. By convention, skip 0 so the zero value
// // of the type means "unset" rather than a valid weekday.
// func TestWeekdayValues(t *testing.T) {
// 	tests := []struct {
// 		day  Weekday
// 		want int
// 	}{
// 		{Monday, 1},
// 		{Tuesday, 2},
// 		{Wednesday, 3},
// 		{Thursday, 4},
// 		{Friday, 5},
// 	}
//
// 	for _, tt := range tests {
// 		if int(tt.day) != tt.want {
// 			t.Errorf("%v: got %d, want %d", tt.day, int(tt.day), tt.want)
// 		}
// 	}
// }

// // In Java: Weekday.MONDAY.toString() returns "MONDAY" (or a custom override).
// // In Go: implement the fmt.Stringer interface (just a String() string method)
// // and fmt.Println, t.Errorf, etc. will use it automatically.
// func TestWeekdayString(t *testing.T) {
// 	tests := []struct {
// 		day  Weekday
// 		want string
// 	}{
// 		{Monday, "Monday"},
// 		{Tuesday, "Tuesday"},
// 		{Wednesday, "Wednesday"},
// 		{Thursday, "Thursday"},
// 		{Friday, "Friday"},
// 	}
//
// 	for _, tt := range tests {
// 		got := tt.day.String()
// 		if got != tt.want {
// 			t.Errorf("Weekday(%d).String(): got %q, want %q", int(tt.day), got, tt.want)
// 		}
// 	}
// }

// ============================================================================
// Step 4: Explicit type conversions
//
// After Step 3 passes, implement ConvertTemperature in types.go.
// Then uncomment the tests below.
// ============================================================================

// // In Java, widening conversions are implicit:
// //
// //	int celsius = 100;
// //	double fahrenheit = celsius * 9.0 / 5 + 32; // int silently widens to double
// //
// // In Go, ALL type conversions must be explicit. This won't compile:
// //
// //	var celsius int = 100
// //	var f float64 = celsius * 9.0 / 5 + 32  // ERROR: mismatched types
// //
// // You must write float64(celsius) to convert explicitly.
// func TestConvertTemperature(t *testing.T) {
// 	tests := []struct {
// 		celsius int
// 		want    float64
// 	}{
// 		{100, 212.0},
// 		{0, 32.0},
// 		{-40, -40.0},
// 	}
//
// 	for _, tt := range tests {
// 		got := ConvertTemperature(tt.celsius)
// 		if got != tt.want {
// 			t.Errorf("ConvertTemperature(%d): got %f, want %f", tt.celsius, got, tt.want)
// 		}
// 	}
// }

// ============================================================================
// Step 5: Pointers
//
// After Step 4 passes, implement Increment and NewInt in types.go.
// Then uncomment the tests below.
// ============================================================================

// // In Java, primitives (int, boolean, etc.) are always passed by value.
// // Objects are always passed by reference (well, the reference is passed by value).
// // There's no way to pass a primitive "by reference."
// //
// // In Go, pointers are explicit. You decide what gets passed by reference:
// //
// //	*int  — "a pointer to an int"
// //	&x    — "the address of x"
// //	*p    — "the value p points to" (dereference)
// func TestIncrement(t *testing.T) {
// 	x := 5
// 	Increment(&x)
// 	if x != 6 {
// 		t.Errorf("after Increment(&x): got %d, want 6", x)
// 	}
// }

// func TestIncrementMultiple(t *testing.T) {
// 	x := 0
// 	Increment(&x)
// 	Increment(&x)
// 	Increment(&x)
// 	if x != 3 {
// 		t.Errorf("after 3x Increment(&x): got %d, want 3", x)
// 	}
// }

// // In Java: Integer x = Integer.valueOf(42); — auto-boxing a primitive.
// // In Go: you can return the address of a local variable. The compiler
// // automatically moves it to the heap (called "escape analysis").
// // This is perfectly safe — no dangling pointers.
// func TestNewInt(t *testing.T) {
// 	p := NewInt(42)
// 	if p == nil {
// 		t.Fatal("NewInt(42): got nil, want non-nil pointer")
// 	}
// 	if *p != 42 {
// 		t.Errorf("NewInt(42): got %d, want 42", *p)
// 	}
// }

// func TestNewIntUniqueness(t *testing.T) {
// 	p1 := NewInt(1)
// 	p2 := NewInt(1)
// 	if p1 == p2 {
// 		t.Error("NewInt should return distinct pointers for each call")
// 	}
// }

// // Verify that the error from Divide has a meaningful message and is non-nil.
// // This is a sneak peek at how Go uses interfaces — you'll explore this
// // more deeply in exercise 03.
// func TestDivideErrorMessage(t *testing.T) {
// 	_, err := Divide(1, 0)
// 	if err == nil {
// 		t.Fatal("expected error")
// 	}
//
// 	// The error interface has a single method: Error() string.
// 	// errors.New and fmt.Errorf both return values satisfying this interface.
// 	msg := err.Error()
// 	if msg != "division by zero" {
// 		t.Errorf("error message: got %q, want \"division by zero\"", msg)
// 	}
// }
