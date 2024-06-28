package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	res := RecursiveFactorial(nb - 1)
	if res < 0 {
		return 0
	}
	ans := nb * res
	return ans
}
