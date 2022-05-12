package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, a := range arr {
		if f(a) {
			return true
		}
	}
	return false
}
