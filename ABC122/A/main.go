package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)

	ans := ""
	if a == "A" {
		ans = "T"
	} else if a == "C" {
		ans = "G"
	} else if a == "G" {
		ans = "C"
	} else if a == "T" {
		ans = "A"
	}
	fmt.Printf("%s\n", ans)
	fmt.Printf("thinking...")
}
