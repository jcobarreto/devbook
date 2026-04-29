package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRouter return a router with the routes configured
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigureRoutes(r)
}
