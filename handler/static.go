package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func StaticFiles(c *gin.Context) {
	url := c.Request.URL.Path
	prefix := path.Base(url)
	var requestPath string
	if prefix == "css" {
		requestPath = "static/css"
	} else if prefix == "js" {
		requestPath = "static/js"
	} else if prefix == "html" {
		requestPath = "static/html"
	}
	c.String(http.StatusOK, requestPath)
}
