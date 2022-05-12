package piscine

func Index(s string, toFind string) int {
	toFindCount := arrayCount(toFind)
	secondIndex := 0
	for _, j := range toFind {
		for i1, i2 := range s {
			if i2 == j {
				if toFindCount == 1 {
					return i1
				} else if toFindCount > 1 {
					for a := 0; a < toFindCount; a++ {
						if s[i1+a] == toFind[a] {
							secondIndex++
						} else {
							return -1
						}
					}
					if secondIndex == toFindCount {
						return i1
					}
				} else {
					return -1
				}
			}
		}
		if secondIndex <= 0 {
			return -1
		}
	}
	return toFindCount
}

func arrayCount(s string) int {
	count := 0
	for i := range s {
		count = i + 1
	}
	return count
}
