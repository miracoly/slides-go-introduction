package interfaces

import (
	// Add "fmt" to this import block when you reach Step 3.
	// Uncomment "math" and "testing" when you uncomment your first test.
	// "math"
	// "testing"
)

// ============================================================================
// Step 1: Structs and methods
//
// Define Circle and Rectangle structs with Area() and Perimeter() methods
// in interfaces.go, then uncomment the tests below.
// ============================================================================

// // No interface yet! These tests call methods directly on the structs.
// // In Java, you'd test the same way — but here, notice there's no Shape
// // interface involved. The methods exist on the struct, not on an interface.
// func TestCircleAreaAndPerimeter(t *testing.T) {
// 	c := Circle{Radius: 5}
//
// 	area := c.Area()
// 	wantArea := math.Pi * 25
// 	if math.Abs(area-wantArea) > 0.0001 {
// 		t.Errorf("Circle.Area(): got %f, want %f", area, wantArea)
// 	}
//
// 	peri := c.Perimeter()
// 	wantPeri := 2 * math.Pi * 5
// 	if math.Abs(peri-wantPeri) > 0.0001 {
// 		t.Errorf("Circle.Perimeter(): got %f, want %f", peri, wantPeri)
// 	}
// }

// func TestRectangleAreaAndPerimeter(t *testing.T) {
// 	r := Rectangle{Width: 3, Height: 4}
//
// 	if r.Area() != 12.0 {
// 		t.Errorf("Rectangle.Area(): got %f, want 12.0", r.Area())
// 	}
// 	if r.Perimeter() != 14.0 {
// 		t.Errorf("Rectangle.Perimeter(): got %f, want 14.0", r.Perimeter())
// 	}
// }

// ============================================================================
// Step 2: Define the Shape interface
//
// After Step 1 passes, define Shape interface and TotalArea in interfaces.go,
// then uncomment the tests below.
// ============================================================================

// // In Java:
// //
// //	public class Circle implements Shape { ... }
// //
// // The "implements" keyword is REQUIRED. Without it, the compiler rejects
// // using Circle as a Shape, even if all methods match.
// //
// // In Go: there is no "implements". Circle already has Area() and Perimeter()
// // from Step 1 — so it ALREADY satisfies Shape. You just had to name the contract.
// func TestCircleSatisfiesShape(t *testing.T) {
// 	c := Circle{Radius: 5}
//
// 	// This assignment would fail at compile time if Circle
// 	// didn't have the right methods. That's the entire check.
// 	var s Shape = c
//
// 	got := s.Area()
// 	want := math.Pi * 25
// 	if math.Abs(got-want) > 0.0001 {
// 		t.Errorf("Circle.Area() via Shape: got %f, want %f", got, want)
// 	}
// }

// func TestRectangleSatisfiesShape(t *testing.T) {
// 	r := Rectangle{Width: 3, Height: 4}
//
// 	var s Shape = r
//
// 	if s.Area() != 12.0 {
// 		t.Errorf("Rectangle.Area() via Shape: got %f, want 12.0", s.Area())
// 	}
// }

// // This is the most important Go design idiom:
// //
// //	"Accept interfaces, return structs."
// //
// // TotalArea accepts []Shape — it doesn't care about concrete types.
// // Any type satisfying Shape works, even types that don't exist yet.
// //
// // In Java, this is conceptually the same (List<Shape>), but every type
// // must explicitly declare "implements Shape". In Go, a type from a
// // completely different package can satisfy Shape without knowing it exists.
// func TestTotalArea(t *testing.T) {
// 	shapes := []Shape{
// 		Circle{Radius: 1},
// 		Rectangle{Width: 2, Height: 3},
// 	}
//
// 	got := TotalArea(shapes)
// 	want := math.Pi + 6.0
//
// 	if math.Abs(got-want) > 0.0001 {
// 		t.Errorf("TotalArea: got %f, want %f", got, want)
// 	}
// }

// func TestTotalAreaEmpty(t *testing.T) {
// 	got := TotalArea([]Shape{})
// 	if got != 0.0 {
// 		t.Errorf("TotalArea(empty): got %f, want 0.0", got)
// 	}
// }

// ============================================================================
// Step 3: fmt.Stringer and Describer interface
//
// After Step 2 passes, implement String(), define Describer, and add
// Describe() in interfaces.go. Then uncomment the tests below.
// NOTE: Add "fmt" to the import block at the top of this file.
// ============================================================================

// // In Java: every class has toString(). You override it.
// // In Go: implement the fmt.Stringer interface (String() string).
// // fmt.Println, fmt.Sprint, etc. will use it automatically.
// //
// // The key insight: fmt.Stringer is just a regular interface.
// // There's nothing special about it — it follows the exact same rules.
// func TestCircleStringer(t *testing.T) {
// 	c := Circle{Radius: 5}
// 	got := fmt.Sprint(c)
// 	want := "Circle(r=5.00)"
//
// 	if got != want {
// 		t.Errorf("fmt.Sprint(Circle{5}): got %q, want %q", got, want)
// 	}
// }

// // Describer is a single-method interface. In Go, these are the most
// // powerful kind — they're easy to implement and highly composable.
// //
// // The stdlib is full of them:
// //   - fmt.Stringer:  String() string
// //   - error:         Error() string
// //   - io.Reader:     Read(p []byte) (n int, err error)
// //   - io.Writer:     Write(p []byte) (n int, err error)
// func TestDescriberPolymorphism(t *testing.T) {
// 	describers := []Describer{
// 		Circle{Radius: 5},
// 		Rectangle{Width: 3, Height: 4},
// 	}
//
// 	wants := []string{
// 		"I am a circle with radius 5.00",
// 		"I am a 3.00x4.00 rectangle",
// 	}
//
// 	for i, d := range describers {
// 		got := d.Describe()
// 		if got != wants[i] {
// 			t.Errorf("Describe()[%d]: got %q, want %q", i, got, wants[i])
// 		}
// 	}
// }

