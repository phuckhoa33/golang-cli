package config

import "github.com/spf13/viper"

// AppConfig Config struct
// This is the struct for the config
// This struct will store app configuration
type AppConfig struct {
	AppName             string
	AppVersion          string
	AppShortDescription string
	AppLongDescription  string
}

// LoadAppConfig function
// This function is used to create a new app config
func LoadAppConfig() AppConfig {
	return AppConfig{
		AppName:             viper.GetString("CLI_NAME"),
		AppVersion:          viper.GetString("CLI_VERSION"),
		AppShortDescription: viper.GetString("CLI_SHORT_DESCRIPTION"),
		AppLongDescription:  viper.GetString("CLI_LONG_DESCRIPTION"),
	}
}
