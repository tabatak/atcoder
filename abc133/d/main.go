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
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])
	}

	// 山1に降った量/2を固定する（算出する）
	x := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			x += as[i]
		} else {
			x -= as[i]
		}
	}

	ans := make([]int, n)
	ans[0] = x / 2
	for i := 0; i < n-1; i++ {
		ans[i+1] = as[i] - ans[i]
	}
	for i := 0; i < n; i++ {
		ans[i] *= 2
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d ", ans[i])
	}

}
