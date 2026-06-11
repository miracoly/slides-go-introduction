## ADDED Requirements

### Requirement: Goroutines and basic channel send
The exercise SHALL include a function `Produce(ch chan<- int, values ...int)` that sends each value to the channel. Tests SHALL launch this function as a goroutine and receive values from the channel.

#### Scenario: Send and receive single value
- **WHEN** `Produce` is called as a goroutine with a channel and the value `42`
- **THEN** receiving from the channel yields `42`

#### Scenario: Send and receive multiple values
- **WHEN** `Produce` is called as a goroutine with values `1, 2, 3`
- **THEN** receiving three times from the channel yields `1`, `2`, `3` in order

### Requirement: Buffered channels
The exercise SHALL include a function `FillBuffer(ch chan<- string, items ...string)` that sends items to a buffered channel. Tests SHALL demonstrate that buffered channels accept sends without blocking until the buffer is full.

#### Scenario: Send to buffered channel
- **WHEN** `FillBuffer` is called with a buffered channel (capacity 3) and items `"a", "b", "c"`
- **THEN** all three sends complete without blocking and the channel length is `3`

#### Scenario: Receive from buffered channel
- **WHEN** three items are sent to a buffered channel and then received
- **THEN** the items arrive in FIFO order: `"a"`, `"b"`, `"c"`

### Requirement: Channel close and range with generator pattern
The exercise SHALL include a function `GenerateSquares(n int) <-chan int` that creates a channel, launches a goroutine to send the squares of 1 through n, closes the channel, and returns the receive-only channel. Tests SHALL use `range` to iterate the results.

#### Scenario: Generate squares of 1 to 5
- **WHEN** `GenerateSquares(5)` is called
- **THEN** ranging over the returned channel yields `1, 4, 9, 16, 25` and then the loop ends (channel closed)

#### Scenario: Generate squares of 0
- **WHEN** `GenerateSquares(0)` is called
- **THEN** ranging over the returned channel yields no values (channel is immediately closed)

#### Scenario: Return type is receive-only
- **WHEN** `GenerateSquares` is called
- **THEN** the return type is `<-chan int` (receive-only channel direction)

### Requirement: Select with first result
The exercise SHALL include a function `FirstResult(ch1, ch2 <-chan string) string` that uses `select` to return whichever channel produces a value first.

#### Scenario: First channel is faster
- **WHEN** `ch1` has a value ready and `ch2` does not
- **THEN** `FirstResult` returns the value from `ch1`

#### Scenario: Second channel is faster
- **WHEN** `ch2` has a value ready and `ch1` does not
- **THEN** `FirstResult` returns the value from `ch2`

### Requirement: Select with timeout
The exercise SHALL include a function `WithTimeout(ch <-chan string, timeout time.Duration) (string, bool)` that returns the value and `true` if received before the timeout, or empty string and `false` if the timeout expires.

#### Scenario: Value received before timeout
- **WHEN** the channel produces `"done"` before the timeout expires
- **THEN** `WithTimeout` returns `"done", true`

#### Scenario: Timeout expires
- **WHEN** no value arrives before the timeout expires
- **THEN** `WithTimeout` returns `"", false`

### Requirement: Fan-in challenge
The exercise SHALL include a commented-out challenge section where participants implement `FanIn(channels ...<-chan string) <-chan string` that merges multiple input channels into a single output channel. Each input channel gets a goroutine that forwards its messages. The output channel is closed when all inputs are exhausted.

#### Scenario: Merge two channels
- **WHEN** `FanIn` is called with two channels that produce `"a"` and `"b"` respectively
- **THEN** both values appear on the output channel (order may vary) and the output channel is closed after both inputs close

#### Scenario: Merge three channels with multiple values
- **WHEN** `FanIn` is called with three channels producing `["x","y"]`, `["a"]`, `["m","n","o"]` respectively
- **THEN** all 6 values appear on the output channel and the channel closes

### Requirement: Progressive unlock test structure
The exercise SHALL use the same progressive unlock pattern as previous exercises: all tests commented out with instructions to uncomment step by step. Imports SHALL be commented out when not yet needed.

#### Scenario: Clean compilation
- **WHEN** the exercise files are compiled without any modifications
- **THEN** they compile without errors

### Requirement: Participants define all functions
The exercise `.go` file SHALL contain only instruction comments and suppressed imports. Participants MUST define all functions themselves. No functions SHALL be pre-defined.

#### Scenario: Empty starting point
- **WHEN** a participant opens `concurrency.go`
- **THEN** they see instruction comments for each step but no pre-defined functions (only suppressed imports)
