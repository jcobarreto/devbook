package routes

import (
	"api/src/controllers"
	"net/http"
)

var PostRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Func:         controllers.CreatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Func:         controllers.GetPosts,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postId}",
		Method:       http.MethodGet,
		Func:         controllers.GetPost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postId}",
		Method:       http.MethodPut,
		Func:         controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postId}",
		Method:       http.MethodDelete,
		Func:         controllers.DeletePost,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/posts",
		Method:       http.MethodGet,
		Func:         controllers.GetPostsByUser,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postId}/like",
		Method:       http.MethodPost,
		Func:         controllers.LikePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postId}/unlike",
		Method:       http.MethodPost,
		Func:         controllers.UnlikePost,
		AuthRequired: true,
	},
}
