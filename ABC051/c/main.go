package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var sx, sy, tx, ty int
	fmt.Fscan(r, &sx)
	fmt.Fscan(r, &sy)
	fmt.Fscan(r, &tx)
	fmt.Fscan(r, &ty)

	dx := tx - sx
	dy := ty - sy

	var ans []string
	// 1回目往復
	// 往路
	for i := 0; i < dx; i++ {
		ans = append(ans, "R")
	}
	for i := 0; i < dy; i++ {
		ans = append(ans, "U")
	}
	// 復路
	for i := 0; i < dx; i++ {
		ans = append(ans, "L")
	}
	for i := 0; i < dy; i++ {
		ans = append(ans, "D")
	}

	// 2回目往復
	// 往路
	ans = append(ans, "D")
	for i := 0; i < dx+1; i++ {
		ans = append(ans, "R")
	}
	for i := 0; i < dy+1; i++ {
		ans = append(ans, "U")
	}
	ans = append(ans, "L")
	// 復路
	ans = append(ans, "U")
	for i := 0; i < dx+1; i++ {
		ans = append(ans, "L")
	}
	for i := 0; i < dy+1; i++ {
		ans = append(ans, "D")
	}
	ans = append(ans, "R")

	fmt.Println(strings.Join(ans, ""))
}
