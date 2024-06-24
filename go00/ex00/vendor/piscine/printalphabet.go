package piscine

import "ft"

func Printalphabet() {
	for i := 'a'; i <= 'z'; i++ {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}
