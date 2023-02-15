package example

import (
	"github.com/gin-gonic/gin"
	"github.com/valyego/internal/pkg/core"
	"github.com/valyego/internal/pkg/errcode"
)

func (ctrl *Example) Create(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(nil); err != nil {
		core.WriteResponse(ctx, errcode.ErrResourceNotFound, nil)
		return
	}
}
