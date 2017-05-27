package dao

import (
	"fmt"

	"mjv.projects/motherBase/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn is a singleton DB connection
var irrArr = make(chan *model.Irregular, 200)
var Conn = newConnection()

func newConnection() *gorm.DB {
	conn, err := gorm.Open("sqlite3", "base.sqlite3")
	if err != nil {
		fmt.Println(err)
	}
	conn.AutoMigrate(&model.Lead{})
	conn.AutoMigrate(&model.Irregular{})
	conn.AutoMigrate(&model.Sheet{})

	return conn
}
