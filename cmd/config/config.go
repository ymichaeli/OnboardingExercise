package config

import (
	"github.com/spf13/viper"
)

type DBConnectionInfo struct {
	Host     string `mapstructure:"HOST"`
	Port     int    `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

type ServerInfo struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

type Config struct {
	DBConfig   DBConnectionInfo `mapstructure:"DB_CONNECTION_INFO"`
	ServerInfo ServerInfo       `mapstructure:"SERVER"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("./cmd/config")
	vp.AutomaticEnv()

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
