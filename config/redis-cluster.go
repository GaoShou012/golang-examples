package config

import (
	"examples/utils"
)

var (
	_                  utils.MicroConfig = &redisClusterConfig{}
	RedisClusterConfig                   = &redisClusterConfig{}
)

type redisClusterConfig struct {
	Addr     []string
	Password string
}

func (c *redisClusterConfig) ConfigKey() string {
	return "redis-cluster"
}
