package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var n int
	fmt.Scan(&n)
	hs := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		hs[i], _ = strconv.Atoi(sc.Text())
	}

	str := "Yes"
	for i := 1; i < n; i++ {

		if hs[i-1] < hs[i] {
			//右の方が大きければ1減らす
			hs[i]--
		}
		if hs[i-1] > hs[i] {
			str = "No"
			break
		}
	}

	fmt.Println(str)
}
