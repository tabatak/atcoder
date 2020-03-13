package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, w int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &w)

	vs := make([]int, n)
	ws := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &vs[i])
		fmt.Fscan(r, &ws[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= w; j++ {
			if j < ws[i] {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i+1][j-ws[i]]+vs[i])
			}
		}
	}

	fmt.Println(dp[0][w])

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
