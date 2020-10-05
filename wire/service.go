package wire

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Redis *redis.ClusterClient
	DB    *gorm.DB
}

func NewService(redis *redis.ClusterClient) *Service {
	return &Service{Redis: redis}
}
