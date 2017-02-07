package main

import (
	"fmt"
	"github.com/Firemango/octopz/admin"
	"github.com/Firemango/octopz/routing"
	"github.com/Firemango/octopz/state"
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
	c := Config{}

	admin.Routes(&ar, &state.Context{})

	err := envconfig.Process("octopz", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	http.ListenAndServe(fmt.Sprintf(":%d", c.AdminPort), ar.ToMux())
}
