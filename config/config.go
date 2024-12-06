package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string
	DBSQLite struct {
		Filename string
	}
	PortSrvUsers    int
	PortSrvConcerts int
}

func New() *Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config
	viper.Unmarshal(&config)
	log.Println(config)
	return &config
}
