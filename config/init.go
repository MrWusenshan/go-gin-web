package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./application.yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
