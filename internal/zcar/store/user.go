package store

import (
	"gorm.io/gorm"
)

// UserStore 定义了 user 在 Store 层需实现的方法
type UserStore interface {
	Get() error
	List() []interface{}
	Store() error
	Update() error
	Destroy() error
}

// user UserStore 接口的实现
type user struct {
	*gorm.DB
}

// newUser 实例化
func newUser(dbc *gorm.DB) *user {
	return &user{dbc}
}

var _ UserStore = (*user)(nil)

func (u *user) Get() error {
	return nil
}

func (u *user) List() []interface{} {
	users := []interface{}{
		"valye", "illusion", "hello",
	}

	return users
}

func (u *user) Store() error {
	return nil
}

func (u *user) Update() error {
	return nil
}

func (u *user) Destroy() error {
	return nil
}
