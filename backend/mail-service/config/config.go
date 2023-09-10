package config

import "github.com/spf13/viper"

type Config struct {
	MemphisHostName        string `mapstructure:"MEMPHIS_HOSTNAME"`
	MemphisApplicationUser string `mapstructure:"MEMPHIS_APPLICATION_USER"`
	MemphisPassword        string `mapstructure:"MEMPHIS_PASSWORD"`
	MemphisAccountID       int    `mapstructure:"MEMPHIS_ACCOUNT_ID"`
	MemphisStationName     string `mapstructure:"MEMPHIS_STATION_NAME"`
	MemphisConsumer        string `mapstructure:"MEMPHIS_CONSUMER"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
