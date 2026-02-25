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
local refill = (delta / 1000000) * rate
local filled_tokens = math.min(capacity, tokens + refill)

-- 4. Check if we can afford the request
local allowed = 0
local new_tokens = filled_tokens

if filled_tokens >= requested then
    allowed = 1
    new_tokens = filled_tokens - requested
end

-- 5. Save state
redis.call("HMSET", key, "tokens", new_tokens, "last_refill", now)
redis.call("EXPIRE", key, 60) -- Cleanup unused keys after 60s

return allowed
