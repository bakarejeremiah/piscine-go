package main

import "github.com/01-edu/z01"

func main() {
	r := []rune("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}
