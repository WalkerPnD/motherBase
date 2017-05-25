package dao

import (
	"fmt"

	"mjv.projects/motherBase/model"

	"github.com/gocarina/gocsv"
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
	go initLeadsCreatePool(conn)

	return conn
}

func LeadsBuldCreate(bs []byte) {
	irregulars := []*model.Irregular{}
	gocsv.UnmarshalString(model.CleanCsv(bs), &irregulars)
	for _, irr := range irregulars {
		irrArr <- irr
	}
}

func initLeadsCreatePool(conn *gorm.DB) {
	for {
		irregular, _ := <-irrArr
		if irregular.LinkedIn != "" {
			lead := model.IrregularToLead(irregular)
			err := conn.Create(lead)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		conn.Create(irregular)
	}
}
