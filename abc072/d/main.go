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
	ps := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ps[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		// i = piだったらswapしてみる
		if i+1 == ps[i] {
			if i != n-1 {
				tmp := ps[i]
				ps[i] = ps[i+1]
				ps[i+1] = tmp
			} else {
				tmp := ps[i]
				ps[i] = ps[i-1]
				ps[i-1] = tmp
			}
			ans++
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, ans)
}
