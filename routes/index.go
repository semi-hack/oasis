package routes

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

type endpoint struct {
	Path     string
	Method   string
	Handlers []gin.HandlerFunc
}

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*", "http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	r.HandleMethodNotAllowed = true
	r.MaxMultipartMemory = 20 << 20
	r.NoRoute(Handle404)
	r.NoMethod(Handle405)


    loadRoutes(r.Group("auth"), authRoutes...)
	loadRoutes(r.Group("recommendation"), unauthenticatedRecommendationRoutes...)
	loadRoutes(r.Group("category"), categoryRoutes...)

    r.Use(UserValidation)
	loadRoutes(r.Group("recommendation"), recommendationRoutes...)

    return r
}

func loadRoutes(r *gin.RouterGroup, endpoints ...endpoint) {
	for _, endpoint := range endpoints {
		r.Handle(endpoint.Method, endpoint.Path, endpoint.Handlers...)
	}
}