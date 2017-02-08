package main

import (
	"fmt"
	"github.com/Firemango/octopz/admin"
	"github.com/Firemango/octopz/routing"
	"github.com/Firemango/octopz/state"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
)

type Config struct {
	RedisHost string `default:"127.0.0.1"`
	RedisPort int    `default:"6379"`
	Port      int    `default:"8080"`
	AdminPort int    `default:"9090"`
}

func main() {
	ar := routing.Router{}
	c := Config{}

	err := envconfig.Process("octopz", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisHost, c.RedisPort),
		Password: "",
		DB:       0,
	})

	ctx := state.Context{Redis: r}

	admin.Routes(&ar, &ctx)

	http.ListenAndServe(fmt.Sprintf(":%d", c.AdminPort), ar.ToMux())
}
