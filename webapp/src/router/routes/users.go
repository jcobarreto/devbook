package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:          "/create-user",
		Method:       http.MethodGet,
		Func:         controllers.LoadSignUpUserPage,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/get-users",
		Method:       http.MethodGet,
		Func:         controllers.LoadUsersPage,
		AuthRequired: true,
	},
}
