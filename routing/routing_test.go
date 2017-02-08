package routing

import (
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

}
