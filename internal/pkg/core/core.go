package core

import (
	"github.com/gin-gonic/gin"
)

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		// httpcode, appcode, message := errcode.Decode(err)
		// ctx.JSON(httpcode, )
	}
}
