package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/model"
)

//PostArticle 提交文章
func PostArticle(c *gin.Context) {

}

//GetAllArticles 获取所有文章
func GetAllArticles() []model.Article {
	article := &model.Article{}

	articles := article.GetAllArticles(0, 30)
	return articles
}

//GetArticleByID 通过id获取文章
func GetArticleByID(c *gin.Context) {
	article := &model.Article{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		render(c, http.StatusNotFound, "article.html", gin.H{
			"errorMessage": "not found page",
		})
		return
	}
	res, err := article.Get(uint(id))
	if err != nil {
		render(c, http.StatusNotFound, "article.html", gin.H{
			"errorMessage": "not found page",
		})
		return
	}
	if err != nil {
		render(c, http.StatusBadRequest, "article.html", gin.H{
			"title": "invaild request id",
		})
		return
	}
	if res.ID == uint(id) {
		render(c, http.StatusOK, "article.html", gin.H{
			"payload": res,
			"title":   res.Title,
		})
	} else {
		render(c, http.StatusUnauthorized, "index.html", gin.H{
			"title": "index",
		})
	}
}

//NewArticle get post article page
func NewArticle(c *gin.Context) {
	if c.Request.Method == "GET" {
		render(c, http.StatusOK, "create_article.html", gin.H{
			"title": "create article",
		})
	} else if c.Request.Method == "POST" {
		title := c.PostForm("title")
		content := c.PostForm("content")
		article := &model.Article{Title: title, Content: content}
		article.Create()
		render(c, http.StatusOK, "submission-successful.html", gin.H{
			"title": "submit success",
		})
	}
}

//
func DeleteArticle(c *gin.Context) {

}
