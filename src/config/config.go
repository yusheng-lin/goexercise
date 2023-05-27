package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	JWT_SECRET        string `mapstructure:"JWT_SECRET"`
	JWT_EXPIRE        int    `mapstructure:"JWT_EXPIRE"`
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	RUN_TLS           bool   `mapstructure:"RUN_TLS"`
}

func NewConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	files, err := os.ReadDir(path)

	if err != nil {
		return
	}
	config = &Config{}
	for _, file := range files {
		viper.SetConfigFile(fmt.Sprintf("%s/%s", path, file.Name()))
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
		err = viper.Unmarshal(config)
	}
	return
}
