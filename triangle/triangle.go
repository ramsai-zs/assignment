package main

func triangle(x, y, z int) string {
	if x <= 0 || y <= 0 || z <= 0 { // any side is less or equal to zero it is  invalid
		return "triangle is Invalid"
	} else if x+y <= z || y+z <= x || x+z <= y { // any two side is combined less than the third side is a degenerate Triangle
		return "Degenerate Triangle"
	} else if x == y && y == z && z == x { // all side is said to be equilateral triangle is Equitriangle
		return "Equilateral Triangle"
	} else if x != y && y != z && z != x { //two side equal is said to be Scalane Triangle
		return "Scalane Triangle"
	} else {
		return "Isosceles Triangle" // no side is equal is said to be Isosceles Triangle
	}
}

func main() {
	a := 15
	b := 20
	c := 30
	triangle(a, b, c)
}
