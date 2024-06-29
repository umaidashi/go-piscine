package piscine

func Compare(a, b string) int {
	lenA := strLen(a)
	lenB := strLen(b)
	min := lenA
	if min > lenB {
		min = lenB
	}
	for i := 0; i < min; i++ {
		if a[i] != b[i] {
			if a[i] > b[i] {
				return 1
			} else {
				return -1
			}
		}
	}
	if lenA > lenB {
		return 1
	} else if lenA < lenB {
		return -1
	}
	return 0
}

func strLen(s string) int {
	cnt := 0
	for range s {
		cnt++
	}
	return cnt
}
