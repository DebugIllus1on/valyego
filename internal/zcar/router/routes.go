package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const GroupName = "zcar"

// Routes 路由
func Routes(ctx *gin.Engine) {
	// 注册 404 Handler
	ctx.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "100001", "message": "router not found.",
		})
	})

	// 路由组
	appgroup := ctx.Group(GroupName)

	// 注册路由
	appgroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1", "message": "There is [Zcar] app",
		})
	})

	appgroup.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
