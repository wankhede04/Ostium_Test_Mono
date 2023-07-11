package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func IntializeConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigType("json")
	viper.SetConfigName("config.json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
