package ipam

import (
	"testing"
	"fmt"
)

func TestRedisIpam_RequestAddress(t *testing.T) {
	ri := &redisIpam{}
	poolID, _ := ri.RequestPool(`192.168.231.0/30`)
	fmt.Println(poolID)
	//address, _ := ri.RequestAddress(poolID, nil)
	//pretty.Println(address.String())
	//ri.ReleaseAddress(poolID, address.IP)
}
