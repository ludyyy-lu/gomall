-- KEYS[1]: 秒杀库存 key
-- ARGV[1]: 扣减的数量（一般是1）

local stock = tonumber(redis.call("GET", KEYS[1]))

if stock == nil then
    return -1  -- key 不存在
end

if stock <= 0 then
    return 0  -- 库存不足
end

redis.call("DECRBY", KEYS[1], ARGV[1])
return 1  -- 扣库存成功
