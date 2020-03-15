package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var alpha = "abcdefghijklmnopqrstuvwxyz"
var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	fmt.Fscan(r, &n)
	dfs("", 0)
}

func dfs(s string, mx int) {
	if len(s) == n {
		fmt.Fprintln(w, s)
		return
	}

	for i := 0; i < mx+1; i++ {
		t := s + string(alpha[i])
		dfs(t, max(i+1, mx))
	}
	return
}

// Utility
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
