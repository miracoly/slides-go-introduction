// Package structsandmethods
package structsandmethods

// BankAccount represents a bank account.
// In Java, this would be a class with private fields and public getters/setters:
//
//	public class BankAccount {
//	    private String owner;
//	    private int balance;
//	    public BankAccount(String owner, int balance) { ... }
//	}
//
// In Go, we use a struct. Exported fields (uppercase) are like public.
// Unexported fields (lowercase) are like private — but scoped to the package, not the type.
type BankAccount struct {
	Owner   string // exported — accessible from any package (like Java's public)
	balance int    // unexported — accessible only within this package (like Java's package-private)
}

// NewBankAccount creates a new BankAccount with the given owner and initial balance.
// In Java, you'd use a constructor: new BankAccount("Alice", 100).
// In Go, there are no constructors. By convention, factory functions named
// NewXxx return a pointer to the initialized struct.
func NewBankAccount(owner string, initialBalance int) *BankAccount {
	panic("TODO: implement")
}

// Deposit adds amount to the account balance.
// This is a method with a POINTER receiver (*BankAccount).
// In Java, all methods on objects implicitly operate on the reference.
// In Go, you must choose: pointer receiver (*T) to modify, value receiver (T) for read-only.
func (a *BankAccount) Deposit(amount int) {
	panic("TODO: implement")
}

// Withdraw subtracts amount from the account balance.
// Returns an error if the balance is insufficient.
// Use errors.New() or fmt.Errorf() to create the error.
func (a *BankAccount) Withdraw(amount int) error {
	panic("TODO: implement")
}

// Balance returns the current balance.
// In Java, you'd write: public int getBalance() { return this.balance; }
// In Go, the convention is just Balance() — no "get" prefix.
func (a *BankAccount) Balance() int {
	panic("TODO: implement")
}

// Summary returns a string like "Alice: 100".
// This uses a VALUE receiver (BankAccount, not *BankAccount).
// It receives a COPY of the struct — it can read fields but cannot modify the original.
// use fmt.Sprintf
func (a BankAccount) Summary() string {
	panic("TODO: implement")
}

// ResetBalance sets balance to 0.
// BUT: this is a VALUE receiver! It operates on a copy.
// Calling this method will NOT affect the original struct.
// This is the key difference from Java, where all object methods share the same reference.
func (a BankAccount) ResetBalance() {
	panic("TODO: implement")
}

// ============================================================================
// CHALLENGE: Struct Embedding (composition over inheritance)
// ============================================================================
//
// In Java, you'd model this with inheritance:
//
//	class Person {
//	    String name;
//	    String email;
//	    String greet() { return "Hi, I'm " + name; }
//	}
//
//	class Employee extends Person {
//	    String company;
//	}
//
// In Go, there is NO inheritance. Instead, you EMBED a struct.
// Embedding promotes the embedded struct's fields and methods to the outer struct.
// This is composition — but it LOOKS like inheritance, which is exactly why
// Java devs need to understand the difference.
//
// YOUR TASK:
// 1. Define a Person struct with exported fields: Name (string), Email (string)
// 2. Add a Greet() method on Person that returns "Hi, I'm <Name>"
// 3. Define an Employee struct with an exported field: Company (string)
//    and an EMBEDDED Person (not a named field — just the type name)
// 4. Uncomment the three embedding tests in structs_test.go
// 5. Run the tests: task test -- 02
//
// Hints:
//   - Embedding syntax:  type Employee struct { Person; Company string }
//   - This promotes Person's fields (Name, Email) and methods (Greet) to Employee
//   - The embedded struct is also accessible as a field: employee.Person
