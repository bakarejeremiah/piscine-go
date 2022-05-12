package piscine

func RecursiveFactorial(nb int) int {
	if nb > 1234 {
		return 0
	}

	if nb < 0 {
		return 0
	}

	if nb == 1 || nb == 0 {
		return 1
	}

	return nb * RecursiveFactorial(nb-1)
}
