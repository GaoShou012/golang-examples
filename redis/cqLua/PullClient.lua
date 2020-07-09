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
local Bind = function(clientInfo,clientServerInfo)
    -- 标记客户的状态信息
    redis.call("hmset",KeyOfInfoOfClient(clientInfo["Id"]),
        "State",2,
        "QueueIndex",0,
        "ClientServerId",clientServerInfo["Id"],
        "ClientServerName",clientServerInfo["Name"]
    )

    -- 保存客户ID到客户列表
    redis.call("hset",KeyOfClientListOfClientServer(clientServerInfo["Id"]),
        clientInfo["Id"],"1"
    )
end