// Package structsandmethods
//
// In this exercise, YOU define everything from scratch — structs, factory
// functions, and methods. Work through the steps in order. Each step tells
// you what to build. After implementing, uncomment the matching tests in
// structs_test.go and run: task test -- 02
package structsandmethods

// Add imports as you need them:
// - "errors" or "fmt" for creating error values (Step 3)
// - "fmt" for string formatting (Step 4)

// ============================================================================
// Step 1: Define a struct and factory function
//
// Uncomment Step 1 tests in structs_test.go, then implement below.
// ============================================================================
//
// In Java, this would be a class with private fields and a constructor:
//
//	public class BankAccount {
//	    private String owner;
//	    private int balance;
//	    public BankAccount(String owner, int balance) { ... }
//	}
//
// In Go, we use a struct. Visibility is controlled by capitalization:
//   - Uppercase first letter → exported (like Java's public)
//   - Lowercase first letter → unexported (like Java's package-private)
//
// YOUR TASK:
// 1. Define a struct called BankAccount with two fields:
//    - Owner string   (exported — uppercase O)
//    - balance int     (unexported — lowercase b)
// 2. Define a factory function called NewBankAccount that takes an owner
//    (string) and initialBalance (int) and returns *BankAccount.
//    Go has no constructors. By convention, factory functions named NewXxx
//    return a pointer to the initialized struct.

// ============================================================================
// Step 2: Pointer receiver methods
//
// After Step 1 tests pass, uncomment Step 2 tests in structs_test.go.
// ============================================================================
//
// In Java, all methods on objects implicitly operate on the reference (this).
// In Go, you must CHOOSE: pointer receiver (*T) to modify, value receiver (T)
// for read-only.
//
// YOUR TASK:
// 1. Define a method called Deposit on *BankAccount that takes an int (amount)
//    and adds it to the balance. Returns nothing.
//
// 2. Define a method called Balance on *BankAccount that returns the current
//    balance as an int.
//    In Java, you'd write: public int getBalance() { return this.balance; }
//    In Go, the convention is just Balance() — no "get" prefix.

// ============================================================================
// Step 3: Error handling in methods
//
// After Step 2 tests pass, uncomment Step 3 tests in structs_test.go.
// ============================================================================
//
// In Java, you'd throw an InsufficientFundsException.
// In Go, errors are values — return them instead of throwing.
//
// YOUR TASK:
// 1. Define a method called Withdraw on *BankAccount that takes an int (amount)
//    and returns an error.
//    - If amount > balance, return an error with message "insufficient funds".
//      Hint: use errors.New("insufficient funds") or fmt.Errorf(...)
//    - Otherwise, subtract amount from balance and return nil.
//    NOTE: Add "errors" or "fmt" to your imports.

// ============================================================================
// Step 4: Value receivers
//
// After Step 3 tests pass, uncomment Step 4 tests in structs_test.go.
// ============================================================================
//
// A value receiver (BankAccount, not *BankAccount) receives a COPY of the struct.
// It can read fields but any modifications only affect the copy, not the original.
//
// This is THE biggest "gotcha" for Java devs. In Java, all methods operate on
// the same object reference. In Go, a value receiver is like getting a photocopy
// of the object — you can scribble on it, but the original is untouched.
//
// YOUR TASK:
// 1. Define a method called Summary on BankAccount (VALUE receiver, no *)
//    that returns a string like "Alice: 100".
//    Hint: use fmt.Sprintf("%s: %d", a.Owner, a.balance)
//    NOTE: Add "fmt" to your imports if you haven't already.
//
// 2. Define a method called ResetBalance on BankAccount (VALUE receiver, no *)
//    that sets balance to 0.
//    This method will NOT affect the original struct — that's the point!
//    The test verifies that the original balance is unchanged after calling it.

// ============================================================================
// CHALLENGE: Struct Embedding (composition over inheritance)
//
// After Step 4 tests pass, uncomment the challenge tests in structs_test.go.
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
// 4. Uncomment the challenge tests in structs_test.go
// 5. Run the tests: task test -- 02
//
// Hints:
//   - Embedding syntax:  type Employee struct { Person; Company string }
//   - This promotes Person's fields (Name, Email) and methods (Greet) to Employee
//   - The embedded struct is also accessible as a field: employee.Person
