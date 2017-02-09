package admin

import (
	"github.com/Firemango/octopz/datastore"
	"github.com/Firemango/octopz/state"
	"io/ioutil"
	"net/http"
)

func RegisterRoute(ctx *state.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		ctx.Datastore.Broadcast(datastore.Message{
			Key:   "test",
			Value: string(body),
		})

		w.Write([]byte("Written?"))

	}
}
