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
		} else {
			isExist, err := redis.Bool(db.RedisConn.Do("EXISTS", sessionID))
			log.Printf("%+v\n", isExist)
			if err != nil || !isExist {
				log.Println(err)
				if !isExist {
					c.SetCookie("session_id", "", -1, "", "", false, true)
				}
				c.JSON(http.StatusForbidden, gin.H{
					"message": "please login",
				})
				c.Abort()
			} else {
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
			if !isExist {
				c.SetCookie("session_id", "", -1, "", "", false, true)
			}
			if err == nil && isExist {
				c.Redirect(http.StatusFound, "http://localhost:8081")
				c.Abort()
			}
		}
		// c.Next()

	}
}
