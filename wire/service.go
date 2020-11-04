package wire

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Service struct {
	//Addr     string
	Redis    *redis.ClusterClient
	DB       *gorm.DB
	ServiceA *ServiceA
}

type ServiceA struct {
}
