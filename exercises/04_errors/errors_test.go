package errors

import (
	// Add "fmt" to this import block when you reach the challenge.
	// Uncomment "errors" and "testing" when you uncomment your first test.
	// "errors"
	// "testing"
)

// ============================================================================
// Step 1: Basic error returns
//
// Implement ParseAge in errors.go, then uncomment the tests below.
// ============================================================================

// // The (value, error) return pattern is the bread and butter of Go.
// // Every function that can fail returns an error as the LAST return value.
// // The caller checks it immediately — no invisible exception flow.
// func TestParseAgeValid(t *testing.T) {
// 	age, err := ParseAge("25")
// 	if err != nil {
// 		t.Fatalf("ParseAge(\"25\"): unexpected error: %v", err)
// 	}
// 	if age != 25 {
// 		t.Errorf("ParseAge(\"25\"): got %d, want 25", age)
// 	}
// }

// func TestParseAgeEmpty(t *testing.T) {
// 	_, err := ParseAge("")
// 	if err == nil {
// 		t.Fatal("ParseAge(\"\"): expected error, got nil")
// 	}
// }

// func TestParseAgeNonNumeric(t *testing.T) {
// 	_, err := ParseAge("abc")
// 	if err == nil {
// 		t.Fatal("ParseAge(\"abc\"): expected error, got nil")
// 	}
// 	// In Go, error messages are lowercase by convention (unlike Java).
// 	// Check that the error message contains useful context.
// 	got := err.Error()
// 	if got == "" {
// 		t.Error("error message should not be empty")
// 	}
// }

// func TestParseAgeNegative(t *testing.T) {
// 	_, err := ParseAge("-5")
// 	if err == nil {
// 		t.Fatal("ParseAge(\"-5\"): expected error, got nil")
// 	}
// }

// ============================================================================
// Step 2: Sentinel errors and errors.Is
//
// After Step 1 passes, define ErrEmpty, ErrOutOfRange, and ValidateAge
// in errors.go. Then uncomment the tests below.
// NOTE: Add "errors" to the import block at the top of this file.
// ============================================================================

// // In Java:
// //   catch (EmptyInputException e) { ... }
// //
// // In Go:
// //   if errors.Is(err, ErrEmpty) { ... }
// //
// // Sentinel errors are package-level variables. They're comparable values,
// // not types. You check identity, not class hierarchy.
// func TestParseAgeEmptySentinel(t *testing.T) {
// 	_, err := ParseAge("")
// 	if !errors.Is(err, ErrEmpty) {
// 		t.Errorf("ParseAge(\"\"): error should be ErrEmpty, got: %v", err)
// 	}
// }

// func TestValidateAgeOK(t *testing.T) {
// 	err := ValidateAge(25)
// 	if err != nil {
// 		t.Errorf("ValidateAge(25): unexpected error: %v", err)
// 	}
// }

// func TestValidateAgeNegative(t *testing.T) {
// 	err := ValidateAge(-1)
// 	if !errors.Is(err, ErrOutOfRange) {
// 		t.Errorf("ValidateAge(-1): expected ErrOutOfRange, got: %v", err)
// 	}
// }

// func TestValidateAgeTooHigh(t *testing.T) {
// 	err := ValidateAge(200)
// 	if !errors.Is(err, ErrOutOfRange) {
// 		t.Errorf("ValidateAge(200): expected ErrOutOfRange, got: %v", err)
// 	}
// }

// // Sentinels are unique values. They must not be equal to each other.
// func TestSentinelIdentity(t *testing.T) {
// 	if errors.Is(ErrEmpty, ErrOutOfRange) {
// 		t.Error("ErrEmpty and ErrOutOfRange should not be equal")
// 	}
// }

// ============================================================================
// Step 3: Custom error types
//
// After Step 2 passes, define ParseError struct with Error() method
// and ParseMonth function in errors.go. Then uncomment the tests below.
// ============================================================================

// // In Java, custom exceptions carry extra data via fields.
// // In Go, a custom error is just a struct with an Error() method.
// // The error interface is satisfied implicitly — same as exercise 03.
// func TestParseErrorMessage(t *testing.T) {
// 	pe := &ParseError{Input: "abc", Message: "not a number"}
// 	want := "parse error: 'abc': not a number"
// 	got := pe.Error()
// 	if got != want {
// 		t.Errorf("ParseError.Error(): got %q, want %q", got, want)
// 	}
// }

// // A *ParseError is assignable to the error interface.
// // No "implements" needed — just having Error() string is enough.
// func TestParseErrorSatisfiesError(t *testing.T) {
// 	var err error = &ParseError{Input: "test", Message: "test error"}
// 	if err == nil {
// 		t.Error("*ParseError assigned to error should not be nil")
// 	}
// }

// func TestParseMonthValid(t *testing.T) {
// 	month, err := ParseMonth("6")
// 	if err != nil {
// 		t.Fatalf("ParseMonth(\"6\"): unexpected error: %v", err)
// 	}
// 	if month != 6 {
// 		t.Errorf("ParseMonth(\"6\"): got %d, want 6", month)
// 	}
// }

