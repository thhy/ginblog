package handler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/db"
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
			render(c, http.StatusUnauthorized, "login.html", gin.H{
				"ErrorTitle":   "Login Failed",
				"ErrorMessage": "Invalid credentials provided"})
		} else {
			sessionID := generateSessionToken()
			jsonUser, err := json.Marshal(*user)
			if err != nil {
			}
			storeSessionID(sessionID, string(jsonUser), 86400)
			//set cookie ttl
			c.SetCookie("session_id", sessionID, 3600*24, "", "", false, true)
			c.Set("is_logged_in", true)
			render(c, http.StatusOK, "login-successful.html", gin.H{
				"title": "loginSuccess",
			})
		}
	} else if c.Request.Method == "GET" {
		render(c, http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	}
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
}

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

//Regist regist user
func Regist(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")

		user := &model.User{Name: username, Password: password}
		err := user.Regist()

		if err != nil {
			render(c, http.StatusOK, "register.html", gin.H{
				"title":      "注册",
				"ErrorTitle": err,
			})
			return
		}
		sessionID := string(rand.Int31())
		jsonUser, err := json.Marshal(*user)
		if err != nil {
		}
		storeSessionID(sessionID, string(jsonUser), 86400)
		c.Set("is_logged_in", true)
		c.SetCookie("session_id", string(sessionID), 3600*24, "", "", false, true)
		render(c, http.StatusOK, "login-successful.html", gin.H{
			"title": "loginSuccess",
		})
	} else if c.Request.Method == "GET" {
		render(c, http.StatusOK, "register.html", gin.H{
			"title": "注册",
		})
	}

}

//Logout clear redis info and cookie
func Logout(c *gin.Context) {
	sessionID, _ := c.Cookie("session_id")
	db.RedisConn.Do("DEL", sessionID)
	c.SetCookie("session_id", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	render(c, http.StatusOK, "logout.html", gin.H{
		"jump": "/",
	})
}

//storeSessionID store sessionid into redis and set expire
func storeSessionID(sessionID, userInfo string, timeout int32) {
	_, err := db.RedisConn.Do("HMSET", sessionID, "info", userInfo)
	if err != nil {
	}
	//set session ttl
	_, err = db.RedisConn.Do("EXPIRE", sessionID, timeout)
	if err != nil {
	}
}
