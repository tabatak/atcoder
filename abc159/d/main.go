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
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &as[i])
		mp[as[i]]++
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	count := make(map[int]int)
	sum := 0
	for k, v := range mp {
		count[k] += v * (v - 1) / 2
		sum += count[k]
	}

	for i := 0; i < n; i++ {
		c := mp[as[i]] - 1
		self := c * (c - 1) / 2

		ans := sum - count[as[i]] + self
		fmt.Fprintln(w, ans)
	}
}
