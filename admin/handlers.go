package admin

import (
	"fmt"
	"github.com/Firemango/octopz/state"
	"net/http"
)

func RegisterRoute(ctx *state.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello")

	}
}
