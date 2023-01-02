package main

import (
	"main/database"
	routes "main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dbErr := database.Init()
	if dbErr != nil {
		panic(dbErr)
	}

	router := gin.New()

	router.Use(gin.Logger())
	routes.StatusRoutes(router)

	router.Run("0.0.0.0:8080")
}
