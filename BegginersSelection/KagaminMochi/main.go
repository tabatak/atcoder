package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	kagamimochis := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&kagamimochis[i])
	}
	sort.Ints(kagamimochis)

	dansu := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			dansu++
		} else {
			if kagamimochis[i] > kagamimochis[i-1] {
				dansu++
			}
		}
	}

	fmt.Println(dansu)
}
