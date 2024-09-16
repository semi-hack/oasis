package main

import (
    "oasis/config"
    "oasis/routes"
)


func main() {
    config.ConnectDatabase()

    r := routes.SetupRouter()

	r.Run(":7070")
}
