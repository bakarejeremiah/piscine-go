package main

import "github.com/01-edu/z01"

func main() {
	var z rune
	z = 'z'
	for z >= 'a' {
		z01.PrintRune(z)
		z--
	}
	z01.PrintRune('\n')
}
