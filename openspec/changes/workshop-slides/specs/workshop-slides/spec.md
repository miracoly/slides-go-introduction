## ADDED Requirements

### Requirement: Slide deck has two-part structure

The slide deck SHALL be divided into Part 1 and Part 2, each ending with an "Exercise Time" slide that links to the workshop repository.

#### Scenario: Part 1 covers exercises 1–3

- **WHEN** the presenter delivers Part 1
- **THEN** the slides cover "Why Go", types & functions, structs & methods, and interfaces before the exercise break

#### Scenario: Part 2 covers exercises 4–6

- **WHEN** the presenter delivers Part 2
- **THEN** the slides cover errors, collections, and concurrency before the exercise break

### Requirement: Why Go motivation section

Part 1 SHALL open with a persuasive "Why Go" section that motivates Java developers to engage with Go. It MUST cover: readability as the primary advantage (anchored by the quote about reading code being 10x harder than writing), explicitness and no magic, fast compilation, single binary deployment, and opinionated tooling (`go fmt`).

#### Scenario: Readability argument

- **WHEN** the "Why Go" section is presented
- **THEN** it frames Go's design philosophy around optimizing for the code reader, not the writer

#### Scenario: No magic argument

- **WHEN** the "Why Go" section is presented
- **THEN** it contrasts Go's explicitness with hidden control flow (no annotations, no DI magic, no implicit behavior)

### Requirement: Concept slides include Go code examples

Each concept slide SHALL lead with a Go code snippet that demonstrates the concept. Bullet points SHALL be minimal (2–3 max) and support the code example.

#### Scenario: Types and functions concepts

- **WHEN** the types & functions slides are shown
- **THEN** they include code examples for short variable declaration, zero values, multiple return values, and pointers

#### Scenario: Structs and methods concepts

- **WHEN** the structs & methods slides are shown
- **THEN** they include code examples for struct definition, factory functions, and pointer vs value receivers

#### Scenario: Interfaces concepts

- **WHEN** the interfaces slides are shown
- **THEN** they include code examples for interface definition, implicit satisfaction, and type switches

#### Scenario: Errors concepts

- **WHEN** the errors slides are shown
- **THEN** they include code examples for error returns, sentinel errors, and error wrapping with %w

#### Scenario: Collections concepts

- **WHEN** the collections slides are shown
- **THEN** they include code examples for slices, maps, and the comma-ok pattern

#### Scenario: Concurrency concepts

- **WHEN** the concurrency slides are shown
- **THEN** they include code examples for goroutines, channels, and select

### Requirement: Slides use reveal-md format

The slides SHALL be written in `slides/slides.md` using reveal-md markdown format with `---` as horizontal slide separators. The file SHALL retain the existing YAML frontmatter configuration.

#### Scenario: Valid reveal-md format

- **WHEN** the slide file is processed by reveal-md
- **THEN** it renders as a valid slide deck with correct transitions and styling

### Requirement: Presentation duration is concise

Part 1 presentation SHALL target ~15 minutes. Part 2 presentation SHALL target ~10 minutes. Slides SHALL not be artificially extended — fewer, tighter slides are preferred.

#### Scenario: Slide count is appropriate

- **WHEN** the slides are counted
- **THEN** Part 1 has roughly 8–12 slides and Part 2 has roughly 5–8 slides
