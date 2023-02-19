package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	AppHost string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"APP_PORT"`

	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`

	JwtSecretKey             string        `mapstructure:"JWT_SECRET_KEY"`
	JwtTokenExpirationLength time.Duration `mapstructure:"JWT_TOKEN_EXPIRATION_LENGTH"`
}

func New(configFilePath string, logger *logrus.Logger) *Config {
	logger.Infoln("[~] loading config..")

	if os.Getenv("APPENV") == "" {
		logger.Infoln("[~] application environment does not set, dev mode will be used..")

		err := os.Setenv("APPENV", "dev")
		if err != nil {
			panic(err)
		}
	}

	logger.Infoln(fmt.Sprintf("running in %s mode", os.Getenv("APPENV")))

	configName := fmt.Sprintf("config.%s", os.Getenv("APPENV"))

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configName)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
