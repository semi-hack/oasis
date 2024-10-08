package routes

import (
	"net/http"
	"oasis/controllers"

	"github.com/gin-gonic/gin"
)

var authRoutes = []endpoint{
	{
		Path: "/login",
		Method: http.MethodPost,
		Handlers: []gin.HandlerFunc{
			controllers.Login,
		},
	},
	{
		Path: "/signup",
		Method: http.MethodPost,
		Handlers: []gin.HandlerFunc{
			controllers.Signup,
		},
	},
}