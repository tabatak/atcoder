package main

import (
	"fmt"
	"sort"
)

type Store struct {
	A int
	B int
}

type Stores []Store

func (s Stores) Len() int {
	return len(s)
}

func (s Stores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Stores) Less(i, j int) bool {
	return s[i].A < s[j].A
}

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	stores := make(Stores, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&stores[i].A)
		fmt.Scan(&stores[i].B)
	}

	sort.Sort(stores)

	sum := 0
	cansmax := m

	for i := 0; i < n; i++ {
		if cansmax >= stores[i].B {
			sum += stores[i].A * stores[i].B
			cansmax -= stores[i].B
		} else {
			sum += stores[i].A * cansmax
			cansmax = 0
		}
		if cansmax == 0 {
			break
		}
	}

	fmt.Println(sum)
}
