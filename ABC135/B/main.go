package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	swapMin := n + 1
	swapMax := 0
	swapMinIndex := 0
	swapMaxIndex := 0
	swap := false
	for i := 1; i < n; i++ {
		if p[i-1] > p[i] {
			swap = true
			if swapMin > p[i] {
				swapMin = p[i]
				swapMinIndex = i
			}
			if swapMax < p[i] {
				swapMax = p[i-1]
				swapMaxIndex = i - 1
			}
		}
	}

	if swap == false {
		fmt.Println("YES")
		return
	}

	//swap
	tmp := p[swapMinIndex]
	p[swapMinIndex] = p[swapMaxIndex]
	p[swapMaxIndex] = tmp

	for i := 1; i < n; i++ {
		if p[i-1] > p[i] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
