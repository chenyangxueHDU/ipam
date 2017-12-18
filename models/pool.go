package models

import (
	"net"
	"ipam/cache"
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type PoolInfo struct {
	ID           string
	NumAddresses int
	IP           string
	Mask         string
}

type poolInfoDao struct{}

func NewPoolInfo(nw *net.IPNet) *PoolInfo {
	ones, bits := nw.Mask.Size()
	return &PoolInfo{
		ID:           nw.String(),
		NumAddresses: int(1<<uint(bits-ones)) - 2, //ip地址个数，全0全1不分配
		IP:           nw.IP.String(),
		Mask:         nw.Mask.String(),
	}
}

func NewPoolInfoDao() PoolInfoDao {
	return &poolInfoDao{}
}

func (dao *poolInfoDao) Insert(pool *PoolInfo) error {
	conn := cache.GetConn()
	defer conn.Close()

	bs, _ := json.Marshal(pool)
	_, err := conn.Do(`SET`, fmt.Sprintf(cache.KeyPoolInfo, pool.ID), bs)
	return err
}
func (dao *poolInfoDao) Get(id string) (*PoolInfo, error) {
	conn := cache.GetConn()
	defer conn.Close()
	info := new(PoolInfo)

	s, err := redis.String(conn.Do(`GET`, fmt.Sprintf(cache.KeyPoolInfo, id)))
	if err != nil {
		return nil, err
	}

	return info, json.Unmarshal([]byte(s), info)
}
func (dao *poolInfoDao) GetAll() []*PoolInfo {

	return nil
}
