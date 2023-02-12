package zcar

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 设置配置文件路径
var configFile = "configs/zcar.yaml"

func NewCommand() *cobra.Command {
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

	return nil
}
