package piscine

import "ft"

func Printdigits() {
	for i := '0'; i <= '9'; i++ {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}