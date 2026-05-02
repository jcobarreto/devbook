package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Func:         controllers.CreatePost,
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
	{
		URI:          "/posts/{postId}/update",
		Method:       http.MethodGet,
		Func:         controllers.LoadEditPostPage,
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
}
