package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, mainPageRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {

		if route.AuthRequired {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Func),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
