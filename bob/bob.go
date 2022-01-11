package main

import "strings"

func bob(str string) string {
	if strings.ToUpper(str) == str && strings.Contains(str, "?") { // strings contain capitals and questions he reply "Calm down..."
		return "Calm down, I know what I'm doing!"
	} else if strings.ToUpper(str) == str && len(str) > 0 { //strings is capitals he replies "Whoa,chill out!"
		return "Whoa,chill out!"
	} else if strings.Contains(str, "?") { //if he questions then he replies "Sure."
		return "Sure."
	} else if str == "" { //if string is empty he replies "Fine. Be that way!"
		return "Fine. Be that way!"
	} else {
		return "Whatever." // else he prints "Whatever."
	}
}

func main() {
	input := "Fine. Be that way!"
	bob(input)
}
