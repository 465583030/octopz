package admin

import (
	"github.com/Firemango/octopz/state"
	"net/http"
)

func RegisterRoute(ctx *state.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := ctx.Redis.Set("key", "value", 0).Err()
		if err != nil {
			panic(err)
		}

		err = ctx.Redis.Publish("routes", "hello").Err()
		if err != nil {
			panic(err)
		}

	}
}
