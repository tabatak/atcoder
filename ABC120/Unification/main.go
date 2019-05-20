package main

import (
	"fmt"
	"math"
)

func main() {
	var s string

	fmt.Scan(&s)

	c0 := float64(0)
	c1 := float64(0)
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			c0++
		} else if s[i] == '1' {
			c1++
		}
	}

	num := int64(math.Min(c0, c1) * 2)
	fmt.Println(num)
}
