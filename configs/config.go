package configs

import "github.com/spf13/viper"

type Config struct {
	Host           string `mapstructure:"REDIS_HOST"`
	Port           string `mapstructure:"REDIS_PORT"`
	Password       string `mapstructure:"REDIS_PASSWORD"`
	DB             int    `mapstructure:"REDIS_DB"`
	LimitByToken   bool   `mapstructure:"LIMIT_BY_TOKEN"`
	RateLimit      int    `mapstructure:"RATE_LIMIT"`
	ExpirationTime int    `mapstructure:"EXPIRATION_TIME"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
