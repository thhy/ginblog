package model_test

import (
	"fmt"
	"testing"

	"github.com/thhy/ginblog/model"
)

func Test_Create(t *testing.T) {
	title := "first article"
	content := "ssssssssssssssssssss"
	article := &model.Article{Title: title, Content: content}
	err := article.Create()
	if err == nil {
		t.Log("pass test")
	} else {
		t.Fatal("test failed")
	}
}

func Test_Find(t *testing.T) {
	keywords := "kk"
	article := &model.Article{}
	articles, err := article.Find(keywords)
	fmt.Printf("%+v", articles)
	if err == nil {
		t.Log("find")
	} else {
		t.Fatal("not found", err)
	}
}

func Test_Regist(t *testing.T) {
	user := &model.User{Name: "czy", Password: "123456"}
	err := user.Regist()
	if err != nil {
		t.Fatal("test failed")
	} else {
		t.Log("pass")
	}
}

func Test_Auth(t *testing.T) {
	user := &model.User{Name: "czy", Password: "123456"}
	pass := user.Auth()
	if !pass {
		t.Fatal("test failed")
	} else {
		t.Log("pass")
	}
}
