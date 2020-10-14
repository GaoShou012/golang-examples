package config

import (
	"fmt"
	"testing"
	"wchat/utils"
)

const EtcdAddress = "192.168.56.101:2379"

func TestNewRedis(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewRedis()
	fmt.Println("redis conf:", conf)
}

func TestNewMysql(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewMysql()
	fmt.Println("mysql conf:", conf)
}

func TestNewKafka(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewKafka()
	fmt.Println("kafka conf:", conf)
}

func TestNewFrontierConfig(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewFrontierConfig()
	fmt.Println("frontier conf:", conf)
}

func TestNewChatRoomConfig(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewChatRoomConfig()
	fmt.Println("chat room conf:", conf)
}

func TestNewTenantApiConfig(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewTenantApiConfig()
	fmt.Println("tenant-api conf:", conf)
}

func TestNewTenantManagementConfig(t *testing.T) {
	utils.Micro.EtcdRegistry(EtcdAddress)
	utils.Micro.LoadSource()
	conf := NewTenantManagementConfig()
	fmt.Println("tenant-management conf:", conf)
}
