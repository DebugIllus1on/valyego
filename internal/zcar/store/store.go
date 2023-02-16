package store

// IStore 定义需要实现的方法
type IStore interface {
	User() UserStore
}
