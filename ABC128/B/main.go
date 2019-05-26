package main

import (
	"fmt"
	"sort"
	"strings"
)

type h struct {
	idx int
	v   int
}

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	var p int
	hm := make(map[string][]h)

	ss := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		fmt.Scan(&p)

		tmpH := h{
			idx: i + 1,
			v:   p,
		}
		if len(hm[s]) == 0 {
			hm[s] = []h{tmpH}
			ss = append(ss, s)
		} else {
			hm[s] = append(hm[s], tmpH)
		}
	}

	sort.Strings(ss)

LOOP:
	for _, city := range ss {
		if strings.TrimSpace(city) == "" {
			continue
		}
		points, _ := hm[city]

		pl := make([]int, 1)
		pm := make(map[int]int, 1)
		for _, point := range points {
			if point.idx == 0 {
				continue LOOP
			}
			pl = append(pl, point.v)
			pm[point.v] = point.idx
		}
		sort.Sort(sort.Reverse(sort.IntSlice(pl)))
		for j := 0; j < len(points); j++ {
			fmt.Println(pm[pl[j]])
		}
	}
}
