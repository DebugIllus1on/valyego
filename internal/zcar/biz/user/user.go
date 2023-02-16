package user

import (
	"context"

	"github.com/valyego/internal/pkg/errcode"
	"github.com/valyego/internal/zcar/store"
)

type UserBiz interface {
	Index(ctx context.Context, request interface{}) ([]interface{}, error)
	Create(ctx context.Context, request interface{}) error
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) *userBiz {
	return &userBiz{ds}
}

func (b *userBiz) Index(ctx context.Context, request interface{}) ([]interface{}, error) {
	users := b.ds.User().List()

	return users, errcode.OK
}

func (b *userBiz) Create(ctx context.Context, request interface{}) error {
	return nil
}
