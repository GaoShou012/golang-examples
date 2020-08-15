package utils

import (
	"fmt"
	microconfig "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	microsourceetcd "github.com/micro/go-micro/v2/config/source/etcd"
)

var Micro micro

type micro struct {
	Config microconfig.Config
	Source source.Source

	etcd struct {
		addr     string
		registry registry.Registry
	}
}

func (m *micro) Init(srcHost string, srcPort int) {
	m.etcd.addr = fmt.Sprintf("%s:%d", srcHost, srcPort)
	m.etcd.registry = etcd.NewRegistry(
		registry.Addrs(m.etcd.addr),
	)

	conf, err := microconfig.NewConfig()
	if err != nil {
		panic(err)
	}
	m.Config = conf

	m.Source = microsourceetcd.NewSource(microsourceetcd.WithAddress(m.etcd.addr))
}

func (m *micro) GetEtcdRegistry() registry.Registry {
	return m.etcd.registry
}

func (m *micro) LoadSource() {
	if err := m.Config.Load(m.Source); err != nil {
		panic(err)
	}
}
