package config

import (
	"fmt"

	"ttemplate/db"

	"github.com/spf13/viper"
)

type Server struct {
	GinMode string
	Address string
	Port    string
}

type DataBase struct {
	Url string
}

type Redis struct {
	Address     string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type Config struct {
	Server   Server
	Database DataBase
	Redis    Redis
}

var Configuration *Config

func LoadAndInitConfig() {
	toml := viper.New()
	toml.SetConfigName("config")
	toml.SetConfigType("toml")
	toml.AddConfigPath(".")
	err := toml.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var Config Config
	if err := toml.Unmarshal(&Config); err != nil {
		panic(err)
	}
	Configuration = &Config
	db.Init(Configuration.Database.Url)
}
