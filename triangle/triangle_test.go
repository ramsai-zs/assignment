package main

import "testing"

func TestTriangle(t *testing.T) {
	tables := []struct {
		s1, s2, s3 int //sides of a triangle
		result     string
	}{
		{s1: 0, s2: 0, s3: 9, result: "triangle is Invalid"},
		{s1: 20, s2: 91, s3: 120, result: "Degenerate Triangle"},
		{s1: 60, s2: 30, s3: 60, result: "Isosceles Triangle"},
		{s1: 55, s2: 55, s3: 55, result: "Equilateral Triangle"},
		{s1: 15, s2: 10, s3: 20, result: "Scalane Triangle"},
	}
	for _, v := range tables {
		output := triangle(v.s1, v.s2, v.s3)

		if output != v.result {
			t.Errorf("accepted %v but got %v", v.result, output)
		}
	}
}
