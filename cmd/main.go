// package main

// import (
// 	"log"
// 	"oasis/cmd/api"
// 	"oasis/db"
	
// )

// func main() {
// 	db.ConnectDatabase()

// 	server := api.NEWAPIServer(":4050", nil)
// 	if err := server.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }