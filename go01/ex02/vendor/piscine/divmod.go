package piscine

func DivMod(a, b int, div, mod *int) {
	if div != nil {
		*div = a / b
	}

	if mod != nil {
		*mod = a % b
	}
}
