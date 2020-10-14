//+build wireinject

package empty_struct

import "github.com/google/wire"

func NewService() *Service {
	wire.Build(
		Provider,
	)
	return nil
}
