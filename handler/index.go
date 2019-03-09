package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/thhy/ginblog/conf"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//InitializeRoutes init route
func InitializeRoutes() {
	//change path to work dir
	os.Chdir(conf.SRCPATH)

	fmt.Println(os.Getwd())
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"title": "Home Page",
		// })
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})
}

//Run start listen
func Run() {
	router.Run(conf.BINDADDRESS)
}
