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
}
