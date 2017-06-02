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
	rows := 0
	irregulars := []*model.Irregular{}
	gocsv.UnmarshalString(model.CleanCsv(bs), &irregulars)

	tx := Conn.Begin()
	for _, irr := range irregulars {
		irr.CleanDatas()
		n := &model.Lead{
			CompanyName: irr.CompanyName,
			FirstName:   irr.FirstName,
			LastName:    irr.LastName,
			JobTitle:    irr.JobTitle,
		}
		Conn.Model(&model.Lead{}).Where(&n).Count(&rows)
		if rows == 0 {
			tx.Save(irr.ToLead())
			continue
		}
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

	childs := []*model.Irregular{}

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

func newChildLeads(bs []byte) []*model.Irregular {
	childs := []*model.Irregular{}
	irregulars := []*model.Irregular{}
	gocsv.UnmarshalString(model.CleanCsv(bs), &irregulars)
	fmt.Println(irregulars[0].Email)

	for _, irr := range irregulars {

		irr.CleanDatas()
		if irr.LinkedIn == "" {
			continue
		}

		if isNew(irr) {
			childs = append(childs, irr)
		}
	}
	return childs
}

func isNew(target *model.Irregular) bool {
	rows := 0
	n := &model.Lead{}
	n.Email = target.Email
	Conn.Model(&model.Lead{}).Where(&n).Count(&rows)
	if rows == 0 {
		return true
	}

	n = &model.Lead{
		CompanyName: target.CompanyName,
		FirstName:   target.FirstName,
		LastName:    target.LastName,
		JobTitle:    target.JobTitle,
	}
	Conn.Model(&model.Lead{}).Where(&n).Count(&rows)
	if rows == 0 {
		return true
	}
	return false
}
