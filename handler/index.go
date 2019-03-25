package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 主页
func Index(c *gin.Context) {
	render(c, http.StatusOK, "index.html", gin.H{
		"payload": GetAllArticles(),
	})
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, statusCode int, templateName string, data gin.H) {
	loggedInInterface, _ := c.Get("is_logged_in")
	log.Println(loggedInInterface)
	data["is_logged_in"], _ = loggedInInterface.(bool)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(statusCode, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(statusCode, data["payload"])
	default:
		// Respond with HTML
		c.HTML(statusCode, templateName, data)
	}
}
