package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"

	"github.com/thhy/ginblog/db"
)

//User auth user
type User struct {
	ID         uint   `xorm:"not null pk 'id' autoincr" json:"id"`
	Name       string `xorm:"varchar(256) not null" json:"name"`
	Password   string `xorm:"varchar(256) not null" json:"password"`
	Level      uint   `xorm:"default(0)" json:"level"`
	CreateTime uint   `xorm:"created" json:"createtime"`
	UpdateTime uint   `xorm:"updated" json:"updatetime"`
	DeleteTime uint   `xorm:"deleted" json:"deletetime"`
}

//connect sqlite and create user database
func init() {
	// }
	exist, err := db.DB.IsTableExist(&User{})
	if err != nil {
		log.Panic("check user failed", err)
	}
	if !exist {
		err := db.DB.CreateTables(&User{})
		if err != nil {
			log.Panic("create user failed", err)
		}
	}
	log.Println("user init")
	log.Printf("%+v\n", db.DB.TableInfo(&User{}))
}

//TableName orm table name
func (article *User) TableName() string {
	return "user"
}

//Regist add user
func (user *User) Regist() error {
	has, err := db.DB.Exist(&User{
		Name: user.Name,
	})
	//check user exist
	if err != nil {
		log.Fatal("query user error", err)
	}
	if has {
		return errors.New("exist user")
	}
	mdpassword := encodePassword(user.Password)
	// user.Password = mdpassword
	tmp := User{Name: user.Name, Password: mdpassword}
	//insert into table user
	affectRow, err := db.DB.Insert(&tmp)
	if err != nil {
		log.Fatal("insert user error", err)
	}
	log.Println("affect row ", affectRow)
	return nil
}

//Auth check valid user
func (user *User) Auth() bool {
	mdpassword := encodePassword(user.Password)
	user.Password = mdpassword
	has, err := db.DB.Get(user)
	if err != nil {
		log.Fatal("query user error", err)
	}

	return has
}

func (user *User) Modify(newPassword string) error {
	//使用旧密码验证
	if !user.Auth() {
		return errors.New("auth failed")
	}

	if user.Password != "" {
		user.Password = encodePassword(newPassword)
	}
	//对密码进行更新
	db.DB.Where("id = ?", user.ID).Update("password = ?", user.Password)
	return nil
}

//encodePassword md5 password
func encodePassword(password string) string {
	salt := "kaikaikai"
	t := fmt.Sprintf("%s%s", password, salt)
	b := []byte(t)
	return fmt.Sprintf("%x", md5.Sum(b))
}

//
/*func (user *User) modifyPassword(NewPassword string) error {
	mdPassword := encodePassword(NewPassword)

}*/
