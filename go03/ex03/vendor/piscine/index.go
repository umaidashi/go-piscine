package piscine

func Index(s, toFind string) int {
	for i, c := range s {
		if c == rune(toFind[0]) {
			for j, t := range toFind {
				if t != rune(s[i+j]) {
					break
				}
			}
			return i
		}
	}
	return -1
}
