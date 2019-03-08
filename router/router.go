package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// var router *gin.Engine

func InitializeRoutes(router *gin.Engine) {

	fmt.Println(os.Getwd())
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
		// c.JSON(http.StatusOK, gin.H{
		// 	"Hello": "World",
		// })
	})
}
