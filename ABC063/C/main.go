package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s[i])
	}

	min := math.MaxInt64
	sum := 0
	for i := 0; i < n; i++ {
		sum += s[i]
		if s[i]%10 != 0 {
			min = minInt(min, s[i])
		}
	}

	if sum%10 != 0 {
		fmt.Println(sum)
	} else if min == math.MaxInt64 {
		fmt.Println(0)
	} else {
		fmt.Println(sum - min)
	}
}

func minInt(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
