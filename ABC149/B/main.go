package main

import (
	"fmt"
)

func main() {
	var a, b, k int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&k)

	fromA := min(a, k)
	a = a - fromA
	k = k - fromA
	fromB := min(b, k)
	b = b - fromB
	fmt.Printf("%d %d\n", a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
