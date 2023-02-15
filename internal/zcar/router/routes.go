package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valyego/internal/pkg/core"
	"github.com/valyego/internal/zcar/controller/v1/example"
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
	exampleCtrl := example.New()
	appgroup.GET("/example", exampleCtrl.Index)

	appgroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "1", "message": "There is [Zcar] app",
		})
	})

	appgroup.GET("/healthz", func(ctx *gin.Context) {
		// ! 响应器测试通过即将删除
		core.WriteResponse(ctx, nil, []interface{}{
			"hello", "world",
		})
	})
}
