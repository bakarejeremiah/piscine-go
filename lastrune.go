package piscine

func LastRune(str string) rune {
	lastrune := []rune(str)
	lens := len([]rune(str))
	return lastrune[lens-1]
}
