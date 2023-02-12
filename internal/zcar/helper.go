package zcar

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const ()

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量前缀为 ZCAR
	// 若是 zcar 将自动转变为大写
	viper.SetEnvPrefix("ZCAR")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "[runtime] config not found >>", viper.ConfigFileUsed())
	} else {
		// 打印当前使用的配置文件
		fmt.Fprintln(os.Stdout, "[runtime] using config file >>", viper.ConfigFileUsed())
	}
}
