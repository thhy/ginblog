package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/modle"
)

func Login(c *gin.Context) {
	username := c.PostForm["username"]
	password := c.PostForm["password"]

	user := modle.AuthUser(userame, password)
	if user.ID == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	} else {

		c.SetCookie("is_login_id", "true")
		c.SetCookie("token", generateSessionToken(), 3600, "", "", false, true)
		c.HTML(http.StatusOK, "login-successful.html", gin.H{
			"title":"loginSuccess"
		})
	}
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func Re