package exists_struct

import "github.com/google/wire"

type Sender struct {
}

type Container struct {
	Sender *Sender
}

func NewContainer() *Container {
	return nil
}

type Service1 struct {
	Sender *Sender
}

var Provider = wire.NewSet(
	//NewContainer,
	//wire.Struct(new(Sender), "*"),
	//wire.FieldsOf(new(*Container), "*"),
	wire.Value(&Container{}),
	wire.Struct(new(Service1), "*"),
)
