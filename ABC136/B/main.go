package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nStr string
	fmt.Scan(&nStr)

	n, _ := strconv.Atoi(nStr)
	l := len(nStr)

	if l == 1 {
		fmt.Println(n)
	} else if l == 2 {
		fmt.Println(9)
	} else if l == 3 {
		fmt.Println(n - 99 + 9)
	} else if l == 4 {
		fmt.Println((999 - 99) + 9)
	} else if l == 5 {
		fmt.Println(n - 9999 + (999 - 99) + 9)
	} else if l == 6 {
		fmt.Println((99999 - 9999) + (999 - 99) + 9)
	}

}
