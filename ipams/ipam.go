package ipam

import (
	"net"
	"fmt"
	"ipam/models"
	"ipam/common"
	"github.com/docker/libnetwork/ipam"
)

type redisIpam struct {
}

var poolInfoDao models.PoolInfoDao

func init() {
	poolInfoDao = models.NewPoolInfoDao()
}

func (ri *redisIpam) RequestPool(pool string) (poolID string, err error) {
	_, nw, err := net.ParseCIDR(pool)
	if err != nil {
		return ``, err
	}
	key := fmt.Sprintf(common.KeyPoolData, nw.String())
	exists, err := common.ExistsKey(key)
	if err != nil {
		return ``, err
	}
	if exists {
		return nw.String(), nil
	}

	info := models.NewPoolInfo(nw)
	if err = poolInfoDao.Insert(info); err != nil {
		return ``, err
	}

	return nw.String(), common.NewBitmap(key)
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
//
//func (ri *redisIpam) ReleaseAddress(poolID string, address net.IP) error {
//	key := fmt.Sprintf(keyIPPool, poolID)
//	exists, err := existsKey(key)
//	if err != nil {
//		return err
//	}
//	if !exists {
//		return ErrKeyNotExists
//	}
//
//	_, nw, _ := net.ParseCIDR(poolID)
//
//	h, err := ipam.GetHostPartIP(address, nw.Mask)
//	if err != nil {
//		return err
//	}
//	ordinal := ipam.IpToUint64(h)
//	return rpush(key, ordinal)
//}
