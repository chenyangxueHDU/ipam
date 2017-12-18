package models

type PoolInfoDao interface {
	Insert(pool *PoolInfo) error
	Get(id int) (*PoolInfo, error)
	GetAll() []*PoolInfo
}
