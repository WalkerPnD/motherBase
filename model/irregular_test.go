package model

import "testing"
import "strconv"

func TestHardBounceToBool(t *testing.T) {
	irregular := Irregular{HardBounce: "yes"}
	irrTestsTrue := []string{
		"yes",
		"Yes",
		" yes",
		" Yes",
		" YEs",
		"Sim",
		" Sim",
		"sim",
		" sim",
		"1",
		" 1",
	}
	irrTestsFalse := []string{
		" oj",
		" aidoja",
		" 23",
		" wuv",
		" 343",
		" 09",
		" 90",
		" jli",
		" mciv",
		"aafa",
		"2",
	}
	r := false

	// Test True Case
	for _, irr := range irrTestsTrue {
		irregular.HardBounce = irr
		r = irregular.HardBounceToBool()
		if !r {
			t.Error("expedted: true" + strconv.FormatBool(r))
		}
	}

	// Test False Case
	for _, irr := range irrTestsFalse {
		irregular.HardBounce = irr
		r = irregular.HardBounceToBool()
		if r {
			t.Error("expedted: true" + strconv.FormatBool(r))
		}
	}
}
