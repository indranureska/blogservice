package utils

import "github.com/spf13/viper"

type Config struct {
	BlogDbURI string `mapstructure:"BLOG_DB_URI"`
}

// LoadConfig reads configuration from file.
// Takes a path as input, and returns a config object or an error.
func LoadConfig(path string) (config Config, err error) {
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
