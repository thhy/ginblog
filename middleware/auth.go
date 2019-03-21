package middleware

import (
	"log"
	"net/http"

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
			res, err := db.RedisConn.Do("KEYS", sessionID)
			log.Printf("%+v\n", res)
			if err != nil {
				log.Panicln(err)
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
			_, err := db.RedisConn.Do("KEYS", sessionID)
			if err == nil {
				c.Redirect(http.StatusFound, "http://localhost:8081")
				c.Abort()
			}
		}
		// c.Next()

	}
}
