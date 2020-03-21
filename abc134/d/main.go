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
	as := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(r, &as[i])
	}

	// 各箱のボール有無を管理
	bs := make([]int, n+1)
	for i := n; i >= 1; i-- {
		sum := 0
		for j := i + i; j <= n; j += i {
			// 自身より大きい倍数に入っている個数をチェック
			sum ^= bs[j]
		}
		bs[i] = sum ^ as[i]
	}

	var ans []int
	for i := 1; i < n+1; i++ {
		if bs[i] == 1 {
			ans = append(ans, i)
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, len(ans))
	for i := 0; i < len(ans); i++ {
		fmt.Fprintf(w, "%d ", ans[i])
	}
}
