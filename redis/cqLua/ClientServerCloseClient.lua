-- 客服解绑客户

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
local KeyOfClientQueue = function() return "cs:clientQueue"; end
-- 客户的信息
local KeyOfInfoOfClient = function(clientId) return "cs:client:info:" .. clientId; end
-- 客户消息Stream
local KeyOfMessageStreamOfClient = function(clientId) return "cs:clientStream:" .. clientId; end

-- 客服信息
local KeyOfInfoOfClientServer = function(clientServerId) return "cs:clientServer:info" .. clientServerId; end
-- 客服的客户列表
local KeyOfClientListOfClientServer = function(clientServerId) return "cs:clientServer:clientList:" .. clientServerId; end
----------------------------------------------------------------------------------------

local clientServerId = ARGV[1];
local clientId = ARGV[2];

local Unbind = function(clientServerInfo,clientInfo)
    -- 从客户列表中，移除客户
    redis.call("hdel",KeyOfClientListOfClientServer(clientServerInfo["Id"]),clientInfo["Id"]);
    -- 设置客户的状态
    -- 清空客户的当前客服信息
    redis.call("hmset",KeyOfInfoOfClient(clientInfo["Id"]),"State",0,"ClientServerId","","ClientServerName","");
    -- 系统消息
    local content = "The client server close the server(ID=" .. clientServerInfo["Id"] .. ",Name=" .. clientServerInfo["Name"] .. ")";
    SystemMessage(KeyOfMessageStreamOfClient(clientId),content);
end

local clientServerInfo = redis.call("hgetall",KeyOfInfoOfClientServer);
local clientInfo = redis.call("hgetall",KeyOfInfoOfClient);
if redis.call("hexists",KeyOfClientListOfClientServer(clientServerInfo["Id"]),clientInfo["Id"]) == 1 then
    Unbind(clientServerInfo,clientInfo);
    return {0,""}
else
    return {1,""}
end
