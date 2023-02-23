package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Token      string
	STATIC_DIR string
	VID_PATH   string
	AUD_PATH   string
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
	log.Println(Conf.STATIC_DIR)
}
