package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all api routes
type Route struct {
	URI                    string
	Method                 string
	function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}
// ToSetUp puts all routes inside the router
func ToSetUp(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.URI, route.function).Methods(route.Method)
	}
	return r
}
