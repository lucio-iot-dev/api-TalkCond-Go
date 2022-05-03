package router

import (
	"api-TalkCond/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return a router with the routes
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.ToSetUp(r)
}