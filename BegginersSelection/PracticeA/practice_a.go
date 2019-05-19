package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var c int
	var t string

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&t)

	sum := a + b + c
	fmt.Printf("%d %s\n", sum, t)
}
