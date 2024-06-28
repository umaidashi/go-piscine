package piscine

func UltimateDivMod(a *int, b *int) {
	if a == nil || b == nil {
		return
	}
	tmp := *a
	*a = *a / *b
	*b = tmp % *b
}
