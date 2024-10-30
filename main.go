package main

import (
	"gin/database"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()
	routes.RegisterRoutes(router)
	routes.BarangRoutes(router)
	routes.JenisBarangRoutes(router)
	router.Run(":8080")
}
