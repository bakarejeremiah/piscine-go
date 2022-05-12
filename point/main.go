package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNumber(point int) {
	c := '0'
	for i := 0; i < point/10; i++ {
		c++
	}
	z01.PrintRune(c)

	d := '0'
	for i := 0; i < point%10; i++ {
		d++
	}
	z01.PrintRune(d)
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNumber(points.y)
	z01.PrintRune('\n')
}
