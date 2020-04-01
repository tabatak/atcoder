package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &m)

	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	b := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &b[i])
	}

	ans := false
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			if check(i, j, a, b) {
				ans = true
				break
			}
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func check(startI, startJ int, a, b []string) bool {
	result := true
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if a[i+startI][j+startJ] != b[i][j] {
				result = false
				break
			}
		}
	}
	return result
}
