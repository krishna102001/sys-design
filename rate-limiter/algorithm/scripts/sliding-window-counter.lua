local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window_size = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- calculate current and previous window start
local curr_window_start = math.floor(now / window_size) * window_size
local prev_window_start = curr_window_start - window_size

-- create redis keys
local curr_key = key .. ":" .. curr_window_start
local prev_key = key .. ":" .. prev_window_start

-- get the count of the request of current and previous
local curr_count = tonumber(redis.call("GET", curr_key) or 0)
local prev_count = tonumber(redis.call("GET", prev_key) or 0)

-- calculate overlap percentage
local time_passed_in_current = now - curr_window_start
local overlap_percentage = 1.0 - (time_passed_in_current / window_size)

-- calculate approx request till now served
local estimated_count = curr_count + math.floor(prev_count * overlap_percentage)

if estimated_count < limit then
    redis.call("INCR", curr_key)
    redis.call('EXPIRE', curr_key, window_size * 2)
    return { 1, limit - estimated_count - 1, curr_window_start + window_size }
else
    return { 0, 0, curr_window_start + window_size }
end
