package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgreDriver  string `mapstructure:"POSTGRES_DRIVER"`
	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`

	Port string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
