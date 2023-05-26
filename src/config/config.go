package config

import "github.com/spf13/viper"

type Config struct {
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
	JWT_EXPIRE int    `mapstructure:"JWT_EXPIRE"`
	DB_CONNSTR string `mapstructure:"DB_CONNSTR"`
	RUN_TLS    bool   `mapstructure:"RUN_TLS"`
}

func NewConfig(path string) (config *Config, err error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	config = &Config{}
	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(config)
	return
}
