package cg

import (
	"google.golang.org/grpc"
	"testing"
)

/**
 * GoLand
 * @author AnnnJ
 * @date 2021/12/10 16:16
 */

func TestConsulGRPC_Register(t *testing.T) {
	c := NewConsulGRPC("192.168.101.30", 8500)
	server := grpc.NewServer()
	c.RegisterHealthCheck(server)
	err := c.Register("192.168.101.23", 50051, "annnj-grpc", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConsulGRPC_DeRegister(t *testing.T) {
	c := NewConsulGRPC("192.168.101.30", 8500)
	err := c.DeRegister("d4d3a713-6f30-4c87-9a95-9c8fd5444a0f")
	if err != nil {
		t.Fatal(err)
	}
}
