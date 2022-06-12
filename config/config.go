package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	var configFile = os.Getenv("APP_ENV_FILE")
	if os.Getenv("APP_ENV_FILE") == "" {
		configFile = ".env"
	}
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
