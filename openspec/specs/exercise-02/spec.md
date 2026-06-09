## ADDED Requirements

### Requirement: Struct definition and literals
The exercise SHALL include a struct `BankAccount` with fields `Owner` (string, exported) and `balance` (int, unexported). Participants MUST create instances using struct literals.

#### Scenario: Create a BankAccount with struct literal
- **WHEN** a `BankAccount` is created with `Owner: "Alice"` and `balance: 100`
- **THEN** the struct fields hold those values

#### Scenario: Zero-valued struct
- **WHEN** a `BankAccount` is declared with `var` (no initialization)
- **THEN** `Owner` is `""` and `balance` is `0`

### Requirement: Factory function (constructor replacement)
The exercise SHALL include a `NewBankAccount(owner string, initialBalance int) *BankAccount` factory function, demonstrating Go's convention for constructors.

#### Scenario: Create via factory function
- **WHEN** `NewBankAccount("Bob", 500)` is called
- **THEN** it returns a `*BankAccount` with `Owner` equal to `"Bob"` and `balance` equal to `500`

#### Scenario: Factory returns a pointer
- **WHEN** `NewBankAccount("Bob", 500)` is called
- **THEN** the return type is `*BankAccount` (pointer), not `BankAccount` (value)

### Requirement: Methods with pointer receiver
The exercise SHALL include `Deposit(amount int)` and `Withdraw(amount int) error` methods on `*BankAccount`, demonstrating that pointer receivers can mutate the struct.

#### Scenario: Deposit increases balance
- **WHEN** `Deposit(50)` is called on a `*BankAccount` with balance `100`
- **THEN** the balance becomes `150`

#### Scenario: Withdraw decreases balance
- **WHEN** `Withdraw(30)` is called on a `*BankAccount` with balance `100`
- **THEN** the balance becomes `70` and the returned error is `nil`

#### Scenario: Withdraw with insufficient funds
- **WHEN** `Withdraw(200)` is called on a `*BankAccount` with balance `100`
- **THEN** the balance remains `100` and a non-nil error with message `"insufficient funds"` is returned

### Requirement: Exported getter for unexported field
The exercise SHALL include a `Balance() int` method that returns the unexported `balance` field, demonstrating Go's visibility model.

#### Scenario: Balance getter
- **WHEN** `Balance()` is called on a `*BankAccount` with balance `100`
- **THEN** it returns `100`

### Requirement: Value vs pointer receiver behavior
The exercise SHALL include a value-receiver method `Summary() string` on `BankAccount` (not pointer) that returns a string like `"Alice: 100"`. A separate test SHALL demonstrate that a value-receiver method that attempts to modify state does NOT affect the original struct.

#### Scenario: Summary returns formatted string
- **WHEN** `Summary()` is called on a `BankAccount` with Owner `"Alice"` and balance `100`
- **THEN** it returns `"Alice: 100"`

#### Scenario: Value receiver does not mutate original
- **WHEN** a value-receiver method `ResetBalance()` is defined on `BankAccount` and sets `balance` to `0`
- **THEN** calling `ResetBalance()` on a `BankAccount` with balance `100` leaves the original balance at `100`

### Requirement: Struct embedding (composition)
The exercise SHALL include an `Address` struct (with `Street` and `City` fields) embedded into a `Person` struct (with `Name` field). This demonstrates Go's composition model as a replacement for Java inheritance.

#### Scenario: Access promoted fields
- **WHEN** a `Person` is created with an embedded `Address`
- **THEN** `Street` and `City` are accessible directly on `Person` (promoted fields)

#### Scenario: Method promotion
- **WHEN** `Address` has a `FullAddress() string` method returning `"Street, City"`
- **THEN** calling `FullAddress()` on `Person` works via method promotion

### Requirement: Embedding is not inheritance
The exercise SHALL include a test demonstrating that embedding does NOT create an is-a relationship — the embedded struct is accessible as a named field.

#### Scenario: Access embedded struct directly
- **WHEN** a `Person` has an embedded `Address`
- **THEN** the `Address` field is accessible as `person.Address` (the embedded struct itself)
