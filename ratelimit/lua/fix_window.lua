local val = redis.call('GET', KEYS[1])
local limit = tonumber(ARGV[2])

if val == false then
    if limit < 1 then
        return "true"
    else
        redis.call('SET', KEYS[1], 1, 'PX', ARGV[1])
        return "false"
    end
else
    local current = tonumber(val)
    if current < limit then
        redis.call('INCR', KEYS[1])
        return "false"
    else
        return "true"
    end
end