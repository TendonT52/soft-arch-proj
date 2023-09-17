package config

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost               string `mapstructure:"POSTGRES_HOST"`
	DBUserName           string `mapstructure:"POSTGRES_USER"`
	DBUserPassword       string `mapstructure:"POSTGRES_PASSWORD"`
	DBName               string `mapstructure:"POSTGRES_DB"`
	DBPort               string `mapstructure:"POSTGRES_PORT"`
	ServerPort           string `mapstructure:"PORT"`
	AccessTokenPublicKey string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`

	AccessTokenPrivateKeyTest string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY_TEST"`
	AccessTokenPublicKeyTest  string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY_TEST"`
	AccessTokenExpiredInTest  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN_TEST"`
	AccessTokenMaxAgeTest     int           `mapstructure:"ACCESS_TOKEN_MAXAGE_TEST"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	})
	if err != nil {
		return nil, err
	}
	return
}
