package example

import (
	"github.com/gin-gonic/gin"
	"github.com/valyego/internal/pkg/core"
	"github.com/valyego/internal/pkg/errcode"
)

func (ctrl *Example) Index(ctx *gin.Context) {
	data := []string{
		"Recorder updates", "Customize selector types of a recording", "Edit user flow while recording",
	}

	core.WriteResponse(ctx, errcode.ErrCarNotFound, data)
}
