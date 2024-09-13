package main

import (
	"log"
	"oasis/cmd/api"
)

func main() {
	server := api.NEWAPIServer(":4050", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}