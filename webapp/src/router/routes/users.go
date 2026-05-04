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
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Func:         controllers.LoadUserProfilePage,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/unfollow",
		Method:       http.MethodPost,
		Func:         controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/follow",
		Method:       http.MethodPost,
		Func:         controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Func:         controllers.LoadUserLoggedProfile,
		AuthRequired: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodGet,
		Func:         controllers.LoadEditUserPage,
		AuthRequired: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodPut,
		Func:         controllers.EditUser,
		AuthRequired: true,
	},
	{
		URI:          "/update-password",
		Method:       http.MethodGet,
		Func:         controllers.LoadUpdatePasswordPage,
		AuthRequired: true,
	},
	{
		URI:          "/update-password",
		Method:       http.MethodPost,
		Func:         controllers.UpdatePassword,
		AuthRequired: true,
	},
}
