package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()

	config.AddConfigPath("./configs/")
	config.AddConfigPath("./")
	config.SetConfigName("config.yaml")
	config.SetConfigType("yaml")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