// func TestParseMonthNonNumeric(t *testing.T) {
// 	_, err := ParseMonth("june")
// 	if err == nil {
// 		t.Fatal("ParseMonth(\"june\"): expected error, got nil")
// 	}
// 	// Type-assert to check that it's a *ParseError with the right Input.
// 	pe, ok := err.(*ParseError)
// 	if !ok {
// 		t.Fatalf("expected *ParseError, got %T", err)
// 	}
// 	if pe.Input != "june" {
// 		t.Errorf("ParseError.Input: got %q, want \"june\"", pe.Input)
// 	}
// }

// func TestParseMonthOutOfRange(t *testing.T) {
// 	_, err := ParseMonth("13")
// 	if err == nil {
// 		t.Fatal("ParseMonth(\"13\"): expected error, got nil")
// 	}
// 	pe, ok := err.(*ParseError)
// 	if !ok {
// 		t.Fatalf("expected *ParseError, got %T", err)
// 	}
// 	if pe.Input != "13" {
// 		t.Errorf("ParseError.Input: got %q, want \"13\"", pe.Input)
// 	}
// }

// func TestParseMonthEmpty(t *testing.T) {
// 	_, err := ParseMonth("")
// 	if !errors.Is(err, ErrEmpty) {
// 		t.Errorf("ParseMonth(\"\"): expected ErrEmpty, got: %v", err)
// 	}
// }

// ============================================================================
// Step 4: Error wrapping with %w
//
// After Step 3 passes, implement ParseUserAge in errors.go,
// then uncomment the tests below.
// ============================================================================

// // In Java:
// //   throw new ServiceException("user age failed", cause);
// //   // later: getCause() walks the chain
// //
// // In Go:
// //   fmt.Errorf("context: %w", err)
// //   // later: errors.Is(err, target) walks the chain
// //
// // The %w verb is special — it WRAPS the error, preserving it in a chain.
// // Using %v instead would just format the string, losing the chain.
// func TestParseUserAgeValid(t *testing.T) {
// 	age, err := ParseUserAge("30")
// 	if err != nil {
// 		t.Fatalf("ParseUserAge(\"30\"): unexpected error: %v", err)
// 	}
// 	if age != 30 {
// 		t.Errorf("ParseUserAge(\"30\"): got %d, want 30", age)
// 	}
// }

// // errors.Is walks the ENTIRE wrap chain. Even though ParseUserAge wraps
// // the error, the original ErrEmpty is still findable.
// func TestParseUserAgePreservesSentinel(t *testing.T) {
// 	_, err := ParseUserAge("")
// 	if !errors.Is(err, ErrEmpty) {
// 		t.Errorf("ParseUserAge(\"\"): expected errors.Is(err, ErrEmpty), got: %v", err)
// 	}
// }

// func TestParseUserAgeWrapsMessage(t *testing.T) {
// 	_, err := ParseUserAge("abc")
// 	if err == nil {
// 		t.Fatal("ParseUserAge(\"abc\"): expected error, got nil")
// 	}
// 	got := err.Error()
// 	want := "parsing user age:"
// 	if len(got) < len(want) || got[:len(want)] != want {
// 		t.Errorf("error should start with %q, got %q", want, got)
// 	}
// }

// ============================================================================
// CHALLENGE: errors.As — extracting custom types from wrapped chains
//
// After Step 4 passes, implement ParseBirthMonth in errors.go,
// then uncomment the tests below.
// NOTE: Add "fmt" to the import block at the top of this file.
// ============================================================================

// // errors.As is like Java's instanceof + cast, but for error chains:
// //
// //   Java:  if (e instanceof ParseException pe) { pe.getInput(); }
// //   Go:    var pe *ParseError
// //          if errors.As(err, &pe) { pe.Input }
// //
// // errors.As walks the wrap chain looking for an error that can be
// // assigned to the target type. If found, it sets the target and returns true.
// func TestParseBirthMonthErrorsAs(t *testing.T) {
// 	_, err := ParseBirthMonth("june")
// 	if err == nil {
// 		t.Fatal("ParseBirthMonth(\"june\"): expected error, got nil")
// 	}
//
// 	// errors.As extracts the *ParseError even through wrapping.
// 	var pe *ParseError
// 	if !errors.As(err, &pe) {
// 		t.Fatalf("errors.As should find *ParseError in chain, got: %v", err)
// 	}
// 	if pe.Input != "june" {
// 		t.Errorf("ParseError.Input: got %q, want \"june\"", pe.Input)
// 	}
// }

// func TestErrorsAsReturnsFalseForNonMatch(t *testing.T) {
// 	err := fmt.Errorf("simple error")
//
// 	var pe *ParseError
// 	if errors.As(err, &pe) {
// 		t.Error("errors.As should return false for a non-*ParseError")
// 	}
// }
