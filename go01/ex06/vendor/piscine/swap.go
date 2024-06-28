package piscine

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

// https://go.dev/wiki/SimultaneousAssignment
// https://go.dev/ref/spec#Assignment_statements
