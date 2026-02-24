local key          = KEYS[1]
local limit        = tonumber(ARGV[1])
local window_size  = tonumber(ARGV[2])
local now          = tonumber(ARGV[3])
local request_id   = tonumber(ARGV[4])

-- current window start
local window_start = now - window_size

-- remove outdated timestamp
redis.call("ZREMRANGEBYSCORE", key, "-inf", window_start)

-- count current request
local current_request = redis.call('ZCARD', key)

if current_request < limit then
    redis.call("ZADD", key, now, request_id)
    redis.call("EXPIRE", key, window_size + 1)
    return { 1, limit - current_request - 1, window_size + now }
else
    return { 0, 0, window_size + now }
end
