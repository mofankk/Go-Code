package config

import (
	"github.com/spf13/viper"
)

func New(cfgName string) (*viper.Viper, error) {
	cfgName += "_local"
	v := viper.New()
	v.SetConfigName(cfgName)
	v.SetConfigType("yaml")
	v.AddConfigPath("../config")
	v.AddConfigPath(".")
	v.AddConfigPath("./")
	v.AddConfigPath("./config")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
