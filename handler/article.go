package handler

import (
	"encoding/json"
	"errors"
	"github.com/thhy/ginblog/logger"
	"net/http"
	"reflect"
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
		autherId, _ := getCurSessionId(c)
		article := &model.Article{Title: title, Content: content, AutherID: autherId}
		article.Create()
		render(c, http.StatusOK, "submission-successful.html", gin.H{
			"title": "submit success",
		})
	}
}

//DeleteArticle delete article by id
func DeleteArticle(c *gin.Context) {
	if c.Request.Method == "GET" {
		articleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			render(c, http.StatusBadRequest, "index.html", gin.H{
				"title": "delete failed",
			})
			c.Abort()
		}
		uArticleId := uint(articleId)
		userInfo := c.MustGet("userInfo").(string)
		var user model.User

		err = json.Unmarshal([]byte(userInfo), &user)
		if err != nil {
			logger.Log(logger.ERROR, err)
			logger.Log(logger.INFO, reflect.TypeOf(userInfo))
			logger.Log(logger.INFO, userInfo)

			render(c, http.StatusBadRequest, "index.html", gin.H{
				"title": "delete failed",
			})
			c.Abort()
		}
		article := &model.Article{ID: uArticleId, AutherID: user.ID}
		logger.Log(logger.DEBUG, user)
		logger.Log(logger.DEBUG, *article)

		if err := article.Delete(); err != nil {
			render(c, http.StatusBadRequest, "index.html", gin.H{
				"title":   "delete failed",
				"content": err,
			})
		}

	}
}

func getCurSessionId(c *gin.Context) (uint, error) {
	userInfo := c.MustGet("userInfo").(string)
	var user model.User

	err := json.Unmarshal([]byte(userInfo), &user)
	if err != nil {
		return 0, errors.New("get user id failed")
	}
	return user.ID, nil
}
