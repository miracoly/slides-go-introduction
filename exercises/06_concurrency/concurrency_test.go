package concurrency

import (
	// Uncomment "testing" when you uncomment your first test.
	// Add "time" when you reach Step 4, and "sort" when you reach the challenge.
	// "sort"
	// "testing"
	// "time"
)

// ============================================================================
// Step 1: Goroutines and basic channel communication
//
// Implement Produce in concurrency.go, then uncomment the tests below.
// ============================================================================

// // The go keyword launches a function in a new goroutine.
// // The channel is the pipe that connects it back to the caller.
// // This is fundamentally different from Java: no Future, no get(),
// // no ExecutorService — just "go" and a channel.
// func TestProduceSingleValue(t *testing.T) {
// 	ch := make(chan int)
// 	go Produce(ch, 42)
//
// 	got := <-ch
// 	if got != 42 {
// 		t.Errorf("expected 42, got %d", got)
// 	}
// }

// func TestProduceMultipleValues(t *testing.T) {
// 	ch := make(chan int)
// 	go Produce(ch, 1, 2, 3)
//
// 	for _, want := range []int{1, 2, 3} {
// 		got := <-ch
// 		if got != want {
// 			t.Errorf("expected %d, got %d", want, got)
// 		}
// 	}
// }

// ============================================================================
// Step 2: Buffered channels
//
// After Step 1 passes, implement FillBuffer in concurrency.go,
// then uncomment the tests below.
// ============================================================================

// // A buffered channel is like Java's ArrayBlockingQueue:
// // sends don't block until the buffer is full.
// // An unbuffered channel (Step 1) blocks on EVERY send.
// func TestFillBufferSendsAll(t *testing.T) {
// 	ch := make(chan string, 3)
// 	FillBuffer(ch, "a", "b", "c")
//
// 	// All three sends completed without blocking (no goroutine needed!)
// 	// because the buffer has capacity 3.
// 	if len(ch) != 3 {
// 		t.Errorf("expected channel length 3, got %d", len(ch))
// 	}
// }

// func TestFillBufferFIFO(t *testing.T) {
// 	ch := make(chan string, 3)
// 	FillBuffer(ch, "a", "b", "c")
//
// 	// Channels are FIFO — first in, first out. Just like a queue.
// 	for _, want := range []string{"a", "b", "c"} {
// 		got := <-ch
// 		if got != want {
// 			t.Errorf("expected %q, got %q", want, got)
// 		}
// 	}
// }

// ============================================================================
// Step 3: Channel close and range — the generator pattern
//
// After Step 2 passes, implement GenerateSquares in concurrency.go,
// then uncomment the tests below.
// ============================================================================

// // "range" over a channel reads values until the channel is CLOSED.
// // If you forget to close, range blocks forever — your program hangs.
// // This is the most common goroutine leak in Go.
// func TestGenerateSquares(t *testing.T) {
// 	ch := GenerateSquares(5)
//
// 	want := []int{1, 4, 9, 16, 25}
// 	var got []int
// 	for v := range ch {
// 		got = append(got, v)
// 	}
//
// 	if len(got) != len(want) {
// 		t.Fatalf("expected %d values, got %d", len(want), len(got))
// 	}
// 	for i := range want {
// 		if got[i] != want[i] {
// 			t.Errorf("square[%d]: expected %d, got %d", i, want[i], got[i])
// 		}
// 	}
// }

// func TestGenerateSquaresZero(t *testing.T) {
// 	ch := GenerateSquares(0)
//
// 	var got []int
// 	for v := range ch {
// 		got = append(got, v)
// 	}
//
// 	if len(got) != 0 {
// 		t.Errorf("expected 0 values, got %d", len(got))
// 	}
// }

// ============================================================================
// Step 4: Select with timeout
//
// After Step 3 passes, implement FirstResult and WithTimeout in
// concurrency.go, then uncomment the tests below.
// NOTE: Add "time" to the import block at the top of this file.
// ============================================================================

// // select is Go's concurrency multiplexer. It waits on multiple channels
// // and executes whichever case is ready first. This has NO equivalent in
// // Java — the closest is CompletionService.take().
// func TestFirstResultCh1Wins(t *testing.T) {
// 	ch1 := make(chan string, 1)
// 	ch2 := make(chan string, 1)
// 	ch1 <- "from ch1"
//
// 	got := FirstResult(ch1, ch2)
// 	if got != "from ch1" {
// 		t.Errorf("expected \"from ch1\", got %q", got)
// 	}
// }

// func TestFirstResultCh2Wins(t *testing.T) {
// 	ch1 := make(chan string, 1)
// 	ch2 := make(chan string, 1)
// 	ch2 <- "from ch2"
//
// 	got := FirstResult(ch1, ch2)
// 	if got != "from ch2" {
// 		t.Errorf("expected \"from ch2\", got %q", got)
// 	}
// }

// func TestWithTimeoutSuccess(t *testing.T) {
// 	ch := make(chan string, 1)
// 	ch <- "done"
//
// 	val, ok := WithTimeout(ch, 100*time.Millisecond)
// 	if !ok {
// 		t.Fatal("expected ok=true, got false (timeout)")
// 	}
// 	if val != "done" {
// 		t.Errorf("expected \"done\", got %q", val)
// 	}
// }

// func TestWithTimeoutExpires(t *testing.T) {
// 	ch := make(chan string) // unbuffered, nobody will send
//
// 	val, ok := WithTimeout(ch, 50*time.Millisecond)
// 	if ok {
// 		t.Fatalf("expected ok=false (timeout), got value %q", val)
// 	}
// 	if val != "" {
// 		t.Errorf("expected empty string on timeout, got %q", val)
// 	}
// }

// ============================================================================
// CHALLENGE: Fan-in — merging multiple channels into one
//
// After Step 4 passes, implement FanIn in concurrency.go,
// then uncomment the tests below.
// NOTE: Add "sort" to the import block at the top of this file.
// ============================================================================

// // Fan-in: many producers, one consumer. Each input gets its own
// // forwarding goroutine. sync.WaitGroup tracks when all are done.
// // This combines everything: goroutines, channels, close, WaitGroup.
// func TestFanInTwoChannels(t *testing.T) {
// 	ch1 := make(chan string, 1)
// 	ch2 := make(chan string, 1)
// 	ch1 <- "a"
// 	ch2 <- "b"
// 	close(ch1)
// 	close(ch2)
//
// 	out := FanIn(ch1, ch2)
//
// 	var got []string
// 	for v := range out {
// 		got = append(got, v)
// 	}
//
// 	sort.Strings(got) // order is non-deterministic
// 	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
// 		t.Errorf("expected [a b], got %v", got)
// 	}
// }

// func TestFanInThreeChannels(t *testing.T) {
// 	make1 := func(vals ...string) <-chan string {
// 		ch := make(chan string, len(vals))
// 		for _, v := range vals {
// 			ch <- v
// 		}
// 		close(ch)
// 		return ch
// 	}
//
// 	out := FanIn(make1("x", "y"), make1("a"), make1("m", "n", "o"))
//
// 	var got []string
// 	for v := range out {
// 		got = append(got, v)
// 	}
//
// 	if len(got) != 6 {
// 		t.Fatalf("expected 6 values, got %d: %v", len(got), got)
// 	}
//
// 	sort.Strings(got)
// 	want := []string{"a", "m", "n", "o", "x", "y"}
// 	for i := range want {
// 		if got[i] != want[i] {
// 			t.Errorf("FanIn[%d]: expected %q, got %q", i, want[i], got[i])
// 		}
// 	}
// }
