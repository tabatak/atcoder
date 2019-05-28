package main

import (
	"fmt"
)

func main() {
	var h int
	var w int
	var hd int
	var wd int
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&hd)
	fmt.Scan(&wd)

	fmt.Println((h - hd) * (w - wd))
}
