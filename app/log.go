package app

import (
	logrusredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

/*
	配置Logrus日志的流向->redis

	appName 项目|服务|模块 名称
*/
func LogrusRedisHook(appName string) error {
	host := Configs.LogRedis.Host
	pass := Configs.LogRedis.Password
	port := Configs.LogRedis.Port
	key := Configs.LogRedis.Key
	db := Configs.LogRedis.DB

	conf := logrusredis.HookConfig{
		Key:      key,
		Format:   "v1",
		App:      appName,
		Host:     host,
		Password: pass,
		Hostname: "",
		Port:     port,
		DB:       db,
		TTL:      3600,
	}

	hook, err := logrusredis.NewHook(conf)
	if err != nil {
		return err
	}

	logrus.AddHook(hook)
	return nil
}
