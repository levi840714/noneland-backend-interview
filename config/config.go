package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	// config file name
	configName = "config"
	// config file paths
	configPaths = []string{
		"./",
		"./config",
	}
)

type conf struct {
	Debug bool
	Host  string
	Port  string
}

var Conf conf

func Init(name ...string) {
	if len(name) > 0 {
		configName = name[0]
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("load config err: %v", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("could not unmarshal config: %v", err)
	}

	log.Printf("config: %+v", Conf)
}
