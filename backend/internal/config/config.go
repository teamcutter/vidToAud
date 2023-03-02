package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Token      string
	VID_PATH   string
	AUD_PATH   string
	DEBUG_PORT string
	PROD_PORT  string
}

var Conf Configuration

func GetConfig() {
	conf := Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./internal/config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	Conf = conf
}
