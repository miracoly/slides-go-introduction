## Why

Exercise 06 is the final exercise of the workshop and the capstone of Part 2. Concurrency is Go's defining feature — goroutines and channels are the reason many teams adopt Go. Java developers know threads and `ExecutorService`, but Go's concurrency model is fundamentally different: goroutines are lightweight (not OS threads), channels replace shared memory for communication, and `select` provides built-in multiplexing. This exercise gives participants hands-on experience with the concurrency primitives they'll encounter in every Go service.

## What Changes

- Create exercise 06 (`exercises/06_concurrency/`) with:
  - `concurrency.go` — step-by-step instruction comments for participants to implement goroutines, channels, and select from scratch
  - `concurrency_test.go` — tests covering goroutine launching, unbuffered and buffered channels, channel direction, range over channels, select with timeout, and a fan-in pattern
- Follows the progressive unlock pattern: all tests commented out, participants uncomment step by step
- Participants define all functions themselves
- Domain: concurrent task execution and result collection — simulating parallel work with sleep-based "work" functions
- Challenge section for a fan-in pattern combining multiple channels

## Capabilities

### New Capabilities
- `exercise-06`: Exercise covering goroutine launching with the `go` keyword, unbuffered channels for synchronization, buffered channels, sending/receiving, channel direction (`chan<-`, `<-chan`), closing channels and range over channels, `select` for multiplexing with timeout, and the fan-in pattern for combining channel results

### Modified Capabilities

## Impact

- New files: `exercises/06_concurrency/concurrency.go`, `exercises/06_concurrency/concurrency_test.go`
- No existing code is affected
- Follows the directory convention from `project-setup` spec
