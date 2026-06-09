## ADDED Requirements

### Requirement: Go module at project root
The project SHALL have a `go.mod` file at the repository root defining the module path. The module name SHALL be `go-introduction`.

#### Scenario: Module initialization
- **WHEN** a participant clones the repository
- **THEN** `go.mod` exists at the root with `module go-introduction` and a Go version directive

### Requirement: Taskfile with test runner
The project SHALL have a `Taskfile.yml` at the root providing go-task targets for running exercise tests.

#### Scenario: Run all exercise tests
- **WHEN** participant runs `task test`
- **THEN** `go test` executes for all packages under `exercises/`

#### Scenario: Run a specific exercise by number
- **WHEN** participant runs `task test -- 01`
- **THEN** `go test` executes only for the package whose directory name starts with `01`

### Requirement: Exercise directory convention
Each exercise SHALL reside in its own directory under `exercises/` following the naming pattern `NN_snake_case_name/` where NN is a zero-padded number. Each exercise directory SHALL contain exactly two files: a `.go` file with stubbed implementations and a `_test.go` file with tests.

#### Scenario: Exercise directory structure
- **WHEN** a new exercise is added
- **THEN** it follows the pattern `exercises/NN_name/name.go` and `exercises/NN_name/name_test.go`

### Requirement: Stubbed functions use panic
All unimplemented exercise functions SHALL use `panic("TODO: implement")` as their body so that the code compiles but tests clearly indicate unfinished work.

#### Scenario: Running tests before implementation
- **WHEN** participant runs `task test` before implementing anything
- **THEN** all tests fail with panic messages containing "TODO: implement"
