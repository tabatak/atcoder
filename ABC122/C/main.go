package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var q int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&q)
	fmt.Scan(&s)

	ls := make([]int, q)
	rs := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&ls[i])
		fmt.Scan(&rs[i])
	}

	for i := 0; i < q; i++ {
		target := s[ls[i]-1 : rs[i]]
		fmt.Println(strings.Count(target, "AC"))
	}
}
