package client

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
	"mjv.projects/motherBase/dao"
)

// UploadCsv Public methods
func UploadCsv(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	files := form.File["files"]

	go filesToLeads(files)

	return c.HTML(http.StatusOK, "<p>done</p>")
}

// APITest shows a simplre response
func APITest(c echo.Context) error {
	return c.HTML(http.StatusOK, "<p>Working</p>")
}

func filesToLeads(files []*multipart.FileHeader) {

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

		dao.LeadsBuldCreate(bs)
	}

	return
}
