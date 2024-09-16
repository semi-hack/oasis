package routes

import (
	"net/http"
	"oasis/controllers"

	"github.com/gin-gonic/gin"
)

var categoryRoutes = []endpoint{
	{
		Path: "/",
		Method: http.MethodPost,
		Handlers: []gin.HandlerFunc{
			controllers.CreateCategory,
		},
	},
	{
		Path: "/",
		Method: http.MethodGet,
		Handlers: []gin.HandlerFunc{
			controllers.GetCategory,
		},
	},
}