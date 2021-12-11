package cg

import (
	"fmt"
	"google.golang.org/grpc"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

/**
 * GoLand
 * @author AnnnJ
 * @date 2021/12/10 15:41
 */

type ConsulGRPC struct {
	Host string
	Port int
}

func NewConsulGRPC(host string, port int) *ConsulGRPC {
	return &ConsulGRPC{host, port}
}

// Register 服务注册
func (c *ConsulGRPC) Register(address string, port int, name string, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.Host, c.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}

	// 生成检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Address = address
	registration.Port = port
	registration.Name = name
	registration.Tags = tags
	registration.ID = fmt.Sprintf("%s", uuid.NewV4())
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}

	return nil
}

// DeRegister 服务注销
func (c *ConsulGRPC) DeRegister(serviceId string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.Host, c.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}

	err = client.Agent().ServiceDeregister(serviceId)
	if err != nil {
		return err
	}
	return nil
}

// RegisterHealthCheck 注册健康检查
func (c *ConsulGRPC) RegisterHealthCheck(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
}
