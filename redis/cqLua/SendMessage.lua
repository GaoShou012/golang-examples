-- 消息加入消息stream
-- 通知所有会员列表

-- local streamName = "teststream";
-- redis.call("xadd","teststream","*","Id","12344")


redis.replicate_commands();
redis.call("xadd","teststream","*","Id","");

