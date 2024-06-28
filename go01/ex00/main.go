package main

import (
	"fmt"
	"piscine"
)

func main() {
	n := 0
	piscine.PointOne(&n)
	fmt.Println(n)
	var a *int
	piscine.PointOne(a)
	fmt.Println(a)
}
