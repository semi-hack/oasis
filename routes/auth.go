package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var authRoutes = []endpoint{
	{
		Path: "/login",
		Method: http.MethodPost,
		Handlers: []gin.HandlerFunc{
			handlers.Login,
		},
	},
}