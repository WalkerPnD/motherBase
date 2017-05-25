package model

import (
	"fmt"
	"strings"
)

// CleanCsv cleans the file
func CleanCsv(bs []byte) string {
	lines := strings.Split(string(bs), "\n")
	lines[0] = CleanCsvHeader(lines[0])
	return strings.Join(lines, "\n")
}

// CleanCsvHeader removes spaces and LowerCase the letters
func CleanCsvHeader(str string) string {
	headers := strings.Split(str, ",")
	for k := range headers {
		headers[k] = strings.TrimSpace(headers[k])
		headers[k] = strings.Title(headers[k])
	}
	str = strings.Join(headers, ",")
	fmt.Println(headers)
	return str
}
