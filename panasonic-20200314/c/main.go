package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)

	x := c - a - b
	if x < 0 {
		fmt.Println("No")
		return
	}
	if 4*a*b < x*x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
