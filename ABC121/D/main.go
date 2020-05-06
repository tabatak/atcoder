package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	ans := calc(b) ^ calc(a-1)
	fmt.Println(ans)
}

func calc(x int) int {
	count := (x + 1) / 2
	ans := count % 2
	if x%2 == 0 {
		ans ^= x
	}
	return ans
}
