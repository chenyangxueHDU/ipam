package models

import "net"

type PoolData struct {
	ID           string
	NumAddresses int
	IP           string
	Mask         string
}

type poolDataDao struct{}

func NewPoolData(nw *net.IPNet) *PoolData {
	ones, bits := nw.Mask.Size()
	return &PoolData{
		ID:           nw.String(),
		NumAddresses: int(1 << uint(bits-ones)),
		IP:           nw.IP.String(),
		Mask:         nw.Mask.String(),
	}
}

func NewPoolDataDao() PoolDataDao {
	return &poolDataDao{}
}

func (dao *poolDataDao) Insert(pool *PoolData) {

}
func (dao *poolDataDao) Get(id int) *PoolData {

	return nil
}
func (dao *poolDataDao) GetAll() []*PoolData {

	return nil
}
