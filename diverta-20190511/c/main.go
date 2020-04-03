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
	fmt.Fscan(r, &n)
	ss := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ss[i])
	}

	suffixACount := 0
	prefixBCount := 0
	aToBCount := 0
	abCount := 0
	for _, s := range ss {
		abCount += strings.Count(s, "AB")

		if s[0] == 'B' && s[len(s)-1] == 'A' {
			aToBCount++
		} else if s[0] == 'B' {
			prefixBCount++
		} else if s[len(s)-1] == 'A' {
			suffixACount++
		}
	}

	if aToBCount > 0 {
		abCount += aToBCount - 1
		if suffixACount > 0 && prefixBCount > 0 {
			abCount += 2
			abCount += min(suffixACount-1, prefixBCount-1)
		} else if suffixACount > 0 {
			abCount++
		} else if prefixBCount > 0 {
			abCount++
		}
	} else {
		abCount += min(suffixACount, prefixBCount)
	}

	fmt.Println(abCount)
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
