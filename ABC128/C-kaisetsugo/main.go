package main

import (
	"fmt"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	var k int
	var s int

	a := make([]int, n) //スイッチごとのつながっている電球を表すビット列
	for i := 0; i < m; i++ {
		fmt.Scan(&k)

		for j := 0; j < k; j++ {
			fmt.Scan(&s)
			s--
			a[s] |= 1 << uint(i)

		}
	}

	var x int
	p := 0
	for i := 0; i < m; i++ {
		fmt.Scan(&x)

		p |= x << uint(i)
	}

	//スイッチの全パターン網羅
	ans := 0
	for s := 0; s < (1 << uint(n)); s++ {
		t := 0

		for i := 0; i < n; i++ {
			if s>>uint(i)&1 == 1 {
				t ^= a[i]
			}
		}

		if t == p {
			ans++
		}
	}

	fmt.Println(ans)
}
