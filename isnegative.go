package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	runei := 'T'
	runej := 'F'
	if nb < 0 {
		z01.PrintRune(runei)
	} else {
		z01.PrintRune(runej)
	}
	z01.PrintRune('\n')
}
