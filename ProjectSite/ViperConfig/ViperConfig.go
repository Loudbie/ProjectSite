package ViperConfig

import (
	"github.com/spf13/viper"
)

var v = viper.New()

func CheckSet() {
	viper.SetConfigName("app") // Имя файла конфигурации без расширения

	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error " + err.Error())
	}
}
