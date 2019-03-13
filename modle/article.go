package modle

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Article article infos
type Article struct {
	gorm.Model
	Title      string `json:"title" gorm:"size:255"`
	Content    string `json:"content"`
	CommentURL string `json:"comment" gorm:"size:1024"`
}

var articleDB *gorm.DB

func init() {
	var err error
	articleDB, err = gorm.Open("sqlite3", "article.db")
	if err != nil {
		panic(err)
	}
	articleDB.LogMode(true)
	// }
	log.Println(articleDB)
	if !articleDB.HasTable(&Article{}) {
		articleDB.CreateTable(&Article{})
	}
}

//BeforeCreate check before create data
/*func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	if article.Content == "" && article.Title == "" {
		log.Fatal("article content or title is null")
	}
	return nil
}*/

//Create create article and insert into database
func Create(title string, content string) bool {
	if !articleDB.HasTable(&Article{}) {
		articleDB.CreateTable(&Article{})
	}
	article := &Article{Title: title, Content: content}
	if articleDB.NewRecord(*article) {
		log.Println("create")
		articleDB.Create(article)
	}
	return true
}

//Get article according id
func Get(id string) Article {
	if !articleDB.HasTable(&Article{}) {
		return Article{}
	}
	var article Article
	articleDB.Where("ID = ?", id).First(&article)
	return article
}

//Find get artiles according keywords
func Find(keywords string) []Article {
	if !articleDB.HasTable(&Article{}) {
		return nil
	}
	var articles []Article
	queryWords := "%" + keywords + "%"
	articleDB.Where("content like ? or title like ?", queryWords, queryWords).Find(&articles)
	return articles
}

//GetAllArticles get all articles
func GetAllArticles() []Article {
	var articles []Article
	articleDB.Find(&articles)
	return articles
}
