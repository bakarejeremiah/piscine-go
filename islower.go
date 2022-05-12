package piscine

func IsLower(s string) bool {
	x := []rune(s)
	len := 0

	for range s {
		len++
	}
	for i := 0; i < len; i++ {
		if x[i] <= 'z' && x[i] >= 'a' {
		} else {
			return false
		}
	}
	return true
}
