package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv string
	App    struct {
		Host string
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Jwt struct {
		SecretKey                string
		TokenExpirationInSeconds int
	}
}

func LoadConfig(configFilePath string) (Config, error) {
	pflag.String("appenv", "dev", "application environment : dev | prod")

	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return Config{}, fmt.Errorf("err bind p flags : %w", err)
	}

	configName := fmt.Sprintf("config.%s", viper.GetString("appenv"))

	fmt.Println(configName)

	viper.SetConfigName(configName)
	viper.AddConfigPath(configFilePath)
	err = viper.ReadInConfig()
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
