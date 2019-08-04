package main

import (
	"bufio"
	"fmt"
	"os"
)

//正解にたどりつけず。。。
func main() {
	var s string
	fmt.Scan(&s)

	cs := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		cs[i] = 1
	}

	lCount := 0
	R := byte('R')
	L := byte('L')
	prev := byte('R')
	for i := 1; i < len(s); i++ {
		current := s[i]

		if prev == R && current == R {
			cs[i] += cs[i-1]
			cs[i-1]--

		} else if prev == R && current == L {
			cs[i] = cs[i-1]
			cs[i-1] = 1

		} else if prev == L && current == R {
			cs[i-lCount-1] += lCount
			// cs[i]++
			lCount = 0

		} else if prev == L && current == L {
			cs[i-lCount]++
			cs[i]--
			lCount++
		}

		prev = current

	}

	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(s); i++ {
		w.Write([]byte(fmt.Sprintf("%d", cs[i])))
		if i != len(s)-1 {
			w.Write([]byte(" "))
		}
	}
	w.Write([]byte("\n"))
	w.Flush()
}
