package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var mainPageRoute = Route{
	URI:          "/home",
	Method:       http.MethodGet,
	Func:         controllers.LoadMainPage,
	AuthRequired: true,
}
