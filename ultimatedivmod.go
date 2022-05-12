package piscine

func UltimateDivMod(a *int, b *int) {
	var div int
	div = *a / *b
	*b = *a % *b
	*a = div
}
