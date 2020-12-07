package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return v
}

var ProviderSet = wire.NewSet(NewConfig)
