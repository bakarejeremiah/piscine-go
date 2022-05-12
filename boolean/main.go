package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	count := 0
	for i := range args {
		count = i + 1
	}
	if isEven(count) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func printStr(str string) {
	arrayStr := []rune(str)
	countt := 0
	for i := range arrayStr {
		countt = i + 1
	}

	for i := 0; i < countt; i++ {
		z01.PrintRune(arrayStr[i])
	}

	z01.PrintRune('\n')
}
