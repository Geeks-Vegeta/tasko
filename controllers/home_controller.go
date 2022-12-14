package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello shreyas mohite",
	})

}

func AddName(c *gin.Context) {
	name := c.Param("name")
	x := fmt.Sprintf("My name is %s", name)
	c.JSON(http.StatusOK, gin.H{
		"message": x,
	})
}
