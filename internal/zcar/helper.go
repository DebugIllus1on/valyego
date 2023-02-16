package zcar

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/valyego/internal/zcar/store"
	"github.com/valyego/pkg/db"
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

func initStore() error {
	dbOptions := &db.DatabaseOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := db.NewDatabase(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)
	fmt.Fprintln(os.Stderr, "[runtime] mysql connected")

	return nil
}

func initRoutes() {

}
