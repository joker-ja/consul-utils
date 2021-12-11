package ci

/**
 * GoLand
 * @author AnnnJ
 * @date 2021/12/10 15:40
 */

type ConsulClient interface {
	Register(address string, port int, name, string, tags []string, id string) error
	DeRegister(serviceId string) error
}
