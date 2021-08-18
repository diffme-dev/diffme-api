package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	MongoUri   string `mapstructure:"MONGODB_URI"`
	RedisUri   string `mapstructure:"REDIS_URI"`
	ElasticUri string `mapstructure:"ELASTIC_URI"`
}

var singleton *Config
var once sync.Once

func GetConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	once.Do(func() {
		var config *Config
		err := viper.Unmarshal(&config)

		if err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
		}

		log.Printf("Config %+v", config)
		singleton = config
	})

	return singleton
}
