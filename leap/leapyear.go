package main

import "fmt"

/* Every year that is divisible by 4 is a leap year
expected for the century year if year is said to be divisible
by 100 that is said to be a leap year then it is also should be divisible by 400
*/

func leap(y int) bool {
	if y%400 == 0 {
		return true
	} else if y%100 == 0 {
		return false
	} else if y%4 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	year := 2100
	x := leap(year)
	fmt.Println(x)
}
