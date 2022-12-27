package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	ListenAddress string
	ListenPort    string
}

type DBConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
}

func init() {
	viper.SetDefault("API.ListenAddress", "127.0.0.1")
	viper.SetDefault("API.ListenPort", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)
	cfg.API = APIConfig{
		ListenAddress: viper.GetString("API.ListenAddress"),
		ListenPort:    viper.GetString("API.ListenPort"),
	}

	cfg.DB = DBConfig{
		Database: viper.GetString("DB.Database"),
		User:     viper.GetString("DB.User"),
		Password: viper.GetString("DB.Password"),
		Host:     viper.GetString("DB.Host"),
		Port:     viper.GetString("DB.Port"),
	}
	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.ListenPort
}

func GetAPI() APIConfig {
	return cfg.API
}
