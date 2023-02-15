package store

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
	DS   *datastore
)

// IStore 定义 Store 层实现的方法
type IStore interface {
	User() // User 用户
}

// datastore IStore 的具体实现
type datastore struct {
	db *gorm.DB
}

var _ IStore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型实例
func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		DS = &datastore{db}
	})

	return DS
}

// Users 返回实现 User 的实例
func (ds *datastore) User() {
	return
}
