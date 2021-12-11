package ch

/**
 * GoLand
 * @author AnnnJ
 * @date 2021/12/10 15:25
 */

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
)

type ConsulHTTP struct {
	Host      string
	Port      int
	ServiceId string
}

func NewConsulHTTP(host string, port int) *ConsulHTTP {
	return &ConsulHTTP{host, port, ""}
}

// Register 服务注册
func (c *ConsulHTTP) Register(address string, port int, name string, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.Host, c.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}

	// 生成检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	// 生成注册对象
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	registration := new(api.AgentServiceRegistration)
	registration.Address = address
	registration.Port = port
	registration.Name = name
	registration.Tags = tags
	registration.ID = serviceId
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}

	c.ServiceId = serviceId
	return nil
}

// DeRegister 服务注销
func (c *ConsulHTTP) DeRegister(serviceId string) error {
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
