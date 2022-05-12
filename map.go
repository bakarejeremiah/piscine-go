package piscine

func Map(f func(int) bool, arr []int) []bool {
	jerry := make([]bool, len(arr))

	for i, eji := range arr {
		jerry[i] = f(eji)
	}

	return jerry
}
