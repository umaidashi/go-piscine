package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var ans int = 1
	for i := 1; i <= nb; i++ {
		if ans*i < ans {
			return 0
		}
		ans *= i
	}
	return int(ans)
}

// fmt.Println(ans, i, ans*i)
