package db

import (
	"log"

	"github.com/thhy/ginblog/conf"

	"github.com/garyburd/redigo/redis"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	//DB database op
	DB *xorm.Engine
	//RedisConn redis connection
	RedisConn redis.Conn
)

func init() {
	var err error
	DB, err = xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		log.Panic(err)
	}

	RedisConn, err = redis.Dial("tcp", conf.REDISNETWORK)
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
