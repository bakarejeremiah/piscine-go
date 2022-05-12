package piscine

func IsPrintable(str string) bool {
	p := []rune(str)
	len := 0
	boolian := true
	for range str {
		len++
	}
	for i := 0; i < len; i++ {
		boolian = boolian && (p[i] >= 32 && p[i] < 127)
	}

	return boolian
}
