package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Func:         controllers.GetUsers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Func:         controllers.GetUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Func:         controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Func:         controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/follow",
		Method:       http.MethodPost,
		Func:         controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/unfollow",
		Method:       http.MethodPost,
		Func:         controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/followers",
		Method:       http.MethodGet,
		Func:         controllers.GetFollowers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/following",
		Method:       http.MethodGet,
		Func:         controllers.GetFollowing,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/update-password",
		Method:       http.MethodPost,
		Func:         controllers.UpdatePassword,
		AuthRequired: true,
	},
}
