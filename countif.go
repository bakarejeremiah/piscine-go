package piscine

func CountIf(f func(string) bool, tab []string) int {
	nb := 0
	for _, a := range tab {
		if f(a) {
			nb++
		}
	}
	return nb
}
