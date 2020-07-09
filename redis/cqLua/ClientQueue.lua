-- 客户排队，可以重复调用
-- 需要注意，索引值是从0开始
-- 如果member的索引值是1，表示前面还有一个member


-- 系统消息
local SystemMessage = function(stream,content)
    redis.call("xadd",stream,"*",
        "SenderType","system",
        "SenderId","",
        "SenderName","系统消息",
        "Content",content
    );
end
-- 客户的排队队列
local KeyOfClientQueue = function() return "cs:client:queue"; end
-- 客户的信息
local KeyOfInfoOfClient = function(clientId) return "cs:client:info:" .. clientId; end
-- 客户消息Stream
local KeyOfMessageStreamOfClient = function(clientId) return "cs:client:stream:" .. clientId; end

-- 客服的客户列表
local KeyOfClientListOfClientServer = function(clientServerId) return "cs:clientServer:clientList:" .. clientServerId; end
-- 客服信息
local KeyOfInfoOfClientServer = function(clientServerId) return "cs:clientServer:info" .. clientServerId; end
----------------------------------------------------------------------------------------
-- 清空客户的当前客服信息
local ResetCurrentClientServerOfClient = function(clientId)
    redis.call("hmset",KeyOfInfoOfClient(clientId),"ClientServerId","","ClientServerName","");
end
----------------------------------------------------------------------------------------

-- 客户ID
local clientId = ARGV[1];
local clientName = ARGV[2];
local time = redis.call("time");
local timestamp = tonumber(time[1]);

-- 加入客户ID到队列里
local Queue = function(clientId)
    redis.call("zadd",KeyOfClientQueue(),timestamp,clientId);
    redis.call("hmset",KeyOfInfoOfClient(clientId),"State",1);
    SystemMessage(KeyOfMessageStreamOfClient(clientId),"客户加入排队");
end

-- 更新当前的排队索引到客户信息
local UpdateQueueIndexOfClient = function(clientId)
    local index = redis.call("zrank",KeyOfClientQueue(),clientId);
    redis.call("hset",KeyOfInfoOfClient(clientId),"QueueIndex",index);
end

redis.replicate_commands();

-- 初始化客户信息
if redis.call("exists",KeyOfInfoOfClient(clientId)) == 0 then
    redis.call("hmset",KeyOfInfoOfClient(clientId),
        "Id",clientId,
        "Name",clientName,
        "State",0,
        "QueueIndex",0,
        "ClientServerId","",
        "ClientServerName",""
    );
end

local state = redis.call("hget",KeyOfInfoOfClient(clientId),"State");

if state == "0" then    -- 客户没有进入排队队列
    -- 客户加入排队队列里
    Queue(clientId);
    -- 更新客户的排队位置信息
    UpdateQueueIndexOfClient(clientId);
elseif state == "1" then  -- 客户正在排队中
    -- 更新客户的排队位置信息
    UpdateQueueIndexOfClient(clientId);
elseif state == "2" then  -- 客户已经分配到客服
end

return redis.call("hgetall",KeyOfInfoOfClient(clientId));