package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/handler"
)

//InitializeRoutes 初始化路由
func InitializeRoutes(router *gin.Engine) {

	fmt.Println(os.Getwd())
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handler.Index)

	u := router.Group("/u")
	{
		u.GET("login", handler.Login)
		u.POST("login", handler.Login)
		u.GET("register", handler.Regist)
		u.POST("register", handler.Regist)
	}

	a := router.Group("/article")
	{
		a.GET("view/:id", handler.GetArticleByID)
		a.GET("create", handler.NewArticle)
		a.POST("create", handler.NewArticle)
	}
}
