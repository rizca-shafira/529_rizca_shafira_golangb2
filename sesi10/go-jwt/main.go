package main

import (
	"golang_sesi10/go-jwt/database"
	"golang_sesi10/go-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8081")
}
