package piscine

func IsAlpha(str string) bool {
	x := []rune(str)
	len := 0

	for range str {
		len++
	}
	for i := 0; i < len; i++ {
		if x[i] >= 'A' && x[i] <= 'Z' || x[i] <= 'z' && x[i] >= 'a' || x[i] >= '0' && x[i] <= '9' {
		} else {
			return false
		}
	}
	return true
}
