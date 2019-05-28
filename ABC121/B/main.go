package main

import (
	"fmt"
)

func main() {
	var n int
	var m int
	var c int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&c)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	ks := make([][]int, n)
	for i := 0; i < n; i++ {
		k := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&k[j])
		}
		ks[i] = k
	}

	ans := 0
	for i := 0; i < n; i++ {
		sum := c
		for j := 0; j < m; j++ {
			sum = sum + ks[i][j]*b[j]
		}
		if sum > 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
