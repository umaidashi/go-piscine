package piscine

import (
	"ft"
)


func PrintCombN(nb int) {
	if nb < 1 || nb > 9 {
		return
	}
	digits := [10]rune{}
	for i := 0; i <= nb; i++ {
		digits[i] = rune('0' + i)
	}
	printCombNRec(digits, nb, 0, 0)
	ft.PrintRune('\n')
}

// fmt.Println("digits:", digits, "nb:", nb, "index:", index, "start:", start)
func printCombNRec(digits [10]rune, nb, index, start int) {
	if index == nb {
		printDigits(digits, nb)
		if digits[0] != rune('9'-nb+1) {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		return
	}
	for i := start; i <= 9; i++ {
		digits[index] = rune('0' + i)
		printCombNRec(digits, nb, index+1, i+1)
	}
}

func printDigits(digits [10]rune, nb int) {
	for i := 0; i < nb; i++ {
		ft.PrintRune(digits[i])
	}
}
