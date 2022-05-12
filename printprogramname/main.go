package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	tekst := []rune(os.Args[0])
	for i := range tekst {
		if i > 1 {
			z01.PrintRune(tekst[i])
		}
	}
	z01.PrintRune('\n')
}
