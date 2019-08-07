package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scan(&s)

	ans := make([]int, len(s))

	for i := 0; i < 2; i++ {
		count := 0
		for j := 0; j < len(s); j++ {
			if s[j] == 'R' {
				count++
			} else {
				ans[j] += count / 2
				ans[j-1] += (count + 1) / 2
				count = 0
			}
		}

		s = ReverseUsingRunes(s)
		ans = ReverseIntSlice(ans)
	}

	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(s); i++ {
		w.Write([]byte(fmt.Sprintf("%d", ans[i])))
		if i != len(s)-1 {
			w.Write([]byte(" "))
		}
	}
	w.Write([]byte("\n"))
	w.Flush()

}

func ReverseUsingRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	for i := 0; i < len(runes); i++ {
		if runes[i] == 'R' {
			runes[i] = 'L'
		} else {
			runes[i] = 'R'
		}
	}
	return string(runes)
}

func ReverseIntSlice(is []int) []int {
	reverse := make([]int, len(is))
	for i, j := 0, len(is)-1; i < len(is); i, j = i+1, j-1 {
		reverse[i] = is[j]
	}
	return reverse
}
