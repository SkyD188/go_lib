package config

import (
	"github.com/spf13/viper"
)

func InitConfig(confType, confPath string, conf interface{}) {
	v := viper.New()

	v.SetConfigType(confType) // 配置文件的类型
	v.SetConfigFile(confPath) // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

}
