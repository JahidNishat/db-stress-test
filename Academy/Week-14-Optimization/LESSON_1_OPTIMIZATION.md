# Week 14: Optimization Techniques

## 1. Pre-allocation
**Problem:** `append(slice, item)` often has to allocate a new, larger array and copy data.
**Fix:** `make([]T, 0, capacity)`.
If you know you need 10,000 items, allocate the memory upfront.

## 2. Removing Hot Path Allocations
In a loop running 100k times/sec:
- **Bad:** `fmt.Sprintf("%d", i)` (Allocates string every time).
- **Good:** `strconv.Itoa(i)` (Optimized, fewer allocations).
- **Best:** Reuse a buffer.

## 3. Throttling the UI
Updating the screen 1000 times a second is wasteful (human eye sees 60fps).
- **Fix:** Only recalculate heavy math (sorting P99) every 100ms or 1s, not every microsecond.
