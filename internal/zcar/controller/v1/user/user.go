package user

import (
	"github.com/gin-gonic/gin"
	"github.com/valyego/internal/pkg/core"
	"github.com/valyego/internal/zcar/biz"
	"github.com/valyego/internal/zcar/store"
)

type UserController struct {
	biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{biz.NewBiz(ds)}
}

func (ctrl *UserController) Index(c *gin.Context) {
	users, err := ctrl.User().Index(c, nil)
	core.WriteResponse(c, err, users)
}
