package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	MongoDBUri string
}

var singleton *Config
var once sync.Once

func GetConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	once.Do(func() {
		singleton = &Config{
			MongoDBUri: os.Getenv("MONGODB_URI"),
		}
	})

	return singleton
}
