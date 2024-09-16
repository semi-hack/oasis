package routes

import (
	"net/http"
	"oasis/controllers"

	"github.com/gin-gonic/gin"
)

var recommendationRoutes = []endpoint{
	{
		Path: "/",
		Method: http.MethodPost,
		Handlers: []gin.HandlerFunc{
			controllers.CreateRecommendation,
		},
	},
}

var unauthenticatedRecommendationRoutes = []endpoint{
	{
		Path: "/",
		Method: http.MethodGet,
		Handlers: []gin.HandlerFunc{
			controllers.GetRecommendations,
		},
	},
}