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

local clientId = ARGV[1];
local clientName = ARGV[2];
local time = redis.call("time");
local timestamp = tonumber(time[1]);

-- 等级用户信息
redis.call("hmset",keyOfInfoOfClient(clientId),
    "Id",clientId,
    "Name",clientName
    "LastLoginTime",timestamp
);
