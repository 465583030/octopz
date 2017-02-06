package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	RedisHost string
	RedisPort int
}

func main() {
	var config Config

	err := envconfig.Process("octopz", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

}
