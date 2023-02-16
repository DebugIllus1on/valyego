package user

import (
	"github.com/valyego/internal/zcar/biz"
	"github.com/valyego/internal/zcar/store"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
