package core

import (
	"github.com/gin-gonic/gin"
	"github.com/valyego/internal/pkg/errcode"
)

// errResponse 发生错误是返回的响应结构
type errResponse struct {
	Code    string      `json:"code"`    // 业务代码
	Message string      `json:"message"` // 错误信息
	Data    interface{} `json:"data"`    // 数据
}

// WriteResponse 错误写入到响应
func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		httpcode, appcode, message := errcode.Decode(err)
		ctx.JSON(httpcode, errResponse{
			Code:    appcode,
			Message: message,
			Data:    nil,
		})

		return
	}

	ctx.JSON(errcode.OK.HTTP, errResponse{
		Code:    errcode.OK.Code,
		Message: errcode.OK.Message,
		Data:    data,
	})
}
