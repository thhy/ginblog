package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//Index 主页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"payload": GetAllArticles(),
	})
}
