package sets

import "github.com/google/wire"

var Provider = wire.NewSet()

var ProviderService = wire.NewSet(NewServiceA,NewServiceB,NewServiceC)

func NewConf() *Conf {
	return nil
}

func NewServiceA(conf *Conf) (*ServiceA, func(), error) {
	return nil, nil, nil
}

func NewServiceB(conf *Conf) (*ServiceB, func(), error) {
	return nil, nil, nil
}

func NewServiceC(conf *Conf) (*ServiceC, func(), error) {
	return nil, nil, nil
}
