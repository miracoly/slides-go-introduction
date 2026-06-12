---
title: Go for Java Developers
theme: white
transition: none
highlight-theme: github
revealOptions:
  controls: true
  progress: true
  slideNumber: true
---

# Go for Java Developers

and maybe for Designers and Journalists

---

## Reading Code is way more difficult than writing it

- Most languages optimize for **writing** code
- Go optimizes for **reading** code
- Simplicity is a feature, not a limitation

---

## No Magic — Explicit Wiring

```go
func main() {
    db := connectDB(os.Getenv("DB_URL"))
    repo := NewUserRepo(db)
    svc := NewUserService(repo)

    http.HandleFunc("/users", svc.List)
    http.ListenAndServe(":8080", nil)
}
```

- Every dependency is created and passed **explicitly**
- No annotations, no injection, no hidden lifecycle
- What you read is what runs

---

## Go's Philosophy

- **Explicit over implicit** — no hidden behavior, no magic defaults
- **Composition over inheritance** — small interfaces, embedded structs
- **Errors are values** — handled where they occur, not thrown across layers
- **One way to do it** — less choice, less debate, more consistency

---

## Lightweight Binaries

| | Spring Boot + JVM | Go |
|---|---|---|
| Binary / JAR | 50–150 MB | 5–15 MB |
| Container image | 300–500 MB | 5–20 MB |
| Memory at idle | ~200 MB | ~10 MB |

- In Kubernetes with 10 microservices across 5 feature namespaces = **50 instances**
- That's **15 GB** idle memory (Spring) vs **500 MB** (Go)
- Faster rollouts, lower cost, smaller attack surface

---

## Developer Experience

- **Compiles in seconds** — not minutes
- **Single binary** — no runtime, no classpath, just copy & run
- **`go fmt`** — one code style, zero debates
- **Built-in tooling** — testing, benchmarks, profiling included
- **Goroutines** — lightweight concurrency, built into the language

---

## Batteries Included

| What you need | Go stdlib |
|---|---|
| HTTP server | `net/http` |
| JSON handling | `encoding/json` |
| SQL database | `database/sql` |
| Testing | `testing` |
| Logging | `log/slog` |

---

## Packages

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Sqrt(16))
}
```

- Every file belongs to a **package** — like Java packages, but simpler
- `main` package + `main` function = entry point

---

## Variables & Zero Values

```go
// Type inference with :=
name := "Alice"
age := 30

// Zero values — no null, no uninitialized memory
var count int     // 0
var label string  // ""
var ok    bool    // false
```

- Every type has a useful **zero value** — no surprises

---

## Constants

```go
const Pi = 3.14159
const MaxRetries = 3
const Greeting = "Hello"
```

- Immutable values, known at **compile time**
- No `final` keyword — just `const`

---

## Loops — Just `for`

```go
// Classic loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
// While-style
for count > 0 {
    count--
}
// Infinite
for {
    break
}
```

- `for` is the **only** loop keyword — no `while`, no `do`

---

## Switch

```go
switch day {
case "Monday":
    fmt.Println("back to work")
case "Saturday", "Sunday":
    fmt.Println("weekend!")
default:
    fmt.Println("midweek")
}
```

- No `break` needed — cases don't fall through by default
- Multiple values per case with commas

---

## Enums with `iota`

```go
type Weekday int

const (
    Monday Weekday = iota // 0
    Tuesday // 1
    Wednesday // 2
    Thursday // 3
    Friday // 4
)
```

- Go has no `enum` keyword — use a custom type + `iota`
- `iota` auto-increments within a `const` block

---

## Functions

```go
func Add(a int, b int) int {        // exported
    return a + b
}

func greet(name string) string {    // unexported
    return "Hello, " + name
}
```

- `func` keyword, parameters are `name type`, return type at the end
- No `public`/`private` keywords — capitalized names are exported

---

## Multiple Return Values

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := Divide(10, 3)
if err != nil {
    log.Fatal(err)
}
```

- Functions return **both** the result and an error — no exceptions

---

## Structs

```go
type BankAccount struct {
    Owner   string  // exported (uppercase)
    balance int     // unexported (lowercase)
}

// Factory function — Go has no constructors
func NewBankAccount(owner string, balance int) *BankAccount {
    return &BankAccount{Owner: owner, balance: balance}
}
```

- Visibility is controlled by **capitalization**, not keywords
- Composition via **embedding**, not inheritance

---

## Pass by Value

```go
func Reset(a BankAccount) {
    a.balance = 0
}

acc := BankAccount{Owner: "Alice", balance: 100}
Reset(acc)
// acc.balance is still 100 — Reset got a copy!
```

- Go passes **everything by value** — structs are copied
- The function works on its own copy, the original is untouched

---

## Pointers

```go
func Reset(a *BankAccount) {
    a.balance = 0
}

acc := BankAccount{Owner: "Alice", balance: 100}
Reset(&acc)
// acc.balance is now 0
```

- `*BankAccount` is a pointer — the function modifies the **original**
- `&acc` passes the address explicitly
- You always know when something can be modified

