local keys = KEYS[1]
local capacity = tonumber(ARGV[1])
local leak_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- fetch current water level
local state = redis.call("HMGET", keys, "water", "last_leak_time")
local water = tonumber(state[1]) or 0
local last_leak_time = tonumber(state[2]) or now

-- calculate the how much request has processed
local time_passed = math.max(0, now - last_leak_time)
local leaked = math.floor(time_passed * leak_rate)

-- update the queue
local water = math.max(0, water - leaked)

-- check if leftover request is able to served or not
if water < capacity then
    water = water + 1
    redis.call("HMSET", keys, "water", water, "last_leak_time", now)

    local time_to_empty = math.ceil(water / leak_rate)
    redis.call("EXPIRE", keys, time_to_empty + 1)

    return { 1, capacity - water, now + 1 }
else
    redis.call("HMSET", keys, "water", water, "last_leak_time", last_leak_time)
    local time_to_empty = math.ceil(water / leak_rate)
    redis.call("EXPIRE", keys, time_to_empty + 1)

    return { 0, 0, now + 1 }
end
