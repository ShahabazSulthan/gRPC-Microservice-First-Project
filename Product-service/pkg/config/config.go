package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBHost        string `mapstructure:"DBHost"`
	DBPort        string `mapstructure:"DBPort"`
	DBUser        string `mapstructure:"DBUser"`
	DBPassword    string `mapstructure:"DBPassword"`
	DBName        string `mapstructure:"DBName"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
