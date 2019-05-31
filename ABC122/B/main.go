package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	length := 0
	for i := 0; i < len(s); i++ {
		tmp := 0
		for j := i; j < len(s); j++ {
			if s[j] != 'A' && s[j] != 'C' && s[j] != 'G' && s[j] != 'T' {
				break
			}
			tmp++
		}
		if length < tmp {
			length = tmp
		}
	}
	fmt.Printf("%d\n", length)
}
