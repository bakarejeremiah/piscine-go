package piscine

import "github.com/01-edu/z01"

func PrintM(num int) {
	s := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for k := 1; k <= num%10; k++ {
		s++
	}
	for k := -1; k >= num%10; k-- {
		s++
	}
	if num/10 != 0 {
		PrintM(num / 10)
	}
	z01.PrintRune(s)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintM(n)
}
