package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, h int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &h)

	as := make([]int, n)
	bs := make([]int, n)
	maxA := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])
		fmt.Fscan(r, &bs[i])

		// 振って最も大きいダメージを与えられる刀を保持
		maxA = max(maxA, as[i])
	}
	sort.Sort(sort.IntSlice(as))
	sort.Sort(sort.IntSlice(bs))

	ans := 0

	// まずは投げてみる
	for i := n - 1; i >= 0; i-- {
		if bs[i] <= maxA {
			break
		}

		h -= bs[i]
		ans++
		if h <= 0 {
			break
		}
	}

	// 投げて不足する場合は、振って攻撃する
	if h > 0 {
		swingCount := (h + maxA - 1) / maxA
		ans += swingCount
	}

	fmt.Println(ans)
}

// Utility

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
