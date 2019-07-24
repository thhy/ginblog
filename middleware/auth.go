package middleware

import (
	"github.com/thhy/ginblog/logger"
	"log"
	"net/http"
	"reflect"

	"github.com/garyburd/redigo/redis"

	"github.com/thhy/ginblog/db"

	"github.com/gin-gonic/gin"
)

//AuthMiddleWare ensure user login
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusForbidden, gin.H{
				"message": "please login",
			})
			c.Abort()
			c.Set("is_logged_in", false)
		} else {
			var userInfo string
			var err error
			if values, err := db.RedisConn.Do("HMGET", sessionID, "info"); err == nil {
				arr, ok := values.([]interface{})
				if ok {
					userInfo = string(arr[0].([]byte))
					logger.Log(logger.DEBUG, reflect.TypeOf(userInfo))
					logger.Log(logger.DEBUG, userInfo)
				}

			} else {
				logger.Log(logger.DEBUG, "err:")
				logger.Log(logger.DEBUG, err)
			}

			if err != nil || userInfo == "" {
				logger.Log(logger.DEBUG, err)
				c.JSON(http.StatusForbidden, gin.H{
					"message": "please login",
				})
				c.Set("is_logged_in", false)
				c.Abort()
			} else {
				logger.Log(logger.DEBUG, "pass auth")
				c.Set("is_logged_in", true)
				c.Set("userInfo", userInfo)
				logger.Log(logger.DEBUG, c.MustGet("userInfo"))
				c.Next()
			}
		}

	}
}

//UnAuthMiddleWare ensure user not login
func UnAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		log.Printf("sessionID:%s\n", sessionID)
		if err == nil {
			isExist, err := redis.Bool(db.RedisConn.Do("EXISTS", sessionID))
			if err == nil && isExist {
				c.Set("is_logged_in", true)
				c.Redirect(http.StatusFound, "http://localhost:8081")
				c.Abort()
				return
			}
		}
		c.Set("is_logged_in", false)
		c.Next()

	}
}

//CheckLoginMiddleWare check user login
func CheckLoginMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		log.Printf("sessionID:%s\n", sessionID)
		if err == nil {
			isExist, err := redis.Bool(db.RedisConn.Do("EXISTS", sessionID))
			if err == nil && isExist {
				c.Set("is_logged_in", true)
				c.Next()
				return
			}
		}
		c.Set("is_logged_in", false)
		c.Next()

	}
}
