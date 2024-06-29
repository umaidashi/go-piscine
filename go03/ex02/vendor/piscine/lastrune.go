package piscine

func LastRune(s string) rune {
	len := strLen(s)
	return rune(s[len-1])
}

func strLen(s string) int {
	cnt := 0
	for range s {
		cnt++
	}
	return cnt
}
