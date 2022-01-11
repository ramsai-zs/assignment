package main

import "testing"

func TestBob(t *testing.T) {
	tables := []struct {
		input, output string
	}{
		{input: "RAM LOCKDOWN?", output: "Calm down, I know what I'm doing!"},
		{input: "GET OUT OF MY AREA", output: "Whoa,chill out!"},
		{input: "Hey out of my area?", output: "Sure."},
		{input: "hehe cooldown", output: "Whatever."},
		{input: "", output: "Fine. Be that way!"},
	}
	for _, v := range tables {
		result := bob(v.input)
		if result != v.output {
			t.Errorf("accepted %v but got %v", v.output, result)
		}
	}
}
