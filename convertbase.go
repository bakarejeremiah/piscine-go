package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	ResBase := 0
	for _, value1 := range nbr {
		for index2, value2 := range baseFrom {
			if value1 == value2 {
				ResBase = ResBase*StringLen(baseFrom) + index2
			}
		}
	}

	x := ""
	for ResBase != 0 {

		x = string(baseTo[ResBase%StringLen(baseTo)]) + x
		ResBase /= StringLen(baseTo)
	}

	return x
}

func StringLen(str string) int {
	i := 0
	for range str {
		i++
	}
	return i
}
