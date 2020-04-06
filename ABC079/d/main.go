package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(r, &h)
	fmt.Fscan(r, &w)

	dist := make([][]int, 10)
	for i := 0; i < 10; i++ {
		row := make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Fscan(r, &row[j])
		}
		dist[i] = row
	}
	for k := 0; k < 10; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	ans := 0
	var tmpW int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(r, &tmpW)
			if tmpW >= 0 {
				ans += dist[tmpW][1]
			}
		}
	}

	fmt.Println(ans)
}

// Utility

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
