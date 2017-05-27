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

	return conn
}

// BulkCreateLeads saves Irregulars and Leads
func BulkCreateLeads(bs []byte) {
	irregulars := []*model.Irregular{}
	gocsv.UnmarshalString(model.CleanCsv(bs), &irregulars)

	tx := Conn.Begin()
	for _, irr := range irregulars {
		irr.CleanDatas()
		if irr.LinkedIn != "" {
			lead := irr.ToLead(Conn)
			sheet := lead.Sheets[0]
			tx.FirstOrCreate(&sheet, model.Sheet{Name: lead.Sheets[0].Name})
			lead.Sheets[0] = sheet
			res := tx.FirstOrCreate(&lead, model.Lead{LinkedIn: lead.LinkedIn})
			if res.RowsAffected == int64(0) {
				lead.Sheets = append(lead.Sheets, sheet)
				tx.Save(lead)
			}
			continue
		}
		tx.Create(irr)
	}
	tx.Commit()
}
