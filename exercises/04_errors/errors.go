// Package errors
//
// In this exercise, YOU define everything from scratch — sentinel errors,
// custom error types, and functions that return errors. Work through the
// steps in order. Each step tells you what to build. After implementing,
// uncomment the matching tests in errors_test.go and run: task test -- 04
package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ============================================================================
// Step 1: Basic error returns
//
// Uncomment Step 1 tests in errors_test.go, then implement below.
// ============================================================================
//
// In Java, you signal failure by throwing an exception:
//
//	public int parseAge(String s) throws IllegalArgumentException {
//	    int age = Integer.parseInt(s);       // throws NumberFormatException
//	    if (age < 0) throw new IllegalArgumentException("negative age");
//	    return age;
//	}
//
// The caller doesn't see the error in the return type. It's invisible.
// If they forget to catch it, it crashes at runtime.
//
// In Go, errors are VALUES returned alongside the result:
//
//	func ParseAge(s string) (int, error) {
//	    // return value, nil        → success
//	    // return 0, errors.New("") → failure
//	}
//
// The caller always sees the error. They can't accidentally ignore it
// (the compiler warns on unused variables).
//
// YOUR TASK:
// 1. Write a function called ParseAge that takes a string and returns
//    (int, error).
//    - If the string is empty, return 0 and an error with the message
//      "empty input". Hint: use errors.New("empty input")
//    - Use strconv.Atoi to convert the string to an int. If it fails,
//      return 0 and an error with a message containing "invalid age".
//      Hint: use fmt.Errorf("invalid age: %s", s)
//    - If the number is negative, return 0 and an error containing
//      "invalid age"
//    - Otherwise, return the parsed age and nil
// 2. Uncomment Step 1 tests, run: task test -- 04

// ============================================================================
// Step 2: Sentinel errors and errors.Is
//
// After Step 1 tests pass, uncomment Step 2 tests in errors_test.go.
// ============================================================================
//
// In Java, you catch specific exception types:
//
//	try { parseAge(input); }
//	catch (NumberFormatException e) { ... }
//	catch (IllegalArgumentException e) { ... }
//
// In Go, the equivalent is a "sentinel error" — a package-level variable
// that callers compare against:
//
//	if errors.Is(err, ErrOutOfRange) { ... }
//
// YOUR TASK:
// 1. Define two sentinel errors at package level:
//    - ErrEmpty: with the message "empty input"
//    - ErrOutOfRange: with the message "out of range"
//    Hint: use errors.New to create them as package-level var declarations
// 2. Update your ParseAge function to return ErrEmpty (instead of a new
//    error) when the input string is empty. This lets callers check with
//    errors.Is(err, ErrEmpty).
// 3. Write a function called ValidateAge that takes an int and returns
//    an error.
//    - If age is negative or greater than 150, return ErrOutOfRange
//    - Otherwise return nil
// 4. Uncomment Step 2 tests, run: task test -- 04

// ============================================================================
// Step 3: Custom error types
//
// After Step 2 tests pass, uncomment Step 3 tests in errors_test.go.
// ============================================================================
//
// In Java, you define custom exception classes:
//
//	public class ParseException extends Exception {
//	    private String input;
//	    public ParseException(String input, String message) {
//	        super(message);
//	        this.input = input;
//	    }
//	}
//
// In Go, the error interface is just:
//
//	type error interface { Error() string }
//
// Any struct with an Error() method satisfies it — just like exercise 03
// taught you. No "extends", no "implements".
//
// YOUR TASK:
// 1. Define a struct called ParseError with two exported fields:
//    Input (string) and Message (string)
// 2. Add an Error method on *ParseError that returns a string formatted as:
//    "parse error: '<input>': <message>"
//    For example: "parse error: 'abc': not a number"
// 3. Write a function called ParseMonth that takes a string and returns
//    (int, error).
//    - If the string is empty, return 0 and ErrEmpty
//    - Use strconv.Atoi to parse it. If that fails, return 0 and a
//      *ParseError with the input string and a message like "not a number"
//    - If the number is less than 1 or greater than 12, return 0 and a
//      *ParseError with the input and a message like "month out of range"
//    - Otherwise, return the parsed month and nil
// 4. Uncomment Step 3 tests, run: task test -- 04

// ============================================================================
// Step 4: Error wrapping with %w
//
// After Step 3 tests pass, uncomment Step 4 tests in errors_test.go.
// ============================================================================
//
// In Java, you wrap exceptions with the cause chain:
//
//	throw new ServiceException("user age failed", originalException);
//	// later: getCause() to walk the chain
//
// In Go, you wrap errors with fmt.Errorf and the %w verb:
//
//	return fmt.Errorf("context: %w", originalErr)
//
// The %w verb (not %v!) preserves the original error in a chain.
// errors.Is can then find it through any number of wraps:
//
//	errors.Is(wrappedErr, ErrEmpty) → true (even through wrapping)
//
// YOUR TASK:
// 1. Write a function called ParseUserAge that takes a string and returns
//    (int, error).
//    - Call your ParseAge function.
//    - If ParseAge returns an error, wrap it with additional context using
//      fmt.Errorf("parsing user age: %w", err)
//    - If successful, return the age and nil
// 2. Uncomment Step 4 tests, run: task test -- 04

// ============================================================================
// CHALLENGE: errors.As — extracting custom types from wrapped chains
//
// After Step 4 tests pass, uncomment the challenge tests in errors_test.go.
// ============================================================================
//
// In Java:
//
//	try { ... }
//	catch (ParseException e) {
//	    System.out.println(e.getInput());  // access custom fields
//	}
//
// In Go, errors.As is like instanceof + cast for error chains:
//
//	var pe *ParseError
//	if errors.As(err, &pe) {
//	    fmt.Println(pe.Input)  // access custom fields
//	}
//
// It walks the wrap chain looking for an error that matches the target type.
//
// YOUR TASK:
// 1. Write a function called ParseBirthMonth that takes a string and returns
//    (int, error).
//    - Call your ParseMonth function.
//    - If ParseMonth returns an error, wrap it:
//      fmt.Errorf("parsing birth month: %w", err)
//    - If successful, return the month and nil
// 2. Uncomment the challenge tests, run: task test -- 04

// Keep the compiler happy — these are used in later steps.
var (
	_ = errors.New
	_ = fmt.Errorf
	_ = strconv.Atoi
	_ = strings.Split
)
