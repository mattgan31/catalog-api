package main

import (
	"example.com/m/v2/database"
	"example.com/m/v2/router"
)

const PORT = ":3000"

func main() {
	database.StartDB()

	router.StartServer().Run(PORT)

}
