package main

import (
	"os"

	"github.com/valyego/internal/zcar"
	_ "go.uber.org/automaxprocs"
)

// 程序默认入口函数
func main() {
	// 执行 zcar 应用程序命令
	if err := zcar.NewCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
