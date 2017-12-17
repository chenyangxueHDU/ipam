package models

type PoolDataDao interface {
	Insert(pool *PoolData)
	Get(id int) *PoolData
	GetAll() []*PoolData
}
