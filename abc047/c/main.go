package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(r, &s)

	c := s[0]
	ans := 0
	for i := 1; i < len(s); i++ {
		if c != s[i] {
			ans++
		}
		c = s[i]
	}

	fmt.Println(ans)
}
