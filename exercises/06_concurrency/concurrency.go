// Package concurrency
//
// In this exercise, YOU define everything from scratch — functions that use
// goroutines, channels, and select. Work through the steps in order.
// Each step tells you what to build. After implementing, uncomment the
// matching tests in concurrency_test.go and run: task test -- 06
package concurrency

import (
	"sync"
	"time"
)

// ============================================================================
// Step 1: Goroutines and basic channel communication
//
// Uncomment Step 1 tests in concurrency_test.go, then implement below.
// ============================================================================
//
// In Java, you spawn threads:
//
//	new Thread(() -> { result = doWork(); }).start();
//	// or
//	ExecutorService pool = Executors.newFixedThreadPool(4);
//	Future<Integer> f = pool.submit(() -> compute());
//	int result = f.get();  // blocks until done
//
// In Go, you launch a goroutine — it's just "go" before a function call:
//
//	go doWork()
//
// Goroutines are MUCH cheaper than Java threads (~2KB stack vs ~1MB).
// You can launch thousands without thinking twice.
//
// To get results back, use a CHANNEL — a typed pipe between goroutines:
//
//	ch := make(chan int)    // create an unbuffered channel of ints
//	go func() {
//	    ch <- 42           // send a value (blocks until someone receives)
//	}()
//	value := <-ch          // receive (blocks until someone sends)
//
// YOUR TASK:
// 1. Write a function called Produce that takes a send-only channel
//    (ch chan<- int) and variadic int values. It sends each value to
//    the channel in order.
//    Note: chan<- means "send only" — you can't read from this channel.
//    This is like a method that only writes to an OutputStream.
// 2. Uncomment Step 1 tests, run: task test -- 06

// ============================================================================
// Step 2: Buffered channels
//
// After Step 1 tests pass, uncomment Step 2 tests in concurrency_test.go.
// ============================================================================
//
// In Java:
//
//	BlockingQueue<String> queue = new ArrayBlockingQueue<>(3);
//	queue.put("a");  // succeeds, doesn't block (capacity left)
//	queue.put("b");  // succeeds
//	queue.put("c");  // succeeds
//	queue.put("d");  // BLOCKS — queue full
//
// In Go, buffered channels work the same way:
//
//	ch := make(chan string, 3)  // buffered channel, capacity 3
//	ch <- "a"                  // doesn't block (buffer has space)
//	ch <- "b"                  // doesn't block
//	ch <- "c"                  // doesn't block
//	ch <- "d"                  // BLOCKS — buffer full
//
// An UNBUFFERED channel (make(chan T)) blocks on EVERY send until
// someone receives. A BUFFERED channel only blocks when the buffer is full.
//
// YOUR TASK:
// 1. Write a function called FillBuffer that takes a send-only channel
//    (ch chan<- string) and variadic string items. It sends each item
//    to the channel in order.
// 2. Uncomment Step 2 tests, run: task test -- 06

// ============================================================================
// Step 3: Channel close and range — the generator pattern
//
// After Step 2 tests pass, uncomment Step 3 tests in concurrency_test.go.
// ============================================================================
//
// In Java, you'd use a Stream to generate a sequence:
//
//	IntStream.rangeClosed(1, n).map(i -> i * i).forEach(System.out::println);
//
// In Go, the "generator" pattern uses a goroutine + channel:
// 1. Create a channel
// 2. Launch a goroutine that sends values and then CLOSES the channel
// 3. Return the channel — the caller uses "range" to iterate
//
// Closing is critical: without it, "range" over a channel blocks forever,
// waiting for more values that will never come.
//
// The return type <-chan int means "receive-only channel". The caller
// can only read from it, not write. This is like returning an InputStream.
//
// YOUR TASK:
// 1. Write a function called GenerateSquares that takes an int n and
//    returns a receive-only channel of ints (<-chan int).
//    - Create a channel
//    - Launch a goroutine that sends the square of each number from 1 to n
//      (i.e., 1*1, 2*2, 3*3, ..., n*n)
//    - Close the channel after all values are sent (use defer!)
//    - Return the channel immediately (the goroutine runs concurrently)
// 2. Uncomment Step 3 tests, run: task test -- 06

// ============================================================================
// Step 4: Select with timeout
//
// After Step 3 tests pass, uncomment Step 4 tests in concurrency_test.go.
// ============================================================================
//
// In Java, you'd use Future with a timeout:
//
//	try {
//	    String result = future.get(100, TimeUnit.MILLISECONDS);
//	} catch (TimeoutException e) { ... }
//
// In Go, "select" lets you wait on MULTIPLE channels at once:
//
//	select {
//	case v := <-ch1:
//	    // ch1 produced a value first
//	case v := <-ch2:
//	    // ch2 produced a value first
//	case <-time.After(100 * time.Millisecond):
//	    // neither produced within 100ms
//	}
//
// select blocks until ONE case is ready, then executes that case.
// If multiple are ready, it picks one at random.
// time.After returns a channel that receives after the duration.
//
// YOUR TASK:
// 1. Write a function called FirstResult that takes two receive-only
//    string channels (ch1, ch2 <-chan string) and returns the first
//    string value that arrives from either channel.
//    Hint: use select with two cases
// 2. Write a function called WithTimeout that takes a receive-only
//    string channel and a time.Duration, and returns (string, bool).
//    - If a value is received before the timeout: return (value, true)
//    - If the timeout expires first: return ("", false)
//    Hint: use select with a channel case and a time.After case
// 3. Uncomment Step 4 tests, run: task test -- 06

// ============================================================================
// CHALLENGE: Fan-in — merging multiple channels into one
//
// After Step 4 tests pass, uncomment the challenge tests.
// ============================================================================
//
// Fan-in is a classic Go concurrency pattern. Multiple producers,
// one consumer. In Java, you'd use a shared ConcurrentLinkedQueue
// or a CompletionService.
//
// In Go, each input channel gets its own forwarding goroutine.
// A coordinator waits for all to finish, then closes the output.
//
// YOUR TASK:
// 1. Write a function called FanIn that takes variadic receive-only
//    string channels (...<-chan string) and returns a single receive-only
//    string channel (<-chan string).
//    - Create an output channel
//    - For each input channel, launch a goroutine that reads all values
//      from that input and sends them to the output channel
//    - Use sync.WaitGroup to wait for all forwarding goroutines to finish,
//      then close the output channel (do this in its own goroutine)
//    - Return the output channel immediately
//
//    Note: this is the one place where sync.WaitGroup shines — you need
//    to know when N goroutines are ALL done. Channels alone can't easily
//    express "wait for N things to finish."
// 2. Uncomment the challenge tests, run: task test -- 06

// Keep the compiler happy — these are used in later steps.
var (
	_ = sync.WaitGroup{}
	_ = time.After
)