---

## Methods — Value Receiver

```go
func (a BankAccount) Summary() string {
    return fmt.Sprintf("%s: %d", a.Owner, a.balance)
}

acc := NewBankAccount("Alice", 100)
fmt.Println(acc.Summary())  // Alice: 100
```

- `(a BankAccount)` — the method receives a **copy**
- Use for methods that only **read** state

---

## Methods — Pointer Receiver

```go
func (a *BankAccount) Deposit(amount int) {
    a.balance += amount
}

acc := NewBankAccount("Alice", 100)
acc.Deposit(50)
fmt.Println(acc.Summary())  // Alice: 150
```

- `(a *BankAccount)` — the method receives the **original**
- Use for methods that **mutate** state

---

## Interfaces

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct{ Radius float64 }

func PrintArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

// won't compile, yet
PrintArea(Circle{Radius: 5})
```

- An interface defines a **contract** — a set of methods
- Any type that has those methods can be passed in

---

## Satisfying an Interface

```go
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Circle satisfies Shape — no "implements" needed
PrintArea(Circle{Radius: 5})
```

- Interfaces are satisfied **implicitly** — just implement the methods

---

## Why Implicit?

- The **consumer** defines the interface, the implementer doesn't need to know it exists
- Interfaces can be defined **after the fact** — even stdlib types already satisfy your interfaces
- Encourages **small interfaces** — 1–2 methods (`io.Reader`, `fmt.Stringer`)

---

## Type Switch

```go
func Describe(s Shape) string {
    switch v := s.(type) {
    case Circle:
        return fmt.Sprintf("Circle r=%.1f", v.Radius)
    case Rectangle:
        return fmt.Sprintf("Rect %.1fx%.1f", v.Width, v.Height)
    default:
        return "unknown"
    }
}
```

- Pattern matching on the **concrete type** behind an interface

---

## Exercise Time — Part 1

### Exercises 1, 2, 3

<https://github.com/miracoly/slides-go-introduction>

---

## Part 2

Error Handling, Collections, Concurrency

---

## Errors Are Values

```go
age, err := strconv.Atoi(input)
if err != nil {
    return 0, errors.New("invalid age")
}
```

- Errors are **returned**, not thrown
- Handled right where they occur — no hidden control flow

---

## Error Wrapping with `fmt.Errorf`

```go
func ParseUserAge(input string) (int, error) {
    age, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("parsing user age: %w", err)
    }
    return age, nil
}
// error: "parsing user age: strconv.Atoi: parsing "abc": invalid syntax"
```

- `%w` wraps the original error — adds **context** without losing the cause

---

## Sentinel Errors

```go
var ErrEmpty = errors.New("empty input")
var ErrOutOfRange = errors.New("out of range")

func ValidateAge(age int) error {
    if age < 0 || age > 150 {
        return ErrOutOfRange
    }
    return nil
}

if errors.Is(err, ErrOutOfRange) { ... }
```

- Package-level variables for **known, expected conditions**
- Callers check with `errors.Is` — works through wrapped errors too

---

## The `error` Interface

```go
// Built-in interface — no import needed
type error interface {
    Error() string
}
```

- Any type with an `Error() string` method is an `error`
- Just another interface — nothing special

---

## Custom Error Types

```go
type ParseError struct {
    Input   string
    Message string
}

func (e *ParseError) Error() string {
    return fmt.Sprintf("parse error: '%s': %s", e.Input, e.Message)
}
```

- `ParseError` satisfies `error` — it has the method
- Use when errors carry **structured data** the caller needs

---

## Using Custom Error Types

```go
func ParseMonth(s string) (int, error) {
    month, err := strconv.Atoi(s)
    if err != nil {
        return 0, &ParseError{Input: s, Message: "not a number"}
    }
    if month < 1 || month > 12 {
        return 0, &ParseError{Input: s, Message: "out of range"}
    }
    return month, nil
}
```

- Return `&ParseError{...}` — a pointer that satisfies `error`
- Callers can extract it with `errors.As` to access the fields

---

## Slices

```go
names := []string{"Alice", "Bob"}       // literal
names = append(names, "Charlie")

scores := make([]int, 0, 10)            // make(type, length, capacity)
scores = append(scores, 42)
```

- `make` creates a slice with a given length and optional capacity
- A slice is a **header** (pointer, length, capacity) — not the data itself

---

## Why Reassign `append`?

```go
names = append(names, "Diana")
//   ^ must reassign!
```

- If the underlying array is full, `append` allocates a **new, larger array**
- The returned slice header points to the new array
- Without reassignment, you'd keep the old, stale header

---

## Iterating with `range`

```go
for i, name := range names {
    fmt.Println(i, name)
}

for _, name := range names {  // ignore index with _
    fmt.Println(name)
}
```

- `range` returns **index and value** on each iteration
- Use `_` to discard values you don't need

---

## Slicing

```go
names := []string{"Alice", "Bob", "Charlie", "Diana"}

