package model

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"

	"github.com/gocarina/gocsv"
)

// LinkedIn is a potential sales contact
type LinkedIn struct {
	CompanyName string `csv:"Company Name"`
	FullName    string `csv:"Full Name"`
	JobTitle    string `csv:"Job Title"`
	City        string `csv:"City"`
	LinkedIn    string `csv:"Linkedin"`
}

func (l *LinkedIn) CleanDatas() {
	l.CompanyName = cleanData(l.CompanyName)
	l.FullName = cleanData(l.FullName)
	l.JobTitle = cleanData(l.JobTitle)
	l.City = strings.TrimSpace(l.City)
	l.LinkedIn = strings.TrimSpace(l.LinkedIn)
}

func (l *LinkedIn) exists(cmp *LinkedIn) bool {
	if l.FullName == cmp.FullName && l.CompanyName == cmp.CompanyName && l.JobTitle == cmp.JobTitle && l.City == cmp.City {
		return true
	}
	return false
}

func JoinContacts(files []*multipart.FileHeader) (string, int) {

	lin := []*LinkedIn{}
	datas := []*LinkedIn{}

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
		gocsv.UnmarshalString(CleanCsv(bs), &datas)
		lin = append(lin, datas...)
	}

	lin = removeDuplicated(lin)
	str, err := gocsv.MarshalString(&lin)
	if err != nil {
		fmt.Println(err)
	}
	return str, len(lin)
}

func removeDuplicated(datas []*LinkedIn) []*LinkedIn {
	newDatas := []*LinkedIn{}

	if len(datas) == 0 || len(datas) == 0 {
		return newDatas
	}

	datas[0].CleanDatas()
	newDatas = append(newDatas, datas[0])
	for _, data := range datas {
		exists := false
		data.CleanDatas()
		for _, nd := range newDatas {
			if nd.exists(data) {
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
