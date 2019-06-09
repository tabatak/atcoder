package main

import (
	"fmt"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	as := make([]int, m)
	ms := make(map[int]bool)
	for i := 0; i < m; i++ {
		fmt.Scan(&as[i])
		ms[as[i]] = true
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if ms[i] {
			continue
		}
		dp[i] = dp[i-1]
		if 2 <= i {
			dp[i] = (dp[i] + dp[i-2]) % 1000000007
		}
	}

	println(dp[n])
}
