package config

import (
	"github.com/spf13/viper"
)

func InitConfig(confPath string, conf interface{}) {
	v := viper.New()

	v.SetConfigFile(confPath) // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

}
