# Contributing to golang-mastery-path

Thank you for your interest in contributing to **golang-mastery-path**! 🎉
This is a progressive learning guide with 150 Go programs from basic to advanced.
Every contribution — new programs, fixes, docs, or ideas — helps Go learners worldwide.

## Code of Conduct

By participating, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).
Please be welcoming, respectful, and inclusive.

## How Can I Contribute?

- Fix a bug or compilation error in an existing program
- Add a new Go program to the curriculum (see open Feature Requests first)
- Improve documentation, comments, or README examples
- Improve tests, CI, linting, or repo tooling

## Getting Started

### 1. Fork and Clone

```bash
# Fork the repo on GitHub, then:
git clone https://github.com/<your-username>/golang-mastery-path.git
cd golang-mastery-path
git remote add upstream https://github.com/imrohit44/golang-mastery-path.git
```

### 2. Set Up Go

- Install Go 1.21 or higher from https://go.dev/dl/
- Verify:

```bash
go version
```

- Initialize / sync dependencies (repo uses stdlib only by default):

```bash
go mod tidy
```

### 3. Run Programs Locally

Each program is self-contained and runnable with `go run`:

```bash
# Run a basics program
go run ./basics/program01.go

# Run an intermediate program
go run ./intermediate/program10.go

# Run an advanced program
go run ./advance/program05.go
```

Or `cd` into a folder and run:

```bash
cd basics
go run program01.go
```

To build all programs and check for compilation errors:

```bash
go build ./...
go vet ./...
```

## Branching and Pull Requests

1. Sync your fork:

```bash
git fetch upstream
git checkout main
git merge upstream/main
```

2. Create a feature branch:

```bash
git checkout -b feat/add-context-timeout-example
# or: fix/program-23-compilation-error
```

3. Make your change. Keep one PR focused on one program / one fix.
4. Commit with a clear message:

```bash
git add .
git commit -m "feat(basics): add program 51 - buffered channels example"
git push origin feat/add-context-timeout-example
```

5. Open a Pull Request against `main` of `imrohit44/golang-mastery-path` and fill out the PR template.

## Requirements Before Opening a Pull Request

All PRs **must** pass formatting and linting:

### 1. `gofmt` (required)

```bash
# Check formatting
gofmt -l .

# Auto-format (should return no output afterwards)
gofmt -w .
```

### 2. `golangci-lint` (required)

Install: https://golangci-lint.run/welcome/install/

```bash
# Run from repo root
golangci-lint run ./...
```

Your PR will not be merged if linting fails. Common issues: unused variables,
ineffectual assignments, `fmt.Sprintf` without formatting directives, missing error checks.

### 3. Manual Checklist

- [ ] Code is `gofmt`-clean
- [ ] `golangci-lint run ./...` passes with no issues
- [ ] `go build ./...` and `go vet ./...` pass
- [ ] New program follows existing naming/style (`programNN.go`, `package main`, commented header explaining concept)
- [ ] README curriculum table updated if you added a new program

## Style Guide for New Programs

- One concept per file, `package main`, runnable via `go run`.
- File header comment: program number, topic, what the learner will learn.
- Prefer stdlib only unless the concept requires an external dependency.
- Use idiomatic Go: `gofmt`, meaningful names, handle errors explicitly, add comments for beginners.
- Keep basics simple, intermediate idiomatic, advanced production-grade.

Example header:

```go
// Program 51: Context with Timeout
// Demonstrates context.WithTimeout for cancelling slow operations.
package main
```

Thank you for helping make Go easier to learn! If unsure, open a Draft PR or Feature Request issue and we'll guide you.
