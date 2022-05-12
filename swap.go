package piscine

func Swap(a *int, b *int) {
	var tmp int

	tmp = *a
	*a = *b
	*b = tmp
}
