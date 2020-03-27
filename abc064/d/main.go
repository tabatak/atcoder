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
	clCount := 0
	i := 0
	for i < n {
		if s[i] == '(' {
			ans = append(ans, "(")
			opCount++
			i++

		} else {
			// )が続く回数を調べる
			clCount++
			i++
			for j := i; j < n; j++ {
				if s[j] != ')' {
					break
				}
				clCount++
				i++
			}

			tmpOpCount := max(0, clCount-opCount)
			for tmpOpCount > 0 {
				ans = append([]string{"("}, ans...)
				tmpOpCount--
			}
			opCount = max(0, opCount-clCount)

			for clCount > 0 {
				ans = append(ans, ")")
				clCount--
			}
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
