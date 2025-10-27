package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/projects/go-projects/echo-apis/go-arrs")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}