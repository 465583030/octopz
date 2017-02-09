package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	RedisHost     string `default:"127.0.0.1"`
	RedisPort     int    `default:"6379"`
	RedisPassword string `default:""`
	Port          int    `default:"8080"`
	AdminPort     int    `default:"9090"`
}

func ConfigFromEnv() Config {
	c := Config{}

	err := envconfig.Process("octopz", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
