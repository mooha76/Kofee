package config

import "github.com/spf13/viper"

type Config struct {
	ServiceUrl string `mapstructure:"SERVICE_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	Sslmode    string `mapstructure:"Sslmode"`
}

var envs = []string{
	"SERVICE_URL",
	"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD",
}

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
