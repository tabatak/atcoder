package main

import (
	"fmt"
)

var mp map[int]int

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	as := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&as[i])
	}

	mp = make(map[int]int)
	ret := 1
	before := 0
	for i := 0; i < m; i++ {
		ret = ret * answer(as[i]-before-1)

		if i > 0 && as[i] == as[i-1]+1 {
			ret = 0
			break
		}

		before = as[i] + 1
	}

	ret = ret * (n - as[len(as)-1] - 1)

	fmt.Println(ret % 1000000007)
}

func answer(n int) int {
	if val, ok := mp[n-1]; ok {
		return val
	}

	if n == 1 {
		mp[n-1] = 1
		return 1
	} else if n == 2 {
		mp[n-1] = 2
		return 2
	} else {
		memo := answer(n-1) + answer(n-2)
		mp[n-1] = memo
		return memo
	}
}
