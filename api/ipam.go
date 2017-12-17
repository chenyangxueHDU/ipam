package api

import "net"

type Ipam interface {
	RequestPool(pool string) (string, error)
	// ReleasePool releases the address pool identified by the passed id
	ReleasePool(poolID string) error
	// Request address from the specified pool ID. Input options or required IP can be passed.
	RequestAddress(string, net.IP) (*net.IPNet, error)
	// Release the address from the specified pool ID
	ReleaseAddress(string, net.IP) error
}
