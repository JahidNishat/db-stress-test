# Crash Course: Redis Lua

## Basics
Redis uses Lua 5.1. It is very simple, like Python but with `local` variables.

**1. Accessing Keys:**
- `KEYS[1]` = The first key you pass from Go.
- `ARGV[1]` = The first argument you pass (like rate, capacity).

**2. Reading from Redis:**
`local value = redis.call("GET", KEYS[1])`
`local fields = redis.call("HMGET", KEYS[1], "tokens", "last_refill")`

**3. Parsing Numbers:**
Redis returns strings. You MUST convert them.
`local tokens = tonumber(fields[1])`

**4. Returning:**
`return 1` (True/Allowed)
`return 0` (False/Denied)

---

## The Logic You Need to Write

```lua
-- Inputs
local key = KEYS[1]
local rate = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local now = tonumber(ARGV[3]) -- Current time in seconds (or micro)
local requested = 1 

-- 1. Read current state
local info = redis.call("HMGET", key, "tokens", "last_refill")
local tokens = tonumber(info[1])
local last_refill = tonumber(info[2])

-- 2. Initialize if missing (First time user)
if not tokens then
    tokens = capacity
    last_refill = now
end

-- 3. Calculate Refill
local delta = math.max(0, now - last_refill)
local filled_tokens = math.min(capacity, tokens + (delta * rate))

-- 4. Check if we can afford the request
local allowed = 0
local new_tokens = filled_tokens

if filled_tokens >= requested then
    allowed = 1
    new_tokens = filled_tokens - requested
end

-- 5. Save State
redis.call("HMSET", key, "tokens", new_tokens, "last_refill", now)
redis.call("EXPIRE", key, 60) -- Cleanup unused keys after 60s

return allowed
```

**Your Task:**
Put this string into a Go variable:
`var script = "..."`
And run it using `rdb.Eval(ctx, script, []string{"my_key"}, rate, capacity, now)`.
