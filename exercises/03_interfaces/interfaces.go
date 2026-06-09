// Package interfaces
//
// In this exercise, YOU define everything from scratch — structs, methods,
// interfaces, and functions. Work through the steps in order.
// Each step tells you what to build. After implementing, uncomment the
// matching tests in interfaces_test.go and run: task test -- 03
package interfaces

import (
	"fmt"
	"math"
)

// ============================================================================
// Step 1: Define structs and methods
//
// Uncomment Step 1 tests in interfaces_test.go, then implement below.
// ============================================================================
//
// In Java, you'd write a class with fields and methods:
//
//	public class Circle {
//	    public double radius;
//	    public double area() { return Math.PI * radius * radius; }
//	    public double perimeter() { return 2 * Math.PI * radius; }
//	}
//
// In Go, data (struct) and behavior (methods) are separate:
//
//	type Circle struct { ... }
//	func (c Circle) Area() float64 { ... }
//
// YOUR TASK:
// 1. Define a struct called Circle with a single field Radius of type float64
// 2. Add two methods on Circle:
//    - Area returns the area (pi * r²). Hint: use math.Pi
//    - Perimeter returns the circumference (2 * pi * r)
//    Both methods return float64.
// 3. Define a struct called Rectangle with fields Width and Height (both float64)
// 4. Add two methods on Rectangle:
//    - Area returns width * height
//    - Perimeter returns 2 * (width + height)
// 5. Uncomment Step 1 tests, run: task test -- 03

// ============================================================================
// Step 2: Define the Shape interface
//
// After Step 1 tests pass, uncomment Step 2 tests in interfaces_test.go.
// ============================================================================
//
// In Java, you'd write:
//
//	public interface Shape {
//	    double area();
//	    double perimeter();
//	}
//	public class Circle implements Shape { ... }
//
// In Go, interfaces are just method sets. Any type that has the right
// methods satisfies it — no "implements" keyword. Circle and Rectangle
// ALREADY satisfy Shape from Step 1. You just need to name the contract.
//
// YOUR TASK:
// 1. Define an interface called Shape with two methods: Area and Perimeter,
//    both returning float64. Circle and Rectangle already have these methods
//    from Step 1 — they will satisfy Shape automatically.
// 2. Write a function called TotalArea that takes a slice of Shape and returns
//    the sum of all their areas as float64.
// 3. Uncomment Step 2 tests, run: task test -- 03

// ============================================================================
// Step 3: Implement fmt.Stringer and a custom Describer interface
//
// After Step 2 tests pass, uncomment Step 3 tests in interfaces_test.go.
// ============================================================================
//
// In Java, every class has toString(). You override it:
//   @Override public String toString() { return "Circle(r=5.00)"; }
//
// In Go, implement the fmt.Stringer interface:
//   type Stringer interface { String() string }
//
// There's nothing magic about it — it's a regular interface.
// fmt.Println, fmt.Sprint, etc. check if the value satisfies Stringer.
//
// YOUR TASK:
// 1. Add a String method on Circle that returns a string like "Circle(r=5.00)".
//    This makes Circle satisfy fmt.Stringer. Hint: use fmt.Sprintf with %.2f
// 2. Define a single-method interface called Describer with one method:
//    Describe, which returns a string.
//    Go proverb: "The bigger the interface, the weaker the abstraction."
//    Single-method interfaces are the norm: fmt.Stringer, error, io.Reader.
// 3. Add a Describe method on Circle that returns a string like
//    "I am a circle with radius 5.00"
// 4. Add a Describe method on Rectangle that returns a string like
//    "I am a 3.00x4.00 rectangle"
// 5. Uncomment Step 3 tests, run: task test -- 03

// ============================================================================
// Step 4: Type assertions and type switch
//
// After Step 3 tests pass, uncomment Step 4 tests in interfaces_test.go.
// ============================================================================
//
// In Java:
//   if (shape instanceof Circle c) { System.out.println(c.getRadius()); }
//
// In Go, you have two tools:
//
// Type assertion (safe cast with comma-ok):
//   c, ok := shape.(Circle)   // ok is false if shape isn't a Circle
//
// Type switch (cleaner than instanceof chains):
//   switch v := s.(type) {
//   case Circle:    // v is Circle here
//   case Rectangle: // v is Rectangle here
//   default:        // unknown type
//   }
//
// YOUR TASK:
// 1. Write a function called DescribeShape that takes a Shape and returns
//    a string. Use a type switch to inspect the concrete type and return:
//    - For a Circle:    "circle with radius X.XX"
//    - For a Rectangle: "rectangle X.XXxX.XX"
//    - For anything else: "unknown shape"
// 2. Uncomment Step 4 tests, run: task test -- 03

// ============================================================================
// Step 5: The nil interface trap (just uncomment — no code to write)
//
// After Step 4 tests pass, uncomment Step 5 tests in interfaces_test.go.
// ============================================================================
//
// THIS IS THE MOST SUBTLE GOTCHA IN GO.
//
// In Java, null is null. Period.
//
// In Go, an interface value has TWO components: (type, value).
// An interface is nil ONLY when BOTH components are nil.
//
//   var s Shape             → nil (no type, no value)
//   var c *Circle = nil
//   var s Shape = c         → NOT nil! (type=*Circle, value=nil)
//
// Read the test carefully. No code to write — just understand it.

// ============================================================================
// CHALLENGE: Interface Composition
//
// After all steps pass, uncomment the challenge tests in interfaces_test.go.
// ============================================================================
//
// In Java:
//   interface ShapeDescriber extends Shape, Describer {}
//
// In Go, you embed interfaces inside another interface — the same syntax
// as struct embedding.
//
// YOUR TASK:
// 1. Define an interface called ShapeDescriber that embeds both Shape
//    and Describer. Any type satisfying both will automatically satisfy
//    ShapeDescriber too.
// 2. Uncomment the challenge tests, run: task test -- 03
//
// Circle and Rectangle already have all the methods — they will satisfy
// ShapeDescriber automatically without any changes.

// Keep the compiler happy — these are used in later steps.
var (
	_ = fmt.Sprintf
	_ = math.Pi
)
