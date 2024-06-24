package piscine

import "ft"

func Printreversealphabet() {
	for i := 'z'; i >= 'a'; i-- {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}