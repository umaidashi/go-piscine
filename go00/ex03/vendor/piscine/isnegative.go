package piscine

import "ft"

func IsNegative(nb int) {
	if nb >= 0 {
		ft.PrintRune('F')
	} else {
		ft.PrintRune('T')
	}
	ft.PrintRune('\n')
}
