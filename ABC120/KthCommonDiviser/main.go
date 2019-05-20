package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var k int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&k)

	i := a
	if a < b {
		i = b
	}
	num := 0
	count := 0
	for {
		if a%i == 0 && b%i == 0 {
			count++
			if count == k {
				num = i
				break
			}
		}
		i--
	}

	fmt.Println(num)
}
