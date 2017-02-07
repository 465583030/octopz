package main

import (
	"fmt"
	"github.com/Firemango/octopz/routing"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type Config struct {
	RedisHost string `default:"locahost"`
	RedisPort int    `default:"6379"`
	Port      int    `default:"8080"`
	AdminPort int    `default:"9090"`
}

func main() {
	ar := routing.Router{}
	var c Config

	err := envconfig.Process("octopz", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	http.ListenAndServe(fmt.Sprintf(":%d", c.AdminPort), ar.ToMux())
}
