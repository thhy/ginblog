package handler

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/model"
)

//Login 用户登录
func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		user := &model.User{Name: username, Password: password}
		success := user.Auth()
		if !success {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"ErrorTitle":   "Login Failed",
				"ErrorMessage": "Invalid credentials provided"})
		} else {

			c.SetCookie("is_login_id", "true", 3600*24, "", "", false, true)
			c.SetCookie("token", generateSessionToken(), 3600, "", "", false, true)
			c.HTML(http.StatusOK, "login-successful.html", gin.H{
				"title": "loginSuccess",
			})
		}
	} else if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	}
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

//Regist regist user
func Regist(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user := &model.User{Name: username, Password: password}
		err := user.Regist()

		if err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"title":      "注册",
				"ErrorTitle": err,
			})
			return
		}

		c.SetCookie("is_login_id", "true", 3600*24, "", "", false, true)
		c.SetCookie("token", generateSessionToken(), 3600, "", "", false, true)
		c.HTML(http.StatusOK, "login-successful.html", gin.H{
			"title": "loginSuccess",
		})
	} else if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "注册",
		})
	}

}
