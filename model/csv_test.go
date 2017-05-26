package model

import "testing"

func TestCleanCsvHeader(t *testing.T) {
	actual := CleanCsvHeader("Company Name,first Name,Last name,Job Title  ,Industry, City ,Email , Linkedin, Nome da Planilha")
	expected := "Company Name,First Name,Last Name,Job Title,Industry,City,Email,Linkedin,Nome Da Planilha"
	if actual != expected {
		t.Errorf("\ngot:  '%s' \nwant: '%s'", actual, expected)
	}
}
