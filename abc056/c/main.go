package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var x int
	fmt.Fscan(r, &x)

	t := 0
	sum := 0
	for sum < x {
		t++
		sum = t * (t + 1) / 2
	}
	fmt.Println(t)
}
