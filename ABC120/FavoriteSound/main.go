package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	num := b / a
	if num <= c {
		fmt.Println(num)
	} else {
		fmt.Println(c)
	}
}
