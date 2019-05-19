package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var a int
	var b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	someSums := 0
	for i := 1; i <= n; i++ {
		sum := 0

		s := strconv.Itoa(i)
		for j := 0; j < len(s); j++ {
			num, _ := strconv.Atoi(string(s[j]))
			sum += num
		}

		if a <= sum && sum <= b {
			someSums += i
		}
	}

	fmt.Println(someSums)
}
