## 1. Exercise 02 — Structs & Methods

- [x] 1.1 Create `exercises/02_structs_and_methods/structs.go` with: `BankAccount` struct (exported `Owner`, unexported `balance`), factory function `NewBankAccount`, methods `Deposit`, `Withdraw`, `Balance`, `Summary`, `ResetBalance` (value receiver), `Address` struct with `FullAddress` method, `Person` struct with embedded `Address` — all function/method bodies using `panic("TODO: implement")`
- [x] 1.2 Create `exercises/02_structs_and_methods/structs_test.go` with tests ordered warm-up → core → think-about-it, covering struct literals, zero-valued struct, factory function, deposit/withdraw, balance getter, value vs pointer receiver mutation, summary, embedding with promoted fields and methods, and direct access to embedded struct

## 2. Verification

- [x] 2.1 Run `task test -- 02` and verify all tests fail with "TODO: implement" panics (not compile errors)
- [x] 2.2 Implement all functions temporarily, verify all tests pass, then revert to `panic("TODO: implement")`
