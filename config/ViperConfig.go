package config

import "github.com/spf13/viper"

func ViperConfig() {
	viper.SetConfigName("application-dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources/")
	e := viper.ReadInConfig()
	PrintError(e, "viper 配置错误")
}
