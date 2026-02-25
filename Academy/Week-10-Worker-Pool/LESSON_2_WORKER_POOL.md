# Week 10: The Worker Pool Pattern (Concurrency)

## 1. Why Worker Pools?
Launching 1 goroutine per request (e.g., 1000/sec) is too much. 
Launching 1 goroutine (Sequential) is too slow.
A **Worker Pool** (Fixed N goroutines) is the perfect balance. It saturates the CPU/Network without overwhelming memory.

## 2. The Implementation (3 Components)
1.  **Context (`context.WithTimeout`):** Controls the **Duration**. When `ctx.Done()` closes, all workers stop.
2.  **Wait Group (`sync.WaitGroup`):** Controls the **Shutdown**. The main function waits for `wg.Wait()` so workers can finish their last task.
3.  **Atomic Counters (`sync/atomic`):** Controls the **Stats**. `success++` is NOT thread-safe. `atomic.AddInt64(&success, 1)` IS safe.

## 3. The Loop Pattern
```go
for {
    select {
    case <-ctx.Done():
        return // Time is up!
    default:
        work() // Keep working!
    }
}
```
This turns a "Run Once" script into a "Run Until Stopped" engine.
