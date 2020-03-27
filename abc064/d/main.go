package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &s)

	var ans []string
	opCount := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			ans = append(ans, "(")
			opCount++
		} else {
			if opCount > 0 {
				// 直前の ( に対応する場合
				opCount--
			} else {
				// 直前に対応する ( が存在しない場合、先頭に ( を挿入
				ans = append([]string{"("}, ans...)
			}
			ans = append(ans, ")")
		}
	}

	for opCount > 0 {
		ans = append(ans, ")")
		opCount--
	}

	fmt.Println(strings.Join(ans, ""))
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
