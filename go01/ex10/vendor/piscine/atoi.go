package piscine

func Atoi(s string) int {
	var result int
	sign := 1
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return result * sign
}
