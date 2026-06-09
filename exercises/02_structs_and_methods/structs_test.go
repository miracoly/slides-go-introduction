package structsandmethods

import "testing"

// ============================================================================
// Warm-up: Creating structs
// ============================================================================

// In Java, you create objects with constructors: new BankAccount("Alice", 100).
// In Go, you create structs with LITERALS — there's no "new" keyword for structs.
//
// Struct literals use field names explicitly:
//
//	BankAccount{Owner: "Alice", balance: 100}
//
// This is more like Java's builder pattern, but built into the language.
func TestStructLiteral(t *testing.T) {
	acc := BankAccount{Owner: "Alice", balance: 100}

	if acc.Owner != "Alice" {
		t.Errorf("Owner: got %q, want \"Alice\"", acc.Owner)
	}
	if acc.balance != 100 {
		t.Errorf("balance: got %d, want 100", acc.balance)
	}
}

// In Java, uninitialized fields get defaults: null for objects, 0 for ints.
// In Go, it's the same idea but more consistent — ALL types have zero values.
// A zero-valued struct has all fields set to their zero values.
//
//	var acc BankAccount → Owner: "", balance: 0
func TestZeroValuedStruct(t *testing.T) {
	var acc BankAccount

	if acc.Owner != "" {
		t.Errorf("Owner zero value: got %q, want \"\"", acc.Owner)
	}
	if acc.balance != 0 {
		t.Errorf("balance zero value: got %d, want 0", acc.balance)
	}
}

// ============================================================================
// Core: Factory functions and methods
// ============================================================================

// In Java: BankAccount acc = new BankAccount("Bob", 500);
// In Go: acc := NewBankAccount("Bob", 500)
//
// Go has no constructors. The convention is a factory function named NewXxx
// that returns a *pointer* to the struct. Why a pointer? So callers can
// call pointer-receiver methods without explicitly taking the address.
func TestNewBankAccount(t *testing.T) {
	acc := NewBankAccount("Bob", 500)

	if acc == nil {
		t.Fatal("NewBankAccount returned nil")
	}
	if acc.Owner != "Bob" {
		t.Errorf("Owner: got %q, want \"Bob\"", acc.Owner)
	}
	if acc.balance != 500 {
		t.Errorf("balance: got %d, want 500", acc.balance)
	}
}

// In Java: acc.deposit(50); — the method implicitly operates on `this`.
// In Go: acc.Deposit(50) — the method has an explicit receiver: func (a *BankAccount) Deposit(...)
//
// The pointer receiver (*BankAccount) means Deposit modifies the original struct.
func TestDeposit(t *testing.T) {
	acc := NewBankAccount("Alice", 100)
	acc.Deposit(50)

	if acc.Balance() != 150 {
		t.Errorf("after Deposit(50): balance got %d, want 150", acc.Balance())
	}
}

func TestDepositMultiple(t *testing.T) {
	acc := NewBankAccount("Alice", 0)
	acc.Deposit(10)
	acc.Deposit(20)
	acc.Deposit(30)

	if acc.Balance() != 60 {
		t.Errorf("after 3 deposits: balance got %d, want 60", acc.Balance())
	}
}

