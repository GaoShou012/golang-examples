//+build wireinject

package wire

import "github.com/google/wire"

var utils = wire.NewSet(NewRedis, NewMysql)

func Init() *Service {
	wire.Build(utils, wire.Struct(new(Service),"*"))
	return nil
}
