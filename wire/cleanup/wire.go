// +build wireinject

package cleanup

import "github.com/google/wire"

func New() (*Mission, func(), error) {
	wire.Build(
		NewPlayer,
		NewMonster,
		NewMission,
	)
	return nil, nil, nil
}
