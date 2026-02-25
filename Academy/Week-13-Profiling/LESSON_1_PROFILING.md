# Week 13: Profiling with pprof

## What is Profiling?
Guessing where your code is slow is for amateurs. Profiling tells you exactly which line of code is using the CPU or RAM.

## Enabling pprof
```go
import _ "net/http/pprof"
go func() { http.ListenAndServe("localhost:6060", nil) }()
```

## Using the Tool
1.  **CPU Profile:** `go tool pprof -http=:8080 .../profile?seconds=5`
    - Shows how much time the CPU spends in each function.
    - Look for big boxes (e.g., `fmt.Sprintf`, `syscall`).

2.  **Heap Profile:** `go tool pprof -http=:8080 .../heap`
    - Shows where memory is allocated.
    - Look for `runtime.mallocgc` (creating too many objects).

## Common Bottlenecks
- **String Formatting:** `fmt.Sprintf` is slow and allocates.
- **Sorting:** `sort.Float64s` is O(N log N).
- **GC Pressure:** Creating short-lived objects (pointers) forces the GC to work hard.
