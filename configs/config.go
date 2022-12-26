package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	ListenAddress string
	ListenPort    int
}

type DBConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     int
}

func init() {
	viper.SetDefault("API.ListenAddress", "127.0.0.1")
}
