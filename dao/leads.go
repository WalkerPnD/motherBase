package dao

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"

	"github.com/gocarina/gocsv"
	"mjv.projects/motherBase/model"
)

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

// CSVsToLeads converts CSVs to slice of Lead struct
func CSVsToLeads(files []*multipart.FileHeader) {
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		defer src.Close()

		bs, err := ioutil.ReadAll(src)
		if err != nil {
			fmt.Println(err)
		}
		BulkCreateLeads(bs)
	}
}
