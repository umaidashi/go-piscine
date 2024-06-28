package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < intsLen(table); i++ {
		for j := i + 1; j < intsLen(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func intsLen(table []int) int {
	cnt := 0
	for range table {
		cnt++
	}
	return cnt
}
