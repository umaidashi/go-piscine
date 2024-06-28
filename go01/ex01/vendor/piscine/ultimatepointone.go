package piscine

func UltimatePointOne(n ***int) {
	if n == nil {
		return
	}
	if *n == nil {
		return
	}
	if **n == nil {
		return
	}

	***n = 1
}
