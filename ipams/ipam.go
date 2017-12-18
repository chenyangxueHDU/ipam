package ipams

import (
	"net"
	"fmt"
	"ipam/models"
	"ipam/cache"
	"github.com/garyburd/redigo/redis"
)

type redisIpam struct {
}

var poolInfoDao models.PoolInfoDao

const (
	availableBit   = 0
	unavailableBit = 1
)

func init() {
	poolInfoDao = models.NewPoolInfoDao()
}

func (ri *redisIpam) RequestPool(pool string) (poolID string, err error) {
	_, nw, err := net.ParseCIDR(pool)
	if err != nil {
		return ``, err
	}
	key := fmt.Sprintf(cache.KeyPoolData, nw.String())
	exists, err := cache.ExistsKey(key)
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

	return nw.String(), cache.NewBitmap(key)
}

func (ri *redisIpam) ReleasePool(poolID string) error {
	return nil
}

func (ri *redisIpam) RequestAddress(poolID string, prefAddress net.IP) (*net.IPNet, error) {
	info, err := poolInfoDao.Get(poolID)

	ordinal, err := ri.getNextAvailable(poolID, info.NumAddresses)
	if err != nil {
		return nil, err
	}

	_, nw, _ := net.ParseCIDR(poolID)
	address := GenerateAddress(uint64(ordinal), nw)

	return &net.IPNet{IP: address, Mask: nw.Mask}, nil
}

func (ri *redisIpam) getNextAvailable(poolID string, numAddresses int) (int, error) {

	conn := cache.GetConn()
	defer conn.Close()
	key := fmt.Sprintf(cache.KeyPoolData, poolID)

	//todo use lua
	next, err := redis.Int(conn.Do(`BITPOS`, key, availableBit))
	if err != nil {
		return 0, err
	}
	//从0开始
	if next >= numAddresses {
		return 0, ErrPoolEmpty
	}

	conn.Do(`SETBIT`, key, next, unavailableBit)

	return next + 1, err
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
