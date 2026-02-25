# Week 12: Reporting & Percentiles

## Why Average Latency is a Lie
If 99 requests take 1ms and 1 request takes 100s:
- **Average:** ~1s. (Looks okay).
- **User Experience:** 1 user is furious.

## Percentiles (P50, P99)
- **P50 (Median):** 50% of users see this speed or faster. The "Normal" experience.
- **P99:** 99% of users see this speed or faster. The "Worst Case" for most people.

## Calculating P99
1.  Store all durations in a slice.
2.  Sort the slice: `sort.Float64s(latencies)`.
3.  Index: `i = int(len * 0.99)`.
4.  Value: `latencies[i]`.

## Memory Trade-off
Storing 1 million floats takes memory.
- **Simple:** Slice (High memory, exact).
- **Advanced:** T-Digest or HDR Histogram (Low memory, approximate).
