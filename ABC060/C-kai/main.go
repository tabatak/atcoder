package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	var t int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &t)
	ts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ts[i])
	}

	sum := 0
	b := ts[0]
	for i := 1; i < n; i++ {
		sum += min(ts[i]-b, t)
		b = ts[i]
	}
	sum += t
	fmt.Println(sum)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
