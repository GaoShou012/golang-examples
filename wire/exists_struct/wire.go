package exists_struct

import "github.com/google/wire"

func New() *Service1 {
	wire.Build(
		Provider,
		)
	return nil
}
