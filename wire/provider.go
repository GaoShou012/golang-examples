package wire

import (
	"github.com/go-redis/redis/v7"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var Provider = wire.NewSet(NewRedis, NewMysql)

func NewMysql() *gorm.DB {
	return &gorm.DB{}
}
func NewRedis() *redis.ClusterClient {
	return &redis.ClusterClient{}
}

func ServiceConfig(service *Service) {

}

type Movie struct {
	Director string
}

func NewMovie(director string) *Movie {
	return &Movie{Director: director}
}

type Director struct {
	Name string
}

func NewDirector(name string) *Director {
	return &Director{Name: name}
}
func NewDirector1() *Director {
	return &Director{Name: "123"}
}

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  *Player
	Monster *Monster
}

func NewMission() Mission {
	p := &Player{Name: "dj"}
	m := &Monster{Name: "kitty"}

	return Mission{p, m}
}
