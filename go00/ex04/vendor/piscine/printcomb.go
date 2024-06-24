package piscine

import "ft"

func PrintComb() {
	first := true
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if (i != j && j != k && i != k) {
					if !first {
						ft.PrintRune(',')
						ft.PrintRune(' ')
					}
					first = false;
					ft.PrintRune(i)
					ft.PrintRune(j)
					ft.PrintRune(k)
				}
			}
		}
	} 
}