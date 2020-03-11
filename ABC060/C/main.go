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
	start := 0
	end := t
	for i := 1; i < n; i++ {
		if ts[i] > end {
			sum += end - start
			start = ts[i]
		}
		end = ts[i] + t

	}
	sum += end - start

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
