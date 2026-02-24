local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window_size = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- calulate current window start
local window_start = math.floor(now / window_size) * window_size
local currentkey = key .. ":" .. window_start

-- get current window counter
local current_count = tonumber(redis.call("GET", currentkey) or 0)

if current_count < limit then
    redis.call("INCR", currentkey)
    if current_count == 0 then
        redis.call("EXPIRE", currentkey, window_size)
    end
    return { 1, limit - current_count - 1, window_start + window_size }
else
    return { 0, 0, window_size + window_start }
end
