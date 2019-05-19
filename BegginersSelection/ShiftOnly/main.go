package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	count := 0
	for {
		oddFlg := false
		for i := 0; i < n; i++ {
			if nums[i]%2 == 0 {
				nums[i] = nums[i] / 2
			} else {
				oddFlg = true
				break
			}
		}
		if oddFlg {
			break
		}
		count++
	}

	fmt.Println(count)
}
