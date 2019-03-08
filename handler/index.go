package handler

import (
	"fmt"
	"ginblog/conf"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

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
