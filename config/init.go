package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("./config/app.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

type AutoConfig struct {
}

func (autoConfig AutoConfig) GetConfigByString(name string) (res string) {
	return viper.GetString(name)
}
