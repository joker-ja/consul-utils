package ch

import "testing"

/**
 * GoLand
 * @author AnnnJ
 * @date 2021/12/10 15:58
 */

func TestConsulHTTP_Register(t *testing.T) {
	c := NewConsulHTTP("192.168.101.30", 8500)
	err := c.Register("192.168.101.23", 8021, "annnj-http", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConsulHTTP_DeRegister(t *testing.T) {
	c := NewConsulHTTP("192.168.101.30", 8500)
	err := c.DeRegister("f7fb84c6-ef9f-482c-bba6-6f6e5cbe069c")
	if err != nil {
		t.Fatal(err)
	}
}