// In Java, you'd throw an InsufficientFundsException.
// In Go, errors are values — Withdraw returns an error instead of throwing.
func TestWithdraw(t *testing.T) {
	acc := NewBankAccount("Alice", 100)
	err := acc.Withdraw(30)

	if err != nil {
		t.Fatalf("Withdraw(30): unexpected error: %v", err)
	}
	if acc.Balance() != 70 {
		t.Errorf("after Withdraw(30): balance got %d, want 70", acc.Balance())
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	acc := NewBankAccount("Alice", 100)
	err := acc.Withdraw(200)

	if err == nil {
		t.Fatal("Withdraw(200): expected error, got nil")
	}
	if err.Error() != "insufficient funds" {
		t.Errorf("error message: got %q, want \"insufficient funds\"", err.Error())
	}
	if acc.Balance() != 100 {
		t.Errorf("balance should be unchanged: got %d, want 100", acc.Balance())
	}
}

// In Java: public int getBalance() { return this.balance; }
// In Go: func (a *BankAccount) Balance() int — no "get" prefix.
//
// The balance field is unexported (lowercase), so code outside this package
// cannot access it directly. The Balance() method is the public API.
func TestBalance(t *testing.T) {
	acc := NewBankAccount("Alice", 42)

	if acc.Balance() != 42 {
		t.Errorf("Balance(): got %d, want 42", acc.Balance())
	}
}

// ============================================================================
// Think about it: Value vs pointer receivers
// ============================================================================

// Summary() uses a VALUE receiver: func (a BankAccount) Summary() string
// This means it receives a COPY of the struct. It can read fields but
// any modifications would only affect the copy, not the original.
//
// In Java, there's no equivalent distinction — all methods operate on the
// same object reference. This is one of the biggest "gotchas" for Java devs.
func TestSummary(t *testing.T) {
	acc := BankAccount{Owner: "Alice", balance: 100}
	got := acc.Summary()

	if got != "Alice: 100" {
		t.Errorf("Summary(): got %q, want \"Alice: 100\"", got)
	}
}

// THIS IS THE KEY TEST. In Java, calling a method on an object always
// modifies the same instance:
//
//	acc.resetBalance(); // acc.balance is now 0
//
// In Go, a VALUE receiver gets a COPY. Modifications don't affect the original.
// ResetBalance() sets balance to 0 on the copy, but the original is unchanged.
func TestValueReceiverDoesNotMutate(t *testing.T) {
	acc := BankAccount{Owner: "Alice", balance: 100}
	acc.ResetBalance() // this operates on a COPY

	if acc.balance != 100 {
		t.Errorf("balance after ResetBalance(): got %d, want 100 (value receiver should not mutate)", acc.balance)
	}
}

// ============================================================================
// CHALLENGE: Struct embedding (composition)
//
// Once you've defined Person, Greet(), and Employee in structs.go,
// uncomment the three tests below and run: task test -- 02
// ============================================================================

// // In Java, you'd use inheritance:
// //
// //	class Employee extends Person {
// //	    String company;
// //	}
// //
// // In Go, there is NO inheritance. Instead, you embed a struct.
// // Embedding PROMOTES the embedded struct's fields to the outer struct.
// // You can access them directly, as if they were declared on Employee.
// func TestEmbeddingPromotedFields(t *testing.T) {
// 	emp := Employee{
// 		Person: Person{
// 			Name:  "Alice",
// 			Email: "alice@example.com",
// 		},
// 		Company: "Acme Corp",
// 	}
//
// 	// Promoted fields — accessible directly on Employee
// 	if emp.Name != "Alice" {
// 		t.Errorf("Name: got %q, want \"Alice\"", emp.Name)
// 	}
// 	if emp.Email != "alice@example.com" {
// 		t.Errorf("Email: got %q, want \"alice@example.com\"", emp.Email)
// 	}
// 	if emp.Company != "Acme Corp" {
// 		t.Errorf("Company: got %q, want \"Acme Corp\"", emp.Company)
// 	}
// }

// // Embedding also promotes METHODS. Person has Greet() — and because
// // Person is embedded in Employee, you can call Greet() on Employee directly.
// //
// // In Java, this "just works" because Employee inherits from Person.
// // In Go, it also "just works" — but the mechanism is promotion, not inheritance.
// func TestEmbeddingPromotedMethods(t *testing.T) {
// 	emp := Employee{
// 		Person:  Person{Name: "Alice", Email: "alice@example.com"},
// 		Company: "Acme Corp",
// 	}
//
// 	got := emp.Greet()
// 	if got != "Hi, I'm Alice" {
// 		t.Errorf("Greet(): got %q, want \"Hi, I'm Alice\"", got)
// 	}
// }

// // Embedding is NOT inheritance. The embedded struct is still a named field.
// // You can access it directly: employee.Person
// //
// // This is the proof that it's composition, not inheritance:
// // the inner type is a real, accessible field — not a hidden base class.
// // In Java, you can't write employee.super as a field — but in Go you can.
// func TestEmbeddedStructAccessible(t *testing.T) {
// 	emp := Employee{
// 		Person:  Person{Name: "Alice", Email: "alice@example.com"},
// 		Company: "Acme Corp",
// 	}
//
// 	// The embedded struct is accessible as a field
// 	p := emp.Person
// 	if p.Name != "Alice" {
// 		t.Errorf("Person.Name: got %q, want \"Alice\"", p.Name)
// 	}
// 	if p.Email != "alice@example.com" {
// 		t.Errorf("Person.Email: got %q, want \"alice@example.com\"", p.Email)
// 	}
// }
