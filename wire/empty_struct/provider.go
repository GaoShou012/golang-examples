package empty_struct

import "github.com/google/wire"

var Provider = wire.NewSet(wire.Struct(new(Service),"*"))

type Service struct {
}