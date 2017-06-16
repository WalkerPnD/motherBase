package apiLead

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"mjv.projects/motherBase/dao"
	"mjv.projects/motherBase/model"
)

// BulkCreate recieves csv file and create Leads on DataBase
func BulkCreate(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.HTML(http.StatusOK, err.Error())
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.HTML(http.StatusOK, "<p>Não há arquivos</p>")
	}

	dao.CSVsToLeads(files)

	return c.HTML(http.StatusOK, "<p>done</p>")
}

// APITest shows a simplre response
func APITest(c echo.Context) error {
	return c.HTML(http.StatusOK, "<p>Working</p>")
}

// CleanCSV removes the existing leads which already in motherBase
func CleanCSV(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.HTML(http.StatusOK, "<p>Não há arquivos</p>")
	}

	csvFile, newContacts := dao.CleanLeadCSV(files)
	if newContacts == 0 {
		return c.HTML(http.StatusOK, "<p>Não há novos contatos</p>")
	}

	fileName := files[0].Filename + "-planilhaFilho.csv"
	c.Response().Header().Set("Content-Type", "application/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)
	return c.String(http.StatusOK, csvFile)
}

// JoinInDatas
func JoinInDatas(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.HTML(http.StatusOK, "<p>Não há arquivos</p>")
	}

	csvFile, newContacts := model.JoinContacts(files)
	if newContacts == 0 {
		return c.HTML(http.StatusOK, "<p>Não há novos contatos</p>")
	}

	fileName := files[0].Filename + "-clean.csv"
	c.Response().Header().Set("Content-Type", "application/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)
	return c.String(http.StatusOK, csvFile)
}

func UpdateDB(c echo.Context) error {
	dao.UpDateDB()
	return c.HTML(http.StatusOK, "<p>Done</p>")
}
