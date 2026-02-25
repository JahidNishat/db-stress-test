# Week 4, Lesson 2: GC Pressure & Escape Analysis

## 1. Garbage Collection (GC)
Go uses a **Tri-color Mark & Sweep** collector. 
- High `allocs/op` = More work for the GC.
- This causes "latency spikes" (p99 latency increases).

## 2. The Solution: Zero-Allocation
We proved that **reusing objects** (BenchmarkGood) is 2x faster than **allocating new ones** (BenchmarkBad).
- **BenchmarkBad:** ~660ns/op (8192 B/op)
- **BenchmarkGood:** ~300ns/op (0 B/op)

## 3. Escape Analysis
The Go compiler decides where to put variables:
- **Stack:** Local variables, very fast, cleaned up instantly.
- **Heap:** Variables that "escape" (return as pointers), tracked by GC.

### How to check:
```bash
go build -gcflags="-m" file.go
```
- If it says `escapes to heap`, it's going to the GC.
- If it says `does not escape`, it stays on the fast Stack.
