package models

import (
	"testing"
	"net"
	"github.com/kr/pretty"
)

func TestNewPoolData(t *testing.T) {
	_, nw, _ := net.ParseCIDR(`192.168.232.1/24`)
	poolData := NewPoolData(nw)
	pretty.Println(poolData)
	pretty.Println(nw.String())
}
