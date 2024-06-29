package piscine

func NRune(s string, nb int) rune {
	if nb < 0 {
		return 0
	}
	index := nb - 1
	if index >= strLen(s) {
		return 0
	}
	return rune(s[index])
}

func strLen(s string) int {
	cnt := 0
	for range s {
		cnt++
	}
	return cnt
}
