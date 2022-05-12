package piscine

func SplitWhiteSpaces(str string) []string {
	i := 0
	bykva := true
	for _, word := range str {
		if word == ' ' || word == '\n' || word == '\t' {
			bykva = true
		} else {
			if bykva {
				i++
				bykva = false
			}
		}
	}
	x := 0
	a := make([]string, i)
	ok := true
	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		a[x] = a[x] + string(c)
		ok = false
	}
	return a
}
