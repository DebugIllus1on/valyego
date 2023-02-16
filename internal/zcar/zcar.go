package zcar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/valyego/internal/pkg/middlewares"
	"github.com/valyego/internal/zcar/router"
)

// 设置配置文件路径
var configFile = "configs/zcar.yaml"

func NewZcarAppCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 命令名
		Use: "zcar",
		// 简短描述
		Short: "简单的zcar应用示例",
		// 详细描述
		// Long: "",
		// 命令出错时不打印帮助信息
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 运行业务代码函数
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}
	// 当运行命令时调用以下函数
	// 初始化配置
	cobra.OnInitialize(initConfig)

	return cmd
}

// run 函数是实际的业务代码入口函数
func run() error {
	// 初始化 Store 数据层
	if err := initStore(); err != nil {
		return nil
	}

	// 设置 Gin 启动模式
	gin.SetMode(viper.GetString("runmode"))
	// 创建 Gin 引擎
	g := gin.New()
	// 默认使用 Gin 日志
	// 默认使用 Gin Recovery 中间件
	g.Use(gin.Logger(), gin.Recovery())

	// 加载中间件
	g.Use(middlewares.XRequestID())
	// 加载路由
	router.Routes(g)

	// 创建 HTTP 服务实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}
	// 打印 Log
	fmt.Println("[runtime] start the server")
	// 运行 HTTP 服务
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()

	// 创建退出等待信号
	quit := make(chan os.Signal, 1)
	// 等待给定的信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 等待退出信号
	<-quit

	// 创建 context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 关闭服务
	if err := httpsrv.Shutdown(ctx); err != nil {
		// TODO: 处理错误日志
		return err
	}

	return nil
}
