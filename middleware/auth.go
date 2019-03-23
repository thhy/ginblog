package middleware

import (
	"log"
	"net/http"

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
			isExist, err := redis.Bool(db.RedisConn.Do("EXISTS", sessionID))
			log.Printf("%+v\n", isExist)
			if err != nil && !isExist {
				log.Panicln(err)
				c.JSON(http.StatusForbidden, gin.H{
					"message": "please login",
				})
				c.Set("is_logged_in", false)
				c.Abort()
			} else {
				c.Set("is_logged_in", true)
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
