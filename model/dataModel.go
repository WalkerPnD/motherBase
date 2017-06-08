package model

import (
	"strings"
)

func cleanData(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Title(s)
	return s
}
