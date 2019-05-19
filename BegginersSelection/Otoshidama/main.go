package main

import (
	"fmt"
)

func main() {
	var n int
	var y int

	fmt.Scan(&n)
	fmt.Scan(&y)

	a := -1
	b := -1
	c := -1

LOOP:
	for i := n; i >= 0; i-- {
		if 10000*i > y {
			continue
		}
		for j := n - i; j >= 0; j-- {
			if 10000*i+5000*j > y {
				continue
			}
			if 10000*i+5000*j+1000*(n-i-j) == y {
				a = i
				b = j
				c = n - i - j
				break LOOP
			}

		}
	}

	fmt.Printf("%d %d %d\n", a, b, c)
}
