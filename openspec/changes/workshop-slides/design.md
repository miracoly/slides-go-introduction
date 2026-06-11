## Context

The workshop repo has 6 completed exercises and a reveal-md template at `slides/`. The existing `slides/slides.md` contains placeholder ProtoBuf content that needs to be replaced with Go workshop slides. The audience is Java developers at the company.

## Goals / Non-Goals

**Goals:**

- Persuasive "Why Go" opening that motivates Java developers to engage
- Concise concept slides with Go code examples — just enough to enable the exercises
- Clear two-part structure with exercise breaks

**Non-Goals:**

- Comprehensive Go language reference (exercises handle depth)
- Java comparison code on slides (done verbally by presenter)
- Exercise instructions on slides (exercises are self-contained)
- Project setup slides (handled before the presentation)

## Decisions

### Single file, two-part structure

All slides go into `slides/slides.md` using reveal-md horizontal slide separators (`---`). Part 1 and Part 2 are separated by an "Exercise Time" slide. No vertical slides needed — keeps navigation simple.

**Alternative considered:** Separate files per part — rejected because reveal-md works best with a single markdown file and the total slide count is small.

### Code-heavy, bullet-light

Each concept slide leads with a Go code snippet, followed by 2–3 bullet points of explanation. Code examples are the "images" of a dev presentation. Keep bullets minimal — the presenter adds context verbally.

### Slide ordering follows exercise dependency

Slides introduce concepts in the exact order exercises need them. No concept is introduced before the exercise that uses it, no exercise references a concept not yet shown.

### Persuasive tone for "Why Go"

The opening section uses a persuasive tone (not academic) to sell Go's developer experience advantages: readability, explicitness, fast compilation, single binary, opinionated tooling. The Fabrice Bellard quote about reading code being 10x harder than writing it anchors the argument.

## Risks / Trade-offs

- **Too little on slides** → Presenter must fill gaps verbally. Mitigation: this is intentional — 10–15 min talks should be tight, exercises do the heavy lifting.
- **Code examples too simple** → May not convey the concept. Mitigation: examples are drawn from patterns used in the actual exercises.
