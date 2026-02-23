local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = 1 -- we consume one token per request

-- Fetch current status from the redis hash
local bucket = redis.call("HMGET", key, "tokens", "last_refill")
local tokens = tonumber(bucket[1])
local last_refill = tonumber(bucket[2])

-- if bucket not exists initalize it full
if not tokens then
    tokens = capacity
    last_refill = now
else
    -- calculate how many token we have to add on the basic of time passed
    local time_passed = math.max(0, now - last_refill)
    local refill_amount = math.floor(time_passed * refill_rate)

    -- add the token but remember don't exceed the capacity of bucket
    tokens = math.min(capacity, tokens + refill_amount)
    if refill_amount > 0 then
        last_refill = now
    end
end

-- check if request is allowed
if tokens >= requested then
    tokens = tokens - requested
    -- save the state and add TTL of 60 second so we can clean up inactive users
    redis.call("HMSET", key, "tokens", tokens, "last_refill", last_refill)
    redis.call("EXPIRE", 60)
    return 1
else
    -- not enough token save state without consuming
    redis.call("HMSET", key, "tokens", tokens, "last_refill", last_refill)
    redis.call("EXPIRE", 60)
    return 0
end
