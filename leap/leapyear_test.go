package main

import "testing"

func TestLeap(t *testing.T) {
	tables := []struct {
		year   int
		result bool
	}{
		{year: 2000, result: true},
		{year: 2010, result: false},
		{year: 2100, result: false},
		{year: 1600, result: true},
		{year: 2017, result: false},
		{year: 2048, result: true},
		{year: 2002, result: false},
		{year: 1896, result: true},
	}

	for _, v := range tables {
		res := leap(v.year)

		if res != v.result {
			t.Errorf("accepted %v but got %v", v.result, res)
		}
	}
}
