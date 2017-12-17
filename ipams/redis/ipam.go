package redis

import (
	"net"
	"fmt"
	"ipam/models"
	"ipam/ipams"
)

const (
	keyIPPool  = `ippool:%s`
)

type redisIpam struct {
}

func (ri *redisIpam) RequestPool(pool string) (poolID string, err error) {
	_, nw, err := net.ParseCIDR(pool)
	if err != nil {
		return ``, err
	}

	key := fmt.Sprintf(keyIPPool, nw.String())
	exists, err := existsKey(key)
	if err != nil {
		return ``, err
	}
	if exists {
		return nw.String(), nil
	}

	poolData := models.NewPoolData(nw)
	args := make([]interface{}, 0, poolData.NumAddresses-2)
	args = append(args, key)
	//ip地址全0和全1不分配
	for i := 1; i < poolData.NumAddresses-1; i++ {
		args = append(args, i)
	}

	return nw.String(), rpush(args...)
}

func (ri *redisIpam) ReleasePool(poolID string) error {
	return nil
}

func (ri *redisIpam) RequestAddress(poolID string, prefAddress net.IP) (*net.IPNet, error) {
	key := fmt.Sprintf(keyIPPool, poolID)
	exists, err := existsKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrKeyNotExists
	}

	ordinal, err := lpop(key)
	if err != nil {
		return nil, err
	}
	_, nw, _ := net.ParseCIDR(poolID)
	address := ipam.GenerateAddress(uint64(ordinal), nw)

	return &net.IPNet{IP: address, Mask: nw.Mask}, nil
}

func (ri *redisIpam) ReleaseAddress(poolID string, address net.IP) error {
	key := fmt.Sprintf(keyIPPool, poolID)
	exists, err := existsKey(key)
	if err != nil {
		return err
	}
	if !exists {
		return ErrKeyNotExists
	}

	_, nw, _ := net.ParseCIDR(poolID)

	h, err := ipam.GetHostPartIP(address, nw.Mask)
	if err != nil {
		return err
	}
	ordinal := ipam.IpToUint64(h)
	return rpush(key, ordinal)
}
