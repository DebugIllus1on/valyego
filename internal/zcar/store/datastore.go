package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	D    *datastore
)

// datastore 是 IStore 的具体实现
type datastore struct {
	dbc *gorm.DB
}

// NewStore 创建 IStore 实例
func NewStore(dbc *gorm.DB) *datastore {
	once.Do(func() {
		D = &datastore{dbc}
	})

	return D
}

// 检查是否已经实现 IStore
var _ IStore = (*datastore)(nil)

func (ds *datastore) User() UserStore {
	return newUser(ds.dbc)
}
