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
		tx.Model(&model.Lead{}).Where(&n).Count(&rows)
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

	childs = removeDuplicated(childs)
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
	if rows != 0 {
		return false
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

func removeDuplicated(datas []*model.Irregular) []*model.Irregular {
	newDatas := []*model.Irregular{}

	if len(datas) == 0 || len(datas) == 0 {
		return newDatas
	}

	datas[0].CleanDatas()
	newDatas = append(newDatas, datas[0])
	for _, data := range datas {
		exists := false
		data.CleanDatas()
		for _, nd := range newDatas {
			if nd.Exists(data) {
				exists = true
				break
			}
		}
		if exists {
			continue
		}
		newDatas = append(newDatas, data)
	}

	return newDatas
}
