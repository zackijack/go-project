package config

import (
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Load() (c *Config, err error) {
	// Search env file in working directory with name ".env" if it exists.
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		zlog.Debug().Msg("Using config file: " + viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&c)

	return
}
