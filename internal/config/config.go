package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Constants consist of environment variables
type Constants struct {
	KAFKA_CONN string
}

//Config consist of Constants
type Config struct {
	Constants
}

// GetEnv : Function for getting the environment variables
func GetEnv() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	if err != nil {
		return nil, err
	}
	config.Constants = constants
	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("kafka-conn-config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")                 // Search the root directory for the configuration file
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		return Constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
