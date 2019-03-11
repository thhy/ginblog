package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/modle"
)

func PostArticle(c *gin.Context) {
	title := c.PostForm["title"]
	content := c.PostForm["content"]

	modle.Create(title, content)
}

func GetAllArticles() []modle.Article {
	articles := modle.GetAllArticles()
	return articles
}

func GetArticleById(id uint) modle.Article {
	article := modle.Get(id)
	return article
}
