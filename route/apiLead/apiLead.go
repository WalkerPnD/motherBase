package apiLead

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"mjv.projects/motherBase/dao"
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

	go dao.CSVsToLeads(files)

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

	c.Response().Header().Set("Content-Type", "application/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=planilhaFilho.csv")
	return c.String(http.StatusOK, csvFile)
}