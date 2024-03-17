package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// NewViper is a function to load config from config.json
// You can change the implementation, for example load from env file, consul, etcd, etc
func NewViper() *viper.Viper {
	config := viper.New()

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	config.SetConfigName("config." + env)

	config.SetConfigType("json")
	// config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}
