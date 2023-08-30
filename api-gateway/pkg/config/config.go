package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	UserService   string `mapstructure:"USER_SRV"`
	MethodService string `mapstructure:"METHOD_SRV"`
}

var envs = []string{"PORT", "USER_SRV", "METHOD_SRV"}

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
