package piscine

import "ft"

func PrintComb2() {
	first := true
	for i := 0; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			if !first {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
			first = false
			ft.PrintRune(rune('0' + (i / 10)))
			ft.PrintRune(rune('0' + (i % 10)))
			ft.PrintRune(' ')
			ft.PrintRune(rune('0' + (j / 10)))
			ft.PrintRune(rune('0' + (j % 10)))
		}
	}
	ft.PrintRune('\n')
}