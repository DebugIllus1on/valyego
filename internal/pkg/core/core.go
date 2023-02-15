package core

import (
	"github.com/gin-gonic/gin"
	"github.com/valyego/internal/pkg/errcode"
)

// ErrResponse 发生错误是返回的响应结构
type ErrResponse struct {
	Code    string `json:"code"`    // 业务代码
	Message string `json:"message"` // 错误信息
}

// WriteResponse 错误写入到响应
func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		httpcode, appcode, message := errcode.Decode(err)
		ctx.JSON(httpcode, ErrResponse{
			Code:    appcode,
			Message: message,
		})

		return
	}

	ctx.JSON(errcode.OK.HTTP, ErrResponse{
		Code:    errcode.OK.Code,
		Message: errcode.OK.Message,
	})
}
