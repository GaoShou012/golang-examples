//+build wireinject

package sets

import "github.com/google/wire"

func NewServiceCA() (*ServiceC, func(), error) {
	wire.Build(
		NewConf,
		ProviderService,
	)
	return nil, nil, nil
}
