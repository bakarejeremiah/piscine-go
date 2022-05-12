package piscine

func IsUpper(str string) bool {
	x := []rune(str)
	len := 0

	for range str {
		len++
	}
	for i := 0; i < len; i++ {
		if x[i] >= 'A' && x[i] <= 'Z' {
		} else {
			return false
		}
	}
	return true
}
