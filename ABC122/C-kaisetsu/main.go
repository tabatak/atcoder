package main

import (
	"fmt"
)

func main() {
	var n int
	var q int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&q)
	fmt.Scan(&s)

	//累積和を計算する
	sum := make([]int, n+1)
	// ACが出る可能性があるのは2文字目から(0-indexedで1から)
	for i := 1; i < n; i++ {
		if s[i-1] == 'A' && s[i] == 'C' {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i]
		}
	}

	ls := make([]int, q)
	rs := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&ls[i])
		fmt.Scan(&rs[i])
	}

	for i := 0; i < q; i++ {
		ans := sum[rs[i]] - sum[ls[i]]
		fmt.Println(ans)
	}
}
