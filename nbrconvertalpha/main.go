package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	miinus := 1
	a := len(s) - 1
	arv := 0
	for i := 0; i <= a; i++ {
		if i == 0 && (int(s[i])-43) == 0 {
		} else if i == 0 && (int(s[i])-45) == 0 {
			miinus = -1
		} else {
			arv = arv*10 + (int(s[i]) - 48)
			if 9 < (int(s[i])-48) || (int(s[i])-48) < 0 {
				return 0
			}
		}
	}
	return arv * miinus
}

func main() {
	bbq := os.Args[1:]
	suurus := 96
	for i := 0; i < len(bbq); i++ {
		if bbq[i] == "--upper" {
			suurus = 64
		} else if Atoi(bbq[i]) < 1 || Atoi(bbq[i]) > 26 {
			z01.PrintRune(rune(32))
		} else {
			z01.PrintRune(rune(Atoi(bbq[i]) + suurus))
		}
	}
	if len(bbq) > 0 {
		z01.PrintRune('\n')
	}
}
