package biz

import (
	"github.com/valyego/internal/zcar/biz/user"
	"github.com/valyego/internal/zcar/store"
)

type IBiz interface {
	// TODO: 业务层接口实现
	User() user.UserBiz
}

var _ IBiz = (*biz)(nil)

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) *biz {
	return &biz{ds}
}

func (b *biz) User() user.UserBiz {
	return user.New(b.ds)
}
