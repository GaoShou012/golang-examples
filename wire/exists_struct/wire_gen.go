// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package exists_struct

// Injectors from wire.go:

func New() *Service1 {
	sender := &Sender{}
	service1 := &Service1{
		Sender: sender,
	}
	return service1
}
