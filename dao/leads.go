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
			tx.Save(irr.ToLead())
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

// CleanLeadCSV returns csvString without repetition
func CleanLeadCSV(files []*multipart.FileHeader) (string, int) {

	childs := []*model.ChildLead{}

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
		childs = append(childs, newChildLeads(bs)...)
	}
	str, err := gocsv.MarshalString(&childs)
	if err != nil {
		fmt.Println(err)
	}
	return str, len(childs)
}

func newChildLeads(bs []byte) []*model.ChildLead {
	childs := []*model.ChildLead{}
	irregulars := []*model.Irregular{}
	gocsv.UnmarshalString(model.CleanCsv(bs), &irregulars)

	for _, irr := range irregulars {
		rows := 0
		irr.CleanDatas()
		if irr.LinkedIn == "" {
			continue
		}

		irrLead := irr.ToLead()
		n := &model.Lead{LinkedIn: irrLead.LinkedIn}
		Conn.Model(&model.Lead{}).Where(&n).Count(&rows)
		if rows == 0 {
			childs = append(childs, irrLead.ToChildLead())
		}
	}
	return childs
}
