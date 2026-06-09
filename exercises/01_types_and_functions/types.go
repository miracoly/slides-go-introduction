// Package typesandfunctions
package typesandfunctions

// Weekday represents a day of the work week.
// In Java, you'd use an enum: enum Weekday { MONDAY, TUESDAY, ... }
// In Go, we define a custom type and use iota to generate sequential values.
type Weekday int

const (
	_ Weekday = iota // skip zero so Monday = 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

// String returns the name of the weekday.
// This is Go's equivalent of Java's toString() — implement the fmt.Stringer interface.
/*

// from fmt package
type Stringer interface {
    String() string
}

*/
func (d Weekday) String() string {
	panic("TODO: implement")
}

// DeclareVariables demonstrates Go's short variable declaration (:=).
// In Java: int i = 42; String s = "hello"; boolean b = true;
// In Go: use := and let the compiler infer the types.
func DeclareVariables() (int, string, bool) {
	panic("TODO: implement")
}

// ZeroValues demonstrates that Go initializes all variables to their zero value.
// In Java, uninitialized local variables cause a compile error.
// In Go, every type has a zero value: 0, "", false, nil.
//
// Declare variables with var (no assignment) and return them.
func ZeroValues() (int, string, bool, *int) {
	panic("TODO: implement")
}

// Divide performs integer division and returns an error on division by zero.
// In Java, you'd throw an ArithmeticException.
// In Go, errors are values — return them as a second return value.
// use errors.New()
func Divide(a, b int) (int, error) {
	panic("TODO: implement")
}

// ConvertTemperature converts Celsius (int) to Fahrenheit (float64).
// Formula: F = C * 9/5 + 32
//
// In Java, int-to-double conversion is implicit: double f = celsius * 9.0/5 + 32;
// In Go, you MUST explicitly convert: float64(celsius). No implicit widening.
func ConvertTemperature(celsius int) float64 {
	panic("TODO: implement")
}

// Increment adds 1 to the value pointed to by p.
// In Java, all objects are references (hidden pointers), and primitives can't be passed by reference.
// In Go, pointers are explicit: *p dereferences, &v takes address.
func Increment(p *int) {
	panic("TODO: implement")
}

// NewInt allocates an int with the given value and returns a pointer to it.
// In Java, you'd use Integer.valueOf() or auto-boxing.
// In Go, you can take the address of a local variable — the compiler moves it to the heap.
func NewInt(value int) *int {
	panic("TODO: implement")
}
