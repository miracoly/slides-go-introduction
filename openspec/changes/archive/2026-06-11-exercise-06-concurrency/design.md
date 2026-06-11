## Context

Sixth and final exercise in the Go workshop for experienced Java developers. This is the last exercise of Part 2. Participants have completed types, structs, interfaces, error handling, and collections. They now tackle Go's signature feature: lightweight concurrency with goroutines and channels.

Java developers know `Thread`, `ExecutorService`, `Future`, `synchronized`, and `ConcurrentHashMap`. Go replaces all of this with two primitives: goroutines (lightweight threads) and channels (typed pipes for communication). The mental model shift is from "shared memory with locks" to "communicate by sharing messages."

## Goals / Non-Goals

**Goals:**
- Show that goroutines are trivially cheap to spawn (`go func()`)
- Teach channels as the primary communication mechanism between goroutines
- Demonstrate unbuffered channels as synchronization points
- Show buffered channels and when they matter
- Teach channel direction for API safety (`chan<-` send-only, `<-chan` receive-only)
- Demonstrate `range` over channels (close + iterate)
- Teach `select` for multiplexing with timeout
- Build a fan-in pattern that combines real concepts

**Non-Goals:**
- `sync.Mutex` / `sync.RWMutex` (important but secondary to channels for an intro)
- `sync.WaitGroup` (useful but introduces shared-state thinking)
- `context.Context` for cancellation (too advanced for a 2-hour workshop)
- Race condition detection / `-race` flag (good to mention in comments, don't exercise)
- Worker pools, pipelines, or complex patterns beyond fan-in

## Decisions

### Step structure: build up from goroutines to select

Each step adds one concept. The progression mirrors how you'd build a real concurrent program:

1. **Goroutines + channels basics** — launch goroutines, send/receive on channels
2. **Buffered channels** — understand blocking behavior differences
3. **Channel direction and range** — close channels, iterate results
4. **Select with timeout** — multiplex channels, handle slow producers
5. **Challenge: Fan-in** — combine multiple channels into one

### Domain: simulated concurrent tasks

Functions simulate "work" by sending results to channels. No real I/O, no external dependencies. Using `time.Sleep` for simulated delays in select/timeout tests.

**Why not real I/O:** Real network calls introduce flakiness, DNS dependencies, and slow test execution. Simulated work with channels tests the same concurrency patterns without the noise.

### Step 1: Goroutines and basic channel communication

Participants write:
- `Produce(ch chan<- int, values ...int)` — sends values to a channel (demonstrates send-only direction)
- A test that launches `Produce` as a goroutine and reads from the channel

This establishes the foundational pattern: launch a goroutine, communicate via channel. The test demonstrates that a channel receive blocks until a send occurs — a key difference from Java's `Future.get()`.

### Step 2: Buffered channels

Participants write:
- `FillBuffer(ch chan<- string, items ...string)` — sends items to a buffered channel

Tests demonstrate that a buffered channel doesn't block until full, contrasting with unbuffered channels that block immediately. This maps to Java's `BlockingQueue` with a fixed capacity.

### Step 3: Channel close and range

Participants write:
- `GenerateSquares(n int) <-chan int` — creates a channel, launches a goroutine that sends squares of 1..n, closes the channel, returns it

This teaches the "generator" pattern: create a channel, do work in a goroutine, close when done. The caller uses `range` to iterate. Closing is important — without it, `range` blocks forever. Return type `<-chan int` shows receive-only direction.

### Step 4: Select with timeout

Participants write:
- `FirstResult(ch1, ch2 <-chan string) string` — returns whichever channel produces a value first, using `select`
- `WithTimeout(ch <-chan string, timeout time.Duration) (string, bool)` — returns the value if received before timeout, or `("", false)` if timeout expires

`select` is Go's multiplexer — like `epoll` but for channels. It's the construct that makes Go's channel-based concurrency practical. The timeout pattern with `time.After` is extremely common in production Go.

### Challenge: Fan-in

Participants write:
- `FanIn(channels ...<-chan string) <-chan string` — merges multiple input channels into a single output channel

This is a classic Go concurrency pattern. It combines goroutines, channels, close, and select into one function. Each input channel gets its own goroutine that forwards messages to the output. A coordinating goroutine waits for all inputs to complete, then closes the output.

### Test design: deterministic where possible, timeouts where not

- Steps 1-3: fully deterministic (send known values, receive them, compare)
- Step 4: uses short timeouts (50ms-100ms) and `time.After` to avoid flaky tests
- Challenge: uses small channel counts and known inputs

## Risks / Trade-offs

- **Concurrency tests can be flaky** → Mitigated by keeping timeouts generous relative to work, using deterministic patterns where possible, and avoiding real I/O.
- **Too ambitious for remaining time** → Steps 1-3 are the core. Step 4 and challenge are stretch goals. Even completing steps 1-2 gives participants a working understanding of goroutines and channels.
- **No WaitGroup** → Some participants may ask how to wait for multiple goroutines. Comments mention WaitGroup exists but frame channels as the idiomatic first choice. This is a deliberate pedagogical simplification.
- **Goroutine leaks in incomplete implementations** → Tests use timeouts to prevent hanging. Comments warn about closing channels and draining goroutines.
