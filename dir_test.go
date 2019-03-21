package main_test

import (
	"log"
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	path := "D:/GOPATH/src/github.com/thhy/ginblog/static/templates"
	dir, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer dir.Close()
	names, _ := dir.Readdirnames(-1)

	for _, n := range names {
		log.Println("file:", n)
	}
	t.Fail()
}
