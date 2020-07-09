package app

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	microconfig "github.com/micro/go-micro/v2/config"
	microetcd "github.com/micro/go-micro/v2/config/source/etcd"

	"sync"
)

var microServices struct {
	init   sync.Once
	client client.Client

	etcd struct {
		init     sync.Once
		registry registry.Registry
	}
}

func EtcdAddr() string {
	url := fmt.Sprintf("%s:%d", Configs.Etcd.Host, Configs.Etcd.Port)
	return url
}

func EtcdRegistry() registry.Registry {
	microServices.etcd.init.Do(func() {
		microServices.etcd.registry = etcd.NewRegistry(
			registry.Addrs(EtcdAddr()),
		)
	})
	return microServices.etcd.registry
}

func EtcdSrouce() (microconfig.Config, source.Source, error) {
	conf, err := microconfig.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	src := microetcd.NewSource(microetcd.WithAddress(EtcdAddr()))
	err = conf.Load(src)
	if err != nil {
		return nil, nil, err
	}
	return conf, src, nil
}

func MicroConfig() (microconfig.Config, error) {
	conf, err := microconfig.NewConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}
func MicroSource()(source.Source,error){
	src := microetcd.NewSource(microetcd.WithAddress(EtcdAddr()))
	return src,nil
}

func ServiceClient() client.Client {
	microServices.init.Do(func() {
		service := micro.NewService(
			micro.Registry(EtcdRegistry()),
		)
		service.Init()
		microServices.client = service.Client()
	})
	return microServices.client
}
