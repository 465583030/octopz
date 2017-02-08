package routing

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestAddRouteToRouter(t *testing.T) {

	r := Router{}
	r.AddRoute(Route{
		Name:    "TestRoute",
		Method:  "GET",
		Pattern: "/test",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {

		}})

	if r.Routes[0].Name != "TestRoute" {
		t.Error("Route was not added.")
	}

	g := r.ToMux()

	g.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tr, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		if tr != "/test" {
			t.Error(fmt.Sprintf("Route failed: %s", tr))
		}
		return nil
	})

}
