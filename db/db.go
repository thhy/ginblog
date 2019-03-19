package db

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

//DB operator database
var DB *xorm.Engine

func init() {
	var err error
	DB, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Panic(err)
	}

	// f, err := os.Create("test.log")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// logger := xorm.NewSimpleLogger(f)
	// logger.ShowSQL(true)
	// logger.SetLevel(core.LOG_DEBUG)
	// DB.SetLogger(logger)
	DB.ShowSQL(true)
	log.Println("db init")
}
