# Lesson 2: Distributed Rate Limiting (Redis + Lua)

## The Problem: Race Conditions Again
In Go (Lesson 1), we used `sync.Mutex` to lock memory.
In a Distributed System (10 servers), we can't share a Mutex.
We use **Redis**.

But if 2 servers read Redis at the same time:
1.  Server A reads `tokens = 5`.
2.  Server B reads `tokens = 5`.
3.  Server A writes `tokens = 4`.
4.  Server B writes `tokens = 4`.
**We lost a subtraction.** This allows traffic to leak.

## The Solution: Redis Lua Scripts
Redis can execute a small script written in **Lua**.
**Crucial Feature:** Redis guarantees the script is **Atomic**.
While the script runs, *no other commands run*. It's like a giant Mutex around the logic.

## The Assignment
We will implement the **Token Bucket** inside Redis using Lua.

### The Data Structure (Redis Hash)
Key: `limiter:{user_id}`
Fields:
- `tokens`: Current count (float)
- `last_refill`: Timestamp (Unix Microseconds)

### The Logic (Lua)
1.  Get current `tokens` and `last_refill`.
2.  Calculate time elapsed (`now - last_refill`).
3.  Calculate refill: `delta * rate`.
4.  New tokens = `min(max_tokens, old_tokens + refill)`.
5.  If `new_tokens >= 1`:
    - `new_tokens -= 1`
    - Update Redis (`tokens`, `now`).
    - Return **1** (Allowed).
6.  Else:
    - Update Redis (`tokens` - just to save the refill, `now`).
    - Return **0** (Denied).

### Setup
1.  Create folder `Academy/Week-02-Distributed-Patterns/redis-limiter`.
2.  `cd` into it.
3.  `go mod init redis-limiter`.
4.  `go get github.com/redis/go-redis/v9`.
5.  Create `main.go`.

**Goal:**
Write the Go code that connects to Redis and runs this Lua script.
I will provide the Lua script skeleton in the next prompt if you get stuck, but try to think about the inputs.
(Inputs: `KEYS[1]` is the user key. `ARGV[1]` is rate. `ARGV[2]` is capacity. `ARGV[3]` is current_time).
