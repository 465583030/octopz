package main

import (
	"fmt"
	"github.com/Firemango/octopz/admin"
	"github.com/Firemango/octopz/config"
	"github.com/Firemango/octopz/datastore"
	"github.com/Firemango/octopz/routing"
	"github.com/Firemango/octopz/state"
	"net/http"
)

func main() {
	ar := routing.Router{}

	c := config.ConfigFromEnv()
	ds := datastore.NewDatastore(c)
	ctx := state.Context{Datastore: &ds}

	admin.Routes(&ar, &ctx)

	http.ListenAndServe(fmt.Sprintf(":%d", c.AdminPort), ar.ToMux())
}
