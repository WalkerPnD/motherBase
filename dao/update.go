package dao

import (
	"fmt"

	"mjv.projects/motherBase/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// UpDataDB creates a new database
func UpDateDB() {
	newConn, err := gorm.Open("sqlite3", "new_base.sqlite3")
	if err != nil {
		fmt.Println(err)
	}
	newConn.AutoMigrate(&model.Lead{})

	leads := []*model.Lead{}
	Conn.Find(&leads)
	fmt.Println(len(leads))
	tx := newConn.Begin()
	for _, l := range leads {
		tx.Save(l)
	}
	tx.Commit()
	tx.Close()
}
