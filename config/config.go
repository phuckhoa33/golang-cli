package config

import "github.com/spf13/viper"

// Config struct
// This is the struct for the config
// This struct will store the configuration for the app
type Config struct {
	App AppConfig
}

// NewConfig function
// This function is used to create a new config
func NewConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	return &Config{
		App: LoadAppConfig(),
	}
}
