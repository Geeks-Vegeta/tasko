package main

import (
	"todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routes.HomeRoutes(router)
	router.Run(":5000")

}
