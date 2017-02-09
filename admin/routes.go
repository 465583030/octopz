package admin

import (
	"github.com/Firemango/octopz/routing"
	"github.com/Firemango/octopz/state"
)

func Routes(r *routing.Router, context *state.Context) {
	r.AddRoute(routing.Route{"RegisterRoute", "POST", "/register_route", RegisterRoute(context)})
}
