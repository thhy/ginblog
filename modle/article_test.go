package modle_test

import (
	"testing"

	"github.com/thhy/ginblog/handler"
)

func Test_Create(t *testing.T) {
	title := "first article"
	content := "ssssssssssssssssssss"
	if handler.Create(title, content) {
		t.Log("pass test")
	} else {
		t.Fatal("test failed")
	}
}

func Test_Find(t *testing.T) {
	keywords := "ss"
	article := handler.Find(keywords)
	if article != nil && len(article) >= 1 {
		t.Log("find")
	} else {
		t.Fatal("not found")
	}
}
