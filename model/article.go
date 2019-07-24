package model

import (
	"errors"
	"log"

	"github.com/thhy/ginblog/db"
)

//Article article infos
type Article struct {
	ID         uint   `xorm:"not null pk 'id' autoincr"`
	Title      string `json:"title" xorm:"VARCHAR(255)"`
	Content    string `json:"content" xorm:"VARCHAR(10240)"`
	CommentURL string `json:"comment" xorm:"VARCHAR(1024)"`
	CreateTime uint   `xorm:"created"`
	UpdateTime uint   `xorm:"updated"`
	DeleteTime uint   `xorm:"deleted"`
	AutherID   uint   `xorm:"autherId"`
}

func init() {
	exist, err := db.DB.IsTableExist(&Article{})
	if err != nil {
		log.Panic(err)
		log.Panic("check Article failed", err)
	}
	if !exist {
		err := db.DB.CreateTables(&Article{})
		if err != nil {
			log.Panic(err)
			log.Panic("create Article failed", err)
		}
	}
	log.Println("article init")
	log.Printf("%+v\n", db.DB.TableInfo(&Article{}))
}

//TableName orm table name
func (article *Article) TableName() string {
	return "article"
}

//Create create article
func (article *Article) Create() error {
	_, err := db.DB.Insert(article)
	if err != nil {
		log.Fatal("insert article error", err)
		return err
	}
	return nil
}

//Get article according article id
func (article *Article) Get(id uint) (*Article, error) {
	var articles []Article
	err := db.DB.Where("id = ?", id).Find(&articles)
	if len(articles) == 0 {
		return nil, errors.New("not found page")
	}
	return &articles[0], err
}

//Delete artcile according articleId
func (article *Article) Delete() error {
	var art Article
	affect, err := db.DB.Where("id = ?", article.ID).And("autherID = ?", article.AutherID).Delete(&art)
	if err != nil {
		return err
	}
	if affect == 0 {
		return errors.New("this artcile is not exist")
	}
	return nil
}

//Modify article
func (article *Article) Modify(id uint) error {
	_, err := db.DB.Where("id = ?", id).Cols("title", "content").Update(&article)
	return err
}

//Find articles by keywords
func (article *Article) Find(keywords string) ([]Article, error) {
	var articles []Article
	err := db.DB.Where("title like ? or content like ?", "%"+keywords+"%", "%"+keywords+"%").Find(&articles)
	return articles, err
}

//GetAllArticles get all articles
func (article *Article) GetAllArticles(start int, count int) []Article {
	var articles []Article
	err := db.DB.Desc("id").Find(&articles)
	if err != nil {
		log.Fatal(err)
	}
	return articles
}
