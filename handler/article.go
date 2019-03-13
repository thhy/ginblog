package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/modle"
)

//PostArticle 提交文章
func PostArticle(c *gin.Context) {

}

//GetAllArticles 获取所有文章
func GetAllArticles() []modle.Article {
	articles := modle.GetAllArticles()
	return articles
}

//GetArticleByID 通过id获取文章
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	article := modle.Get(id)
	log.Println("article.ID:? requesID:?", article.ID, id)
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "article.html", gin.H{
			"title": "invaild request id",
		})
		return
	}
	if article.ID == uint(intID) {
		c.HTML(http.StatusOK, "article.html", gin.H{
			"payload": article,
			"title":   article.Title,
		})

	} else {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{
			"title": "index",
		})
	}
}

//NewArticle get post article page
func NewArticle(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "create_article.html", gin.H{
			"title": "create article",
		})
	} else if c.Request.Method == "POST" {
		title := c.PostForm("title")
		content := c.PostForm("content")

		modle.Create(title, content)
		c.HTML(http.StatusOK, "submission-successful.html", gin.H{
			"title": "submit success",
		})
	}
}
