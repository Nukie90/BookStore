package infrastructure

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration from the file
func LoadConfig(file string) (*viper.Viper, error) {
	appConfig := viper.New()

	configPath := "./config"

	if _, err := os.Stat(configPath); err == nil {
		appConfig.AddConfigPath(configPath)

		appConfig.SetConfigName("common")

		if commonErr := appConfig.ReadInConfig(); commonErr != nil {
			return nil, commonErr
		}

		if len(file) > 0 {
			appConfig.SetConfigName(file)
			fmt.Println("Loading config file: ", file)
			if err := appConfig.MergeInConfig(); err != nil {
				return nil, err
			}
		}
	}
	appConfig.AutomaticEnv()

	return appConfig, nil

}


