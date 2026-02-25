# Week 2: Distributed Patterns & Rate Limiting

**Goal:** Move from protecting *memory* (Mutex) to protecting *services* (Rate Limits).

## The Concept: Protecting the System
In Week 1, we protected a variable (`counter`) from 1000 goroutines using a Mutex.
In Week 2, we protect our **API** from 1,000,000 users using a **Rate Limiter**.

If you don't rate limit:
1.  One user spams your API.
2.  Your database CPU hits 100%.
3.  **Everyone else suffers.** (The "Noisy Neighbor" problem).

---

## Lesson 1: The Token Bucket Algorithm

Imagine a bucket that holds **Tokens**.
1.  **Capacity:** The bucket holds max 10 tokens.
2.  **Refill Rate:** We add 1 token every second.
3.  **Consumption:** To make a request, a user must take 1 token.
4.  **Empty Bucket:** If the bucket is empty, the request is **REJECTED** (HTTP 429 Too Many Requests).

**Why this is standard:**
It allows "Bursts" (you can make 10 requests instantly), but enforces a long-term average (1 per second).

---

## Assignment: The In-Memory Limiter

We will build a pure Go Rate Limiter from scratch (no libraries).

**Struct:**
```go
type RateLimiter struct {
    mu         sync.Mutex
    tokens     float64
    maxTokens  float64
    refillRate float64 // tokens per second
    lastRefill time.Time
}
```

**Logic (The `Allow()` function):**
1.  Lock the Mutex.
2.  Calculate time passed since `lastRefill`.
3.  Add new tokens based on time * rate.
4.  Cap tokens at `maxTokens`.
5.  If `tokens >= 1`:
    - `tokens -= 1`
    - Update `lastRefill = now`
    - Return `true` (Allowed)
6.  Else:
    - Return `false` (Denied)
7.  Unlock.

**Task:**
1.  Create `Academy/Week-02-Distributed-Patterns/limiter.go`.
2.  Implement `NewRateLimiter(rate, max)` and `Allow()`.
3.  Write a `main` function that simulates a burst of 20 requests.
    - Rate: 1 token/sec. Max: 5.
    - Loop 20 times. Print "Allowed" or "Denied".
    - `time.Sleep(200 * time.Millisecond)` between requests.

**Go.**
