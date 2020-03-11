package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	min := math.MaxInt32
	sum := 0
	for i := 0; i < n; i++ {
		sum += s[i]
		if s[i]%10 != 0 {
			min = int(math.Min(float64(min), float64(s[i])))
		}
	}

	if sum%10 != 0 {
		println(sum)
	} else if min == math.MaxInt32 {
		println(0)
	} else {
		println(sum - min)
	}
}
