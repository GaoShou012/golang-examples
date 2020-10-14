package config

import (
	"github.com/google/wire"
	uuid "github.com/satori/go.uuid"
	"wchat/utils"
)

var Provider = wire.NewSet(
	NewRedis,
	NewMysql,
	NewKafka,
	NewFrontierConfig,
	NewChatRoomConfig,
	NewTenantApiConfig,
	NewTenantManagementConfig,
)

func NewRedis() *RedisConf {
	conf := &RedisConf{}
	utils.Micro.LoadConfigMust(conf, "redis-cluster")
	return conf
}

func NewMysql() *MysqlConf {
	conf := &MysqlConf{}
	utils.Micro.LoadConfigMust(conf, "mysql")
	return conf
}

func NewKafka() *KafkaClusterConfig {
	conf := &KafkaClusterConfig{}
	utils.Micro.LoadConfigMust(conf, "kafka-cluster")
	return conf
}

func NewFrontierConfig() *FrontierConfig {
	conf := &FrontierConfig{}
	utils.Micro.LoadConfigMust(conf, "frontier")
	conf.FrontierId = uuid.NewV4().String()
	return conf
}

func NewChatRoomConfig() *ChatRoomServiceConfig {
	conf := &ChatRoomServiceConfig{}
	utils.Micro.LoadConfigMust(conf, "chat-room-service")
	return conf
}

func NewTenantApiConfig() *TenantApi {
	conf := &TenantApi{}
	utils.Micro.LoadConfigMust(conf, "tenant-api-service")
	return conf
}

func NewTenantManagementConfig() *TenantManagement {
	conf := &TenantManagement{}
	utils.Micro.LoadConfigMust(conf, "tenant-management-service")
	return conf
}
