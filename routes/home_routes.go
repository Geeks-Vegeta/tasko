package routes

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func HomeRoutes(router *gin.Engine) {
	router.GET("/", controllers.HomeController)
	router.GET("/name/:name", controllers.AddName)

}
