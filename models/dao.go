package models

type PoolInfoDao interface {
	Insert(pool *PoolInfo) error
	Get(id string) (*PoolInfo, error)
	GetAll() []*PoolInfo
}
