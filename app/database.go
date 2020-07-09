package app

import (
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v7"
)

func Gorm() (*gorm.DB, error) {
	user := Configs.DB.User
	pass := Configs.DB.Password
	host := Configs.DB.Host
	defaultDatabase := Configs.DB.DefaultDatabase

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		user,
		pass,
		host,
		defaultDatabase,
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Redis() *redis.Client {
	host := Configs.Redis.Host
	port := Configs.Redis.Port
	user := Configs.Redis.User
	pass := Configs.Redis.Password
	db := Configs.Redis.DB

	r := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               fmt.Sprintf("%s:%d", host, port),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           user,
		Password:           pass,
		DB:                 db,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
	return r
}
