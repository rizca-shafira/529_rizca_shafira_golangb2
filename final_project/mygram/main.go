package main

import (
	"final_project/mygram/database"
	"final_project/mygram/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8081")
}