first2 := names[0:2]   // ["Alice", "Bob"]
last2 := names[2:]     // ["Charlie", "Diana"]
middle := names[1:3]   // ["Bob", "Charlie"]
```

- `s[low:high]` — from `low` (inclusive) to `high` (exclusive)
- **No copy** — the sub-slice shares the same underlying array
- Use `slices.Clone(s)` if you need an independent copy

---

## Deleting from a Slice

```go
import "slices"

names := []string{"Alice", "Bob", "Charlie"}

// Delete "Bob" (index 1)
names = slices.Delete(names, 1, 2)
// ["Alice", "Charlie"]
```

- `slices.Delete(s, i, j)` removes elements from index `i` to `j`
- From the `slices` package in the standard library

---

## Maps

```go
counts := make(map[string]int)
counts["hello"] = 1  // set
fmt.Println(counts["hello"])  // get → 1
fmt.Println(counts["missing"])  // get → 0 (zero value)

counts["new"]++  // 0 + 1 → works without init!
```

- Built-in hash map with `map[K]V` syntax
- Missing keys return the **zero value** — no `containsKey` needed

---

## Maps — Comma-Ok & Iteration

```go
// Comma-ok — check if a key exists
val, ok := counts["missing"]
// val = 0, ok = false

// Iterate — order is NOT guaranteed
for key, val := range counts {
    fmt.Println(key, val)
}
```

- **Comma-ok** distinguishes "missing key" from "zero value of 0"
- Map iteration order is **randomized** — unlike Java's `LinkedHashMap`

---

## Deleting from a Map

```go
counts := map[string]int{"a": 1, "b": 2, "c": 3}

delete(counts, "b")
// counts is now {"a": 1, "c": 3}
```

- `delete(map, key)` — built-in function, no import needed
- Deleting a non-existent key is a **no-op** — no error, no panic

---

## Maps as Sets

```go
seen := make(map[string]bool)

for _, word := range words {
    if !seen[word] {
        seen[word] = true
        result = append(result, word)
    }
}
```

- Go has no `Set` type — use `map[T]bool` instead
- `seen[word]` returns `false` (zero value) if key doesn't exist

---

## Built-in Concurrency

```go
go doWork()          // launch a goroutine — that's it
ch <- result         // send result through a channel
```

- `go` doesn't create an OS thread — it creates a **goroutine** (~2 KB vs ~1 MB per OS thread)
- The Go runtime **multiplexes** thousands of goroutines onto a few OS threads
- Context switching in **user space**, not via the kernel — orders of magnitude cheaper
- Channels replace shared memory

---

## Goroutines & Channels

```go
ch := make(chan string)  // unbuffered channel

go func() {
    var result string
    // other work ...
    ch <- result    // send — blocks until received
    fmt.Println("result sent")
}()

msg := <-ch  // receive — blocks until sent
```

- `go` launches a lightweight concurrent function
- Channels are typed pipes for **safe communication**
- Unbuffered channels **block on every send** until someone receives

---

## Directional Channels

```go
func Produce(ch chan<- int, values ...int) {
    for _, v := range values {
        ch <- v
    }
}

func Consume(ch <-chan int) int {
    return <-ch
}
```

- `chan<- int` — **send-only**, can't read from it
- `<-chan int` — **receive-only**, can't write to it
- The compiler enforces the direction — mistakes are caught at build time

---

## Buffered Channels

```go
ch := make(chan string, 3)           // buffer capacity: 3

ch <- "a"                            // doesn't block (buffer has space)
ch <- "b"                            // doesn't block
ch <- "c"                            // doesn't block
ch <- "d"                            // BLOCKS — buffer full
```

- Buffered channels only block when the **buffer is full**
- Useful when producer and consumer run at different speeds

---

## Closing Channels

```go
ch := make(chan int)

go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)                        // signal: no more values
}()
```

- `close(ch)` tells receivers that **no more values will be sent**
- Sending to a closed channel **panics** — only the sender should close

---

## Range over Channels

```go
for val := range ch {
    fmt.Println(val)                 // 1, 2, 3 — then stops
}
```

- `range` reads from the channel until it's **closed**
- Without close, `range` blocks forever — your program hangs

---

## Select

```go
select {
case order := <-highPriority:
    process(order)
case order := <-lowPriority:
    process(order)
case <-time.After(5 * time.Second):
    fmt.Println("no orders, idling")
}
```

- `select` waits on **multiple channels** simultaneously
- First ready channel wins — if both are ready, one is picked at random
- `time.After` returns a channel that fires after the duration

---

## Defer

```go
func ReadFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()    // runs when the function returns

    return io.ReadAll(f)
}
```

- `defer` schedules a call to run when the **enclosing function exits**
- Guarantees cleanup — even if the function returns early or panics

---

## WaitGroup — Wait for N Goroutines

```go
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Go(func() {
        doWork()
    })
}

wg.Wait()  // blocks until all goroutines finish
```

- `wg.Go` launches a goroutine and tracks it automatically
- `Wait()` blocks until all goroutines launched with `Go` are done

---

## Exercise Time — Part 2

### Exercises 4, 5, 6

<https://github.com/miracoly/slides-go-introduction>
