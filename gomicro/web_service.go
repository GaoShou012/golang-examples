package main

import (
	"fmt"
	"github.com/config"
	"github.com/micro/go-micro/v2/web"
	"github.com/utils"
	"time"
)

func main() {
	utils.Micro.Init("192.168.1.38", 2379)
	utils.Micro.LoadSource()
	config.TenantServiceConfig.LoadFromMicroSourceEtcd()

	//go func() {
	//	service := web.NewService(
	//		web.Name(config.TenantServiceConfig.GetServiceName()),
	//		web.Address(fmt.Sprintf(":%d", config.TenantServiceConfig.Port)),
	//		web.Registry(utils.Micro.GetEtcdRegistry()),
	//		web.RegisterTTL(time.Second*10),
	//	)
	//	if err := service.Run(); err != nil {
	//		panic(err)
	//	}
	//}()

	service := web.NewService(
		web.Name(config.TenantServiceConfig.GetServiceName()),
		web.Address(fmt.Sprintf(":%d", config.TenantServiceConfig.Port)),
		web.Registry(utils.Micro.GetEtcdRegistry()),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)
	if err := service.Run(); err != nil {
		panic(err)
	}


}
