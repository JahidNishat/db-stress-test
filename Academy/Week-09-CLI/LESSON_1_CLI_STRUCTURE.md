# Week 9: Go CLI Architecture (The Senior Standard)

## 1. Project Layout
A Senior Go project is not just a `main.go`. It follows the **Standard Go Project Layout**:

- **`cmd/app/main.go`**: The entry point. It should contain **Zero Logic**. It only parses flags and calls `internal`.
- **`internal/`**: Private library code. Cannot be imported by other projects.
- **`pkg/`**: Public library code. Can be imported.

### Why separate `main` and `runner`?
So you can reuse your logic!
- **CLI:** `go run cmd/dbstress/main.go` -> calls `runner.Run`.
- **HTTP API:** `POST /test` -> calls `runner.Run`.
- **Unit Test:** `TestStress` -> calls `runner.Run`.

## 2. Cobra (The CLI Library)
We use `spf13/cobra` for robust flag parsing.
- **Commands:** `rootCmd`, `runCmd`.
- **Flags:** `rootCmd.Flags().IntVarP(...)`.

## 3. Interfaces (Polymorphism)
We defined `Workload` interface so our runner doesn't care if it's running a "Bank Transfer" or "E-commerce" test.
```go
type Workload interface {
    Setup(db *sql.DB) error
    Run(db *sql.DB) error
}
```
This is **Dependency Inversion**. The high-level `Runner` depends on an abstraction, not a concrete implementation.
