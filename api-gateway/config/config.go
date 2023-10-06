package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	RESTPort       string `mapstructure:"REST_PORT"`
	UserServiceHost string `mapstructure:"USER_SERVICE_HOST"`
	UserServicePort string `mapstructure:"USER_SERVICE_PORT"`
	PostServiceHost string `mapstructure:"POST_SERVICE_HOST"`
	PostServicePort string `mapstructure:"POST_SERVICE_PORT"`
	ReportServiceHost string `mapstructure:"REPORT_SERVICE_HOST"`
	ReportServicePort string `mapstructure:"REPORT_SERVICE_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	})
	if err != nil {
		return nil, err
	}

	return config, nil
}
