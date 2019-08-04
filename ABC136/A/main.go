package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	tmp := a - b
	if c >= tmp {
		fmt.Println(c - tmp)
	} else {
		fmt.Println(0)
	}
}
