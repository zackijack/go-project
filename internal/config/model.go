package config

import "github.com/spf13/viper"

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	AppHost string `mapstructure:"APP_HOST"`
	AppPort int    `mapstructure:"APP_PORT"`
}

func init() {
	// Set config default value.
	viper.SetDefault("APP_NAME", "go-project")
	viper.SetDefault("APP_HOST", "0.0.0.0")
	viper.SetDefault("APP_PORT", 8080)
}
