// package api

// import (
// 	"database/sql"
// 	"net/http"

// 	"oasis/routes"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	db "oasis/db/sqlc"
// )

// type APIServer struct {
// 	addr string
// 	// store *db.Store
// }

// // NEWAPIServer returns a pointer to an APIServer which is a structure
// // that holds the state required for running the API server.
// func NEWAPIServer(addr string) *APIServer {
// 	return &APIServer{addr: addr}
// }

// // Run starts the API server
// func (s *APIServer) Run() error {

// 	router := gin.Default()
// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"*"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 	}))
	
// 	return http.ListenAndServe(s.addr, router)
// }
