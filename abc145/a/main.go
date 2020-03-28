package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var R int
	fmt.Fscan(r, &R)

	fmt.Println(R * R)
}
