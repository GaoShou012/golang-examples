-- 客服，拉取需要服务的客户

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
-- 客服ID
local clientServerId = ARGV[1];
local time = redis.call("time");
local timestamp = tonumber(time[1]);

-- 弹出客户ID
local PopClientFromQueue = function()
    local arr = redis.call("zrange",KeyOfClientQueue(),0,0);
    local clientId = arr[1];

    if clientId == nil then
        return nil;
    else
        -- 移除ClientId
        redis.call("zrem",KeyOfClientQueue(),clientId);
        return clientId;
    end
end

-- 分配客服
local Bind = function(clientInfo,clientServerInfo,timestamp)
    redis.call("hmset",KeyOfClientInfo(clientInfo["Id"]),
        "State",2,
        "QueueIndex",0,
        "ClientServerId",clientServerInfo["Id"],
        "ClientServerName",clientServerInfo["Name"]
    );
    redis.call("hset",KeyOfClientListOfClientServer(clientServerInfo["Id"]),clientInfo["Id"],timestamp);

    -- 系统消息，客服加入服务
    local content = "客服(" .. clientServerInfo["Name"] .. ")加入服务";
    SystemMessage(KeyOfMessageStreamOfClient(clientInfo["Id"]),content);
end

redis.replicate_commands();

-- 弹出一个客户
local clientId = PopClientFromQueue();
if clientId == nil then return ""; end

local clientInfo = redis.call("hmget",KeyOfInfoOfClient(clientId),"Id","Name");
local clientServerInfo = redis.call("hmget",KeyOfInfoOfClientServer(clientServerId),"Id","Name");

-- 客服分配
Bind(clientInfo,clientServerInfo,timestamp);

-- 返回客户ID
return clientId;
