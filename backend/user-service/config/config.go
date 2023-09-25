package config

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerHost     string `mapstructure:"SERVER_HOST"`
	ServerPort     string `mapstructure:"SERVER_PORT"`

	REDISHost     string `mapstructure:"REDIS_HOST"`
	REDISPassword string `mapstructure:"REDIS_PASSWORD"`
	REDISDB       int    `mapstructure:"REDIS_DB"`
	REDISPort     string `mapstructure:"REDIS_PORT"`
	REDISTimeout  int    `mapstructure:"REDIS_TIMEOUT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenPrivateKey string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenSecret    string        `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge    int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`

	MemphisHostName        string `mapstructure:"MEMPHIS_HOSTNAME"`
	MemphisApplicationUser string `mapstructure:"MEMPHIS_APPLICATION_USER"`
	MemphisPassword        string `mapstructure:"MEMPHIS_PASSWORD"`
	MemphisAccountID       int    `mapstructure:"MEMPHIS_ACCOUNT_ID"`
	MemphisStationName     string `mapstructure:"MEMPHIS_STATION_NAME"`
	MemphisStationNameTest string `mapstructure:"MEMPHIS_STATION_NAME_TEST"`
	MemphisProducer        string `mapstructure:"MEMPHIS_PRODUCER"`

	MemphisRetentionVal     int    `mapstructure:"MEMPHIS_RETENTION_VAL"`
	MemphisRetentionValTest int    `mapstructure:"MEMPHIS_RETENTION_VAL_TEST"`
	MemphisReplicas         int    `mapstructure:"MEMPHIS_REPLICAS"`
	MemphisIdempotency      int    `mapstructure:"MEMPHIS_IDEMPOTENCY"`
	MemphisPartitions       int    `mapstructure:"MEMPHIS_PARTITIONS"`
	MemphisSchemaName       string `mapstructure:"MEMPHIS_SCHEMA_NAME"`

	Pepper    string `mapstructure:"PEPPER"`
	EmailCode string `mapstructure:"EMAIL_CODE"`

	MigrationPath string `mapstructure:"MIGRATION_PATH"`
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
