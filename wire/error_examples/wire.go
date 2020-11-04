//+build wireinject

package error_examples

import "github.com/google/wire"

func NewService() *Service {
	wire.Build(
		wire.Struct(new(Service), "*"),
	)
	return nil
}
