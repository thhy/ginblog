package modle

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

//User auth user
type User struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"size:256"`
	Password string `gorm:"size:256"`
	Level    uint
}

var userDB *gorm.DB

//connect sqlite and create user database
func init() {
	var err error
	userDB, err = gorm.Open("sqlite3", "article.db")
	if err != nil {
		panic(err)
	}
	userDB.LogMode(true)
	// }
	log.Println(userDB)
	if !userDB.HasTable(&User{}) {
		userDB.CreateTable(&User{})
	}
}

//AuthUser check user is valid
func AuthUser(username string, password string) User {
	var user User
	mdpassword := encodePassword(password)
	userDB.Where("name = ? and password = ?", username, mdpassword).First(&user)
	return user
}

//encodePassword md5 password
func encodePassword(password string) string {
	salt := "kaikaikai"
	t := fmt.Sprintf("%s%s", password, salt)
	b := []byte(t)
	return fmt.Sprintf("%x", md5.Sum(b))
}

//Regist regist user
func Regist(username string, password string) (User, error) {
	mdpassword := encodePassword(password)
	var user User
	// user := User{Name:username, Password:mdpassword}
	if userDB.Where("name = ? and password = ?", username, mdpassword).First(&user).RecordNotFound() {
		if !userDB.NewRecord(&user) {
			userDB.Create(user)
		}
		return user, nil
	}
	return User{}, fmt.Errorf("exist user")

}
