package server

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	"kr-consul/internal/conf"
	"log"
)

// NewConsulServer 使用Consul作为注册中心
func NewConsulServer(c *conf.Server) registry.Registrar {
	// 创建consul客户端
	configs := api.DefaultConfig()
	// 从conf/conf.proto定义Consul配置与configs/config.yml配置文件中读取consul的配置
	configs.Address = c.Consul.Addr

	// 创建consul客户端
	consulClient, err := api.NewClient(configs)
	if err != nil {
		log.Fatal(err)
	}

	// 创建consul注册中心
	r := consul.New(consulClient)
	return r
}
