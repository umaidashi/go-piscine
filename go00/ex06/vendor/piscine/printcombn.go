package piscine

import "ft"

func PrintCombN(nb int) {
	if nb < 1 || nb > 9 {
		return
	}
	digits := [10]rune{}
	for i := 0; i <= nb; i++ {
		digits[i] = rune('0' + i)
	}
	printDigits(digits, nb)
	ft.PrintRune('\n')
}

func printDigits(digits [10]rune, nb int) {
	for i := 0; i < nb; i++ {
		ft.PrintRune(digits[i])
	}
}