
-- 限流对象
local key = KEYS[1]
-- 窗口大小
local window= tonumber(ARGV[1])
-- 限流阈值
local threshold = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local min = now - window


redis.call('ZREMRANGEBYSCORE', key, '-inf', min)

local count = redis.call('ZCOUNT', key, '-inf', '+inf')

if count < threshold then
    redis.call('ZADD', key, now, now)
    redis.call('EXPIRE', key, window)
    return "false"
else
    return "true"
end