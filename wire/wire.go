//+build wireinject

package wire

import "github.com/google/wire"

func NewService() *Service {
	wire.Build(
		//wire.Value(Service{Addr: "12321"}),
		Provider, wire.Struct(new(ServiceA), "*"),
		wire.Struct(new(Service), "*"),
	)
	return nil
}

type Foo struct {
	FX
}

type Fabc struct {
	FX
}

type FX int

var pset = wire.NewSet(wire.Value(FX(123)))
var p1set = wire.NewSet(wire.Struct(new(Foo), "*"))

func injectFoo() *Foo {
	wire.Build(
		//p1set,
		//wire.Value(422),
		pset,
		p1set,
		//wire.Struct(new(Foo),"*"),
	)
	return nil
}

func newFABC() *Fabc {
	wire.Build(
		pset,
		wire.Struct(new(Fabc), "*"),
	)
	return nil
}

// wire.go
func InitPlayer() *Player {
	wire.Build(NewMission, wire.FieldsOf(new(Mission), "Player"))
	return nil
}

func InitMonster() *Monster {
	wire.Build(
		NewMission,
		wire.FieldsOf(new(Mission), "Monster"),
	)
	return nil
}

var Set1 = wire.NewSet(NewDirector1)

func initMovie() *Movie {
	wire.Build(
		Set1,
		wire.FieldsOf(new(*Director), "Name"),
		wire.Struct(new(Movie), "*"),
		//NewMovie,
	)
	return nil
}
