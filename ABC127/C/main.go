package main

import (
	"fmt"
)

func main() {
	var n int
	var m int
	var l int
	var r int

	fmt.Scan(&n)
	fmt.Scan(&m)

	min := 1
	max := n
	for i := 0; i < m; i++ {
		fmt.Scan(&l)
		fmt.Scan(&r)

		if min < l {
			min = l
		}
		if max > r {
			max = r
		}
	}

	if max >= min {
		fmt.Println(max - min + 1)
	} else {
		fmt.Println(0)
	}
}
