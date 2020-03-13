package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var h, n int
	fmt.Fscan(r, &h)
	fmt.Fscan(r, &n)

	as := make([]int, n)
	bs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])
		fmt.Fscan(r, &bs[i])
	}

	// モンスターの体力iを減らすための最小の体力値
	dp := make([]int, h+1)
	for i := 0; i < h+1; i++ {
		dp[i] = math.MaxInt64
	}
	// 0減らすためには体力の消耗は不要
	dp[0] = 0

	// すべての技を確認
	for i := 0; i < n; i++ {
		//
		for j := 0; j <= h; j++ {
			if dp[j] != math.MaxInt64 {
				// 現在減らした体力 + 技で減らせる体力（モンスターの体力以下の範囲）
				nj := min(j+as[i], h)
				// 消耗する体力を考慮しても最小の体力であれば更新
				dp[nj] = min(dp[nj], dp[j]+bs[i])
			}
		}
	}

	fmt.Println(dp[h])

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
