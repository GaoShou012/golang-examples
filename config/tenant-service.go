package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/utils"
)

var TenantServiceConfig = &tenantServiceConfig{serviceName: "tenant-service"}

type tenantServiceConfig struct {
	serviceName string
	Port        int `json:"port"`
}

func (c *tenantServiceConfig) LoadFromMicroSourceEtcd() {
	val := utils.Micro.Config.Get("micro", "config", fmt.Sprintf("%s-config",c.serviceName))
	if bytes.Equal(val.Bytes(), []byte("null")) {
		panic(fmt.Sprintf("tenant-service config is %v", string(val.Bytes())))
	}
	if err := json.Unmarshal(val.Bytes(), c); err != nil {
		panic(err)
	}
}

func (c *tenantServiceConfig) GetServiceName() string {
	return c.serviceName
}
