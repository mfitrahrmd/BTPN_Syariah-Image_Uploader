package config

import (
	"fmt"
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

func LoadConfig(configFilePath string) (Config, error) {
	configName := fmt.Sprintf("config.%s", os.Getenv("APPENV"))

	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configName)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("err read in config : %w", err)
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("err unmarshaling config : %w", err)
	}

	return config, err
}
