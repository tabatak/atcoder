package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	ans := 0
	for i := a; i <= b; i++ {
		ans ^= i
	}

	fmt.Println(ans)
}
