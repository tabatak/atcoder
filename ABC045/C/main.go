package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	sum := 0
	for i := 0; i < (1 << uint(len(s)-1)); i++ {
		// iのbitが立っている個所に+を挿入していく
		tmp := uint(i)
		start := 0
		for j := 1; j <= len(s); j++ {
			if tmp&1 == 1 {
				n, _ := strconv.Atoi(s[start:j])
				sum += n
				start = j
			}

			if j == len(s) {
				n, _ := strconv.Atoi(s[start:j])
				sum += n
			}
			tmp = tmp >> 1
		}

	}

	fmt.Println(sum)
}