// ============================================================================
// Step 4: Type assertions and type switch
//
// After Step 3 passes, implement DescribeShape in interfaces.go,
// then uncomment the tests below.
// ============================================================================

// // In Java:
// //
// //	if (shape instanceof Circle c) {
// //	    System.out.println(c.getRadius());
// //	}
// //
// // In Go:
// //
// //	if c, ok := shape.(Circle); ok {
// //	    fmt.Println(c.Radius)
// //	}
// //
// // The comma-ok pattern is Go's safe cast. No ClassCastException risk.
// func TestTypeAssertionSuccess(t *testing.T) {
// 	var s Shape = Circle{Radius: 5}
//
// 	c, ok := s.(Circle)
// 	if !ok {
// 		t.Fatal("type assertion to Circle failed")
// 	}
// 	if c.Radius != 5 {
// 		t.Errorf("Radius: got %f, want 5", c.Radius)
// 	}
// }

// func TestTypeAssertionFailure(t *testing.T) {
// 	var s Shape = Rectangle{Width: 3, Height: 4}
//
// 	_, ok := s.(Circle)
// 	if ok {
// 		t.Error("type assertion to Circle should have failed for Rectangle")
// 	}
// }

// // In Java: a chain of instanceof checks.
// // In Go: a type switch — much cleaner.
// //
// //	switch v := s.(type) {
// //	case Circle:    // v is Circle here
// //	case Rectangle: // v is Rectangle here
// //	default:        // unknown type
// //	}
// func TestDescribeShapeCircle(t *testing.T) {
// 	got := DescribeShape(Circle{Radius: 5})
// 	want := "circle with radius 5.00"
// 	if got != want {
// 		t.Errorf("DescribeShape(Circle): got %q, want %q", got, want)
// 	}
// }

// func TestDescribeShapeRectangle(t *testing.T) {
// 	got := DescribeShape(Rectangle{Width: 3, Height: 4})
// 	want := "rectangle 3.00x4.00"
// 	if got != want {
// 		t.Errorf("DescribeShape(Rectangle): got %q, want %q", got, want)
// 	}
// }

// ============================================================================
// Step 5: The nil interface trap (just uncomment — no code to write)
//
// Read the comments carefully. This test demonstrates a gotcha, not a task.
// ============================================================================

// // THIS IS THE MOST SUBTLE GOTCHA IN GO.
// //
// // In Java, null is null. Period.
// //   Object o = null;
// //   if (o == null) { ... } // always true
// //
// // In Go, an interface value has TWO components: (type, value).
// // An interface is nil ONLY when BOTH are nil.
// //
// //   var s Shape = nil         → s is nil (no type, no value)
// //   var c *Circle = nil
// //   var s Shape = c           → s is NOT nil! (type=*Circle, value=nil)
// //
// // This catches EVERY Go developer at some point. Understanding it now
// // will save you hours of debugging later.
// func TestNilInterfaceTrap(t *testing.T) {
// 	// Case 1: A truly nil interface — no type, no value.
// 	var nilShape Shape
// 	if nilShape != nil {
// 		t.Error("a zero-valued interface should be nil")
// 	}
//
// 	// Case 2: An interface holding a nil pointer.
// 	// The interface knows the TYPE (*Circle) even though the VALUE is nil.
// 	// Therefore the interface itself is NOT nil.
// 	var nilCircle *Circle = nil
// 	var shapeWithNilCircle Shape = nilCircle
//
// 	if shapeWithNilCircle == nil {
// 		t.Error("an interface holding a nil *Circle should NOT be nil — it has type information")
// 	}
// }

// ============================================================================
// CHALLENGE: Interface composition
//
// After all steps pass, define ShapeDescriber in interfaces.go,
// then uncomment the tests below.
// ============================================================================

// // In Java:
// //
// //	interface ShapeDescriber extends Shape, Describer {}
// //
// // In Go, interface composition works the same way:
// //
// //	type ShapeDescriber interface {
// //	    Shape
// //	    Describer
// //	}
// //
// // Any type that satisfies ALL embedded interfaces satisfies the composed one.
// func TestCircleSatisfiesShapeDescriber(t *testing.T) {
// 	c := Circle{Radius: 5}
//
// 	// This assignment verifies Circle satisfies the composed interface.
// 	var sd ShapeDescriber = c
//
// 	got := sd.Describe()
// 	if got != "I am a circle with radius 5.00" {
// 		t.Errorf("Describe(): got %q, want \"I am a circle with radius 5.00\"", got)
// 	}
//
// 	area := sd.Area()
// 	want := math.Pi * 25
// 	if math.Abs(area-want) > 0.0001 {
// 		t.Errorf("Area(): got %f, want %f", area, want)
// 	}
// }

// // Rectangle also satisfies ShapeDescriber — it has all the methods.
// func TestRectangleSatisfiesShapeDescriber(t *testing.T) {
// 	r := Rectangle{Width: 3, Height: 4}
//
// 	var sd ShapeDescriber = r
//
// 	got := sd.Describe()
// 	if got != "I am a 3.00x4.00 rectangle" {
// 		t.Errorf("Describe(): got %q, want \"I am a 3.00x4.00 rectangle\"", got)
// 	}
//
// 	if sd.Area() != 12.0 {
// 		t.Errorf("Area(): got %f, want 12.0", sd.Area())
// 	}
// }
