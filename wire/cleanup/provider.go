package cleanup

import "fmt"

type Player struct {
	Name string
}
type Monster struct {
	Name string
}

type Mission struct {
	Player  *Player
	Monster *Monster
}

func NewPlayer() (*Player, func(), error) {
	p := &Player{Name:"p1"}

	cleanup := func() {
		fmt.Println("new player cleanup")
	}

	return p, cleanup, nil
}

func NewMonster() (*Monster, func(), error) {
	p := &Monster{Name:"m123"}

	cleanup := func() {
		fmt.Println("new monster cleanup")
	}

	return p, cleanup, nil
}

func NewMission(player *Player, monster *Monster) (*Mission, func(), error) {
	p := &Mission{
		Player:  player,
		Monster: monster,
	}

	cleanup := func() {
		fmt.Println("new mission cleanup")
	}

	return p, cleanup, nil
}
