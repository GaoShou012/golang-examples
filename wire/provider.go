package wire

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

func NewMysql() *gorm.DB {
	return &gorm.DB{}
}
func NewRedis() *redis.ClusterClient {
	return &redis.ClusterClient{}
}
