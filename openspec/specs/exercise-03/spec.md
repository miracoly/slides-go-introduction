## ADDED Requirements

### Requirement: Shape interface with implicit satisfaction
The exercise SHALL define a `Shape` interface with methods `Area() float64` and `Perimeter() float64`. Structs `Circle` (with field `Radius float64`) and `Rectangle` (with fields `Width, Height float64`) SHALL implicitly satisfy this interface by implementing both methods.

#### Scenario: Circle satisfies Shape
- **WHEN** a `Circle{Radius: 5}` is used as a `Shape`
- **THEN** `Area()` returns `78.53981633974483` (pi * r^2) and `Perimeter()` returns `31.41592653589793` (2 * pi * r)

#### Scenario: Rectangle satisfies Shape
- **WHEN** a `Rectangle{Width: 3, Height: 4}` is used as a `Shape`
- **THEN** `Area()` returns `12.0` and `Perimeter()` returns `14.0`

#### Scenario: No implements keyword
- **WHEN** `Circle` and `Rectangle` are defined
- **THEN** neither type contains any explicit interface satisfaction declaration — the compiler verifies it structurally

### Requirement: Implementing fmt.Stringer
The exercise SHALL include a `String() string` method on `Circle` so it satisfies `fmt.Stringer`. This demonstrates that stdlib interfaces work identically to user-defined ones.

#### Scenario: Circle as Stringer
- **WHEN** `fmt.Sprint(Circle{Radius: 5})` is called
- **THEN** it returns `"Circle(r=5.00)"`

### Requirement: Accept interfaces as parameters
The exercise SHALL include a function `TotalArea(shapes []Shape) float64` that accepts a slice of the `Shape` interface and returns the sum of all areas. This demonstrates the "accept interfaces" idiom.

#### Scenario: Total area of mixed shapes
- **WHEN** `TotalArea` is called with `[]Shape{Circle{Radius: 1}, Rectangle{Width: 2, Height: 3}}`
- **THEN** it returns the sum of their areas: `pi + 6.0`

### Requirement: Describer interface and polymorphism
The exercise SHALL define a `Describer` interface with a single method `Describe() string`. Multiple types SHALL satisfy it, demonstrating polymorphism through interfaces rather than inheritance.

#### Scenario: Multiple types satisfy Describer
- **WHEN** `Circle{Radius: 5}` implements `Describe()` returning `"I am a circle with radius 5.00"`
- **THEN** it can be used anywhere a `Describer` is expected

#### Scenario: Rectangle describes itself
- **WHEN** `Rectangle{Width: 3, Height: 4}` implements `Describe()` returning `"I am a 3.00x4.00 rectangle"`
- **THEN** it can be used anywhere a `Describer` is expected

### Requirement: Type assertion
The exercise SHALL include tests demonstrating type assertions to extract a concrete type from an interface value, paralleling Java's `instanceof` + cast pattern.

#### Scenario: Successful type assertion
- **WHEN** a `Shape` variable holding a `Circle{Radius: 5}` is asserted to `Circle`
- **THEN** the assertion succeeds and the `Circle` value is accessible with `Radius == 5`

#### Scenario: Failed type assertion with comma-ok
- **WHEN** a `Shape` variable holding a `Rectangle` is asserted to `Circle` using comma-ok syntax
- **THEN** `ok` is `false` and the zero value of `Circle` is returned

### Requirement: Type switch
The exercise SHALL include a function `DescribeShape(s Shape) string` that uses a type switch to return different descriptions based on the concrete type.

#### Scenario: Type switch on Circle
- **WHEN** `DescribeShape` is called with a `Circle{Radius: 5}`
- **THEN** it returns `"circle with radius 5.00"`

#### Scenario: Type switch on Rectangle
- **WHEN** `DescribeShape` is called with a `Rectangle{Width: 3, Height: 4}`
- **THEN** it returns `"rectangle 3.00x4.00"`

#### Scenario: Type switch with unknown type
- **WHEN** `DescribeShape` is called with an unknown Shape implementation
- **THEN** it returns `"unknown shape"`

### Requirement: Nil interface trap
The exercise SHALL include a test demonstrating that an interface holding a nil pointer is NOT itself nil. This is a critical Go gotcha.

#### Scenario: Nil pointer in interface is not nil
- **WHEN** a `*Circle` nil pointer is assigned to a `Shape` variable
- **THEN** the `Shape` variable is NOT nil (because it holds type information)
- **THEN** a directly nil `Shape` variable (no assignment) IS nil

### Requirement: Interface composition challenge
The exercise SHALL include a commented-out challenge section where participants define a composed interface `ShapeDescriber` that combines `Shape` and `Describer`, then uncomment tests verifying that `Circle` and `Rectangle` satisfy the composed interface.

#### Scenario: Composed interface
- **WHEN** `ShapeDescriber` is defined as combining `Shape` and `Describer`
- **THEN** `Circle` and `Rectangle` both satisfy `ShapeDescriber` because they implement all methods from both interfaces
