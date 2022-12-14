package main

import (
	"todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.HomeRoutes(router)
	router.Run("localhost:5000")

}
