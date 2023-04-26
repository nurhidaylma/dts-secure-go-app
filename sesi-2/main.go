package main

import (
	"github.com/nurhidaylma/dts-secure-go-app/database"
	"github.com/nurhidaylma/dts-secure-go-app/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
