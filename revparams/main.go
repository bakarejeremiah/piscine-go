package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := os.Args
	count := 0
	for range result {
		count++
	}
	for i := count - 1; i > 0; i-- {
		for _, i2 := range result[i] {
			z01.PrintRune(i2)
		}
		z01.PrintRune(10)
	}
}
