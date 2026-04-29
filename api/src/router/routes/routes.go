package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all API routes
type Route struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := UserRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, PostRoutes...)

	for _, route := range routes {

		if route.AuthRequired {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
