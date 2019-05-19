package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)

	alice := 0
	bob := 0
	aliceFlg := true
	for i := n - 1; i >= 0; i-- {
		if aliceFlg {
			//Alice
			alice += nums[i]
			aliceFlg = false
		} else {
			//Bob
			bob += nums[i]
			aliceFlg = true
		}
	}

	fmt.Println(alice - bob)
}
