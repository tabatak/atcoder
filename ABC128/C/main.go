package main

import (
	"fmt"
)

type denkyu struct {
	sw []int
	p  int
}

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	var k int
	var s int
	denkyus := make(map[int]denkyu, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&k)
		d := denkyu{
			sw: make([]int, k),
		}
		for j := 0; j < k; j++ {
			fmt.Scan(&s)
			d.sw[j] = s
		}
		denkyus[i] = d
	}

	var p int
	for i := 0; i < m; i++ {
		fmt.Scan(&p)
		d := denkyus[i]
		d.p = p
		denkyus[i] = d
	}

	fmt.Println(denkyus)

	fmt.Println("ギブアップ")
}
