package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	as := make([]int, n)
	mx := -1
	mxIndex := -1
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])

		if mx < as[i] {
			mx = as[i]
			mxIndex = i
		}
	}

	half := mx / 2
	half2 := mx/2 + 1

	t := 0
	for i := 0; i < n; i++ {
		if i == mxIndex {
			continue
		}
		if abs(half-t) > abs(half-as[i]) {
			t = as[i]
		}
		if abs(half2-t) > abs(half2-as[i]) {
			t = as[i]
		}
	}

	fmt.Printf("%d %d", mx, t)
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
