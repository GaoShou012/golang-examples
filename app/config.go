package app

import "github.com/jinzhu/configor"

const (
	configFilePath = "E:\\work\\golang\\examples\\app\\config.yml"
)

var Configs struct {
	Etcd struct {
		Host string
		Port int
	}
	DB struct {
		User            string
		Password        string
		Host            string
		Port            string
		DefaultDatabase string
		Dns             string
	}
	Redis struct {
		Host     string
		User     string
		Port     int
		Password string
		DB       int
	}
	RabbitMq struct {
		User     string
		Password string
		Host     string
		Port     int
	}
	LogRedis struct {
		Host     string
		Port     int
		Password string
		DB       int
		Key      string
	}
}

func LoadConfig() error {
	// 加载配置文件
	err := configor.Load(&Configs, configFilePath)
	if err != nil {
		return err
	}
	return nil
}
