package config

import "github.com/spf13/viper"

type Config struct {
	AuthServiceUrl     string `mapstructure:"AUTH_SERVICE_URL"`
	UserServiceUrl     string `mapstructure:"USER_SERVICE_URL"`
	Port               string `mapstructure:"PORT"`
	MerchantServiceUrl string `mapstructure:"MERCHANT_SERVICE_URL"`
}

var envs = []string{"AUTH_SERVICE_URL", "USER_SERVICE_URL", "PORT"}

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			return
		}
	}
	err = viper.Unmarshal(&config)

	return
}
