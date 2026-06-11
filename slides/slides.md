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

---

## Reading Code is more difficult [than writing it]

*— Fabrice Bellard, creator of FFmpeg*

- Most languages optimize for **writing** code
- Go optimizes for **reading** code
- Simplicity is a feature, not a limitation

---

## No Magic

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

## Developer Experience

- **Compiles in seconds** — not minutes
- **Single binary** — no runtime, no classpath, just copy & run
- **`go fmt`** — one code style, zero debates
- **Built-in tooling** — testing, benchmarks, profiling included

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

## Pointers — Explicit, Not Hidden

```go
func Increment(p *int) {
    *p++        // dereference to modify
}

x := 42
Increment(&x)  // pass address explicitly
// x is now 43
```

- `*int` is a pointer to an int, `&x` takes the address
- You always know when something can be modified

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

## Methods

```go
// Pointer receiver — modifies the original
func (a *BankAccount) Deposit(amount int) {
    a.balance += amount
}

// Value receiver — works on a copy
func (a BankAccount) Summary() string {
    return fmt.Sprintf("%s: %d", a.Owner, a.balance)
}
```

- Pointer receiver (`*T`): use when the method **mutates** state
- Value receiver (`T`): use when it only **reads** state

---

## Interfaces

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
// Circle satisfies Shape — no "implements" needed
```

- Interfaces are satisfied **implicitly** — just implement the methods

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

<https://github.com/marc-uhlig/go-introduction>

---

## Part 2

Error Handling, Collections, Concurrency

---

## Errors Are Values

```go
age, err := strconv.Atoi(input)
if err != nil {
    return fmt.Errorf("invalid age: %w", err)
}
```

- Errors are **returned**, not thrown
- Handled right where they occur — no hidden control flow
- `%w` wraps the original error for later inspection

---

## Sentinel Errors & Custom Types

```go
var ErrNotFound = errors.New("not found")

type ParseError struct {
    Input   string
    Message string
}

func (e *ParseError) Error() string {
    return fmt.Sprintf("parse error: '%s': %s", e.Input, e.Message)
}

// Check wrapped errors
if errors.Is(err, ErrNotFound) { ... }
```

- Sentinel errors for **known conditions**, custom types for **structured data**

---

## Slices

```go
names := []string{"Alice", "Bob"}
names = append(names, "Charlie")

for i, name := range names {
    fmt.Println(i, name)
}
```

- Dynamic arrays — grow with `append`
- `range` iterates with index and value

---

## Maps

```go
counts := map[string]int{}
counts["hello"]++  // zero value (0) + 1

// Comma-ok pattern
val, ok := counts["missing"]
// val = 0, ok = false
```

- Built-in hash map with `map[K]V` syntax
- **Comma-ok** distinguishes "missing key" from "zero value"

---

## Goroutines & Channels

```go
ch := make(chan string)

go func() {
    ch <- "hello from goroutine"
}()

msg := <-ch  // blocks until a message arrives
```

- `go` launches a lightweight concurrent function
- Channels are typed pipes for **safe communication**

---

## Select

```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case <-time.After(3 * time.Second):
    fmt.Println("timeout!")
}
```

- `select` waits on **multiple channels** simultaneously
- First ready channel wins — like a switch for concurrency

---

## Exercise Time — Part 2

### Exercises 4, 5, 6

<https://github.com/marc-uhlig/go-introduction>
