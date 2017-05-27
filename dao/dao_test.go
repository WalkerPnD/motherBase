package dao

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestBulkCreateLeads(t *testing.T) {
	file := "/Users/Walker/godev/src/mjv.projects/motherBase/test/test.csv"
	//os.Remove("/Users/Walker/godev/src/mjv.projects/motherBase/dao/base.sqlite3")

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
	}

	BulkCreateLeads(bs)
}

func TestCleanLeadCSV(t *testing.T) {
	file := "/Users/Walker/godev/src/mjv.projects/motherBase/test/test.csv"
	//os.Remove("/Users/Walker/godev/src/mjv.projects/motherBase/dao/base.sqlite3")

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
	}

	CleanLeadCSV(bs)
}
